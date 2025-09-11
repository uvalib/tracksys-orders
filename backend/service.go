package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ServiceContext contains common data used by all handlers
type serviceContext struct {
	Version         string
	TrackSysURL     string
	GDB             *gorm.DB
	ConfirmTemplate *template.Template
	SMTP            smtpConfig
	Dev             devConfig
}

// InitializeService sets up the service context for all API handlers
func initializeService(version string, cfg *configData) *serviceContext {
	ctx := serviceContext{Version: version,
		TrackSysURL: cfg.tracksysURL,
		Dev:         cfg.dev,
		SMTP:        cfg.smtp}

	log.Printf("INFO: connecting to db...")
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		cfg.db.User, cfg.db.Pass, cfg.db.Host, cfg.db.Name)
	gdb, err := gorm.Open(mysql.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ctx.GDB = gdb
	log.Printf("INFO: db connection established")

	log.Printf("INFO: load confirmation email template")
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
	tpl, err := template.New("confirmation.html").Funcs(funcMap).ParseFiles("./templates/confirmation.html")
	if err != nil {
		log.Fatalf("unable to load confirmation template: %s", err.Error())
	}
	ctx.ConfirmTemplate = tpl

	return &ctx
}

func (svc *serviceContext) authenticate(c *gin.Context) {
	log.Printf("Checking authentication headers...")
	log.Printf("Dump all request headers ==================================")
	for name, values := range c.Request.Header {
		for _, value := range values {
			log.Printf("%s=%s\n", name, value)
		}
	}
	log.Printf("END header dump ===========================================")

	computingID := c.GetHeader("remote_user")
	if svc.Dev.authUser != "" {
		computingID = svc.Dev.authUser
		log.Printf("Using dev auth user ID: %s", computingID)
	}
	if computingID == "" {
		log.Printf("ERROR: Expected auth header not present in request. Not authorized.")
		c.Redirect(http.StatusFound, "/forbidden")
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/granted?user=%s", computingID))
}

func (svc *serviceContext) healthCheck(c *gin.Context) {
	type hcResp struct {
		Healthy bool   `json:"healthy"`
		Message string `json:"message,omitempty"`
	}
	hcMap := make(map[string]hcResp)
	hcMap["ts-orders"] = hcResp{Healthy: true}

	hcMap["ts-database"] = hcResp{Healthy: true}
	sqlDB, err := svc.GDB.DB()
	if err != nil {
		hcMap["ts-database"] = hcResp{Healthy: false, Message: err.Error()}
	} else {
		err := sqlDB.Ping()
		if err != nil {
			hcMap["ts-database"] = hcResp{Healthy: false, Message: err.Error()}
		}
	}

	c.JSON(http.StatusOK, hcMap)
}

func (svc *serviceContext) getVersion(c *gin.Context) {
	build := "unknown"

	// cos our CWD is the bin directory
	files, _ := filepath.Glob("../buildtag.*")
	if len(files) == 1 {
		build = strings.Replace(files[0], "../buildtag.", "", 1)
	}

	vMap := make(map[string]string)
	vMap["version"] = Version
	vMap["build"] = build
	c.JSON(http.StatusOK, vMap)
}

type emailRequest struct {
	Subject string
	To      []string
	ReplyTo string
	CC      string
	From    string
	Body    string
}

// SendEmail will and send an email to the specified recipients
func (svc *serviceContext) SendEmail(request *emailRequest) error {
	mail := gomail.NewMessage()
	mail.SetHeader("Subject", request.Subject)
	mail.SetHeader("To", request.To...)
	mail.SetHeader("From", request.From)
	if request.ReplyTo != "" {
		mail.SetHeader("Reply-To", request.ReplyTo)
	}
	if len(request.CC) > 0 {
		mail.SetHeader("Cc", request.CC)
	}
	mail.SetBody("text/html", request.Body)

	if svc.Dev.fakeSMTP {
		log.Printf("Email is in dev mode. Logging message instead of sending")
		log.Printf("==================================================")
		mail.WriteTo(log.Writer())
		log.Printf("==================================================")
		return nil
	}

	log.Printf("Sending %s email to %s", request.Subject, strings.Join(request.To, ","))
	if svc.SMTP.Pass != "" {
		dialer := gomail.Dialer{Host: svc.SMTP.Host, Port: svc.SMTP.Port, Username: svc.SMTP.User, Password: svc.SMTP.Pass}
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		return dialer.DialAndSend(mail)
	}

	log.Printf("Sending email with no auth")
	dialer := gomail.Dialer{Host: svc.SMTP.Host, Port: svc.SMTP.Port}
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer.DialAndSend(mail)
}
