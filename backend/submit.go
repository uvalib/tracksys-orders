package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type itemInfo struct {
	Title       string `json:"title"`
	Pages       string `json:"pages"`
	CallNumber  string `json:"callNumber"`
	Author      string `json:"author"`
	Published   string `json:"published"`
	Location    string `json:"location"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type requestInfo struct {
	DueDate       time.Time `json:"dueDate"`
	Instructions  string    `json:"specialInstructions"`
	IntendedUseID int64     `json:"intendedUseID"`
	IntendedUse   string    `json:"-"`
	Format        string    `json:"-"`
	Resolution    string    `json:"-"`
}

type submission struct {
	User      customerJSON  `json:"user"`
	Addresses []addressJSON `json:"addresses"`
	Request   requestInfo   `json:"request"`
	Items     []itemInfo    `json:"items"`
	OrderID   int64         `json:"-"`
}

type orderRec struct {
	ID            int64     `db:"id"`
	CustomeerID   int64     `db:"customer_id"`
	DateDue       time.Time `db:"date_due"`
	Instructions  string    `db:"special_instructions"`
	DateSubmitted time.Time `db:"date_request_submitted"`
	Status        string    `db:"order_status"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u orderRec) TableName() string {
	return "orders"
}

type itemRec struct {
	ID            int64          `db:"id"`
	OrderID       int64          `db:"order_id"`
	IntendedUseID int64          `db:"intended_use_id"`
	Title         string         `db:"title"`
	Pages         string         `db:"pages"`
	CallNumber    sql.NullString `db:"call_number"`
	Author        sql.NullString `db:"author"`
	Published     sql.NullString `db:"year"`
	Location      sql.NullString `db:"location"`
	Link          sql.NullString `db:"source_url"`
	Description   sql.NullString `db:"description"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u itemRec) TableName() string {
	return "order_items"
}

func (svc *serviceContext) submitRequest(c *gin.Context) {
	log.Printf("INFO: order request submitted ")
	var req submission
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ERROR: unable to parse order submission payload: %s", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("INFO: lookup intended use %d", req.Request.IntendedUseID)
	q := svc.DB.NewQuery("select id,description,deliverable_format,deliverable_resolution from intended_uses where id={:id}")
	q.Bind(dbx.Params{"id": req.Request.IntendedUseID})
	var useRec intendedUseRec
	dbErr := q.One(&useRec)
	if dbErr != nil {
		log.Printf("ERROR: unable to get intended use %d: %s", req.Request.IntendedUseID, dbErr.Error())
		c.String(http.StatusInternalServerError, dbErr.Error())
		return
	}
	req.Request.IntendedUse = useRec.Name
	req.Request.Format = useRec.Format
	req.Request.Resolution = useRec.Resolution.String

	log.Printf("INFO: load confirmation email template")
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dateFmt": func(d time.Time) string {
			return d.Format("2006-01-02")
		},
	}
	tpl, err := template.New("confirmation.html").Funcs(funcMap).ParseFiles("./templates/confirmation.html")
	if err != nil {
		log.Printf("ERROR: unable to load confirmation template: %s", err.Error())
		c.String(http.StatusOK, "ok")
		return
	}

	log.Printf("INFO: create order and order items")
	order := orderRec{CustomeerID: req.User.ID, DateDue: req.Request.DueDate, Instructions: req.Request.Instructions,
		DateSubmitted: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Status: "requested"}
	log.Printf("INFO: create order %+v", order)
	err = svc.DB.Model(&order).Insert()
	if err != nil {
		log.Printf("ERROR: unable to create order: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("INFO: add %d items to order", len(req.Items))
	for _, item := range req.Items {
		rec := itemRec{OrderID: order.ID, IntendedUseID: req.Request.IntendedUseID, Title: item.Title, Pages: item.Pages}
		if item.CallNumber != "" {
			rec.CallNumber.String = item.CallNumber
			rec.CallNumber.Valid = true
		}
		if item.Author != "" {
			rec.Author.String = item.Author
			rec.Author.Valid = true
		}
		if item.Published != "" {
			rec.Published.String = item.Published
			rec.Published.Valid = true
		}
		if item.Location != "" {
			rec.Location.String = item.Location
			rec.Location.Valid = true
		}
		if item.Link != "" {
			rec.Link.String = item.Link
			rec.Link.Valid = true
		}
		if item.Description != "" {
			rec.Description.String = item.Description
			rec.Description.Valid = true
		}
		itemErr := svc.DB.Model(&rec).Insert()
		if err != nil {
			log.Printf("ERROR: unable to add items to order: %s", itemErr.Error())
			c.String(http.StatusInternalServerError, itemErr.Error())
			return
		}
	}
	log.Printf("INFO: all items created")

	log.Printf("INFO: generate confirmation email for %s", req.User.Email)
	req.OrderID = order.ID
	var renderedEmail bytes.Buffer
	err = tpl.Execute(&renderedEmail, req)
	if err != nil {
		log.Printf("ERROR: unable to generate confirmation email for %s: %s", req.User.Email, err.Error())
	} else {
		to := make([]string, 0)
		to = append(to, req.User.Email)
		sub := fmt.Sprintf("UVA Digital Production Group - Request #%d Confirmation", req.OrderID)
		eRequest := emailRequest{Subject: sub, To: to, From: svc.SMTP.Sender, CC: svc.SMTP.Sender, Body: renderedEmail.String()}
		sendErr := svc.SendEmail(&eRequest)
		if sendErr != nil {
			log.Printf("ERROR: Unable to send confirm email to %s: %s", req.User.Email, sendErr.Error())
		}
	}

	c.String(http.StatusOK, "ok")
}
