package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type submission struct {
	DateDue             time.Time   `json:"dateDue"`
	IntendedUseID       int64       `json:"intendedUseID"`
	SpecialInstructions string      `json:"specialInstructions"`
	Items               []orderItem `json:"items"`
}

type order struct {
	ID                   int64
	CustomerID           int64
	OrderStatus          string
	DateDue              time.Time
	SpecialInstructions  string
	DateRequestSubmitted time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type orderItem struct {
	ID            int64  `json:"-"`
	OrderID       int64  `json:"-"`
	IntendedUseID int64  `json:"-"`
	Title         string `json:"title"`
	Pages         string `json:"pages"`
	CallNumber    string `json:"callNumber"`
	Author        string `json:"author"`
	Year          string `json:"published"`
	Location      string `json:"location"`
	SourceURL     string `json:"link"`
	Description   string `json:"description"`
}

func (svc *serviceContext) submitRequest(c *gin.Context) {
	customerID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	log.Printf("INFO: order request submitted by customer id %d", customerID)

	var req submission
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("ERROR: unable to parse order submission payload: %s", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// NOTE: intended use and user are needed for the confirmation email
	log.Printf("INFO: lookup intended use %d", req.IntendedUseID)
	var useRec intendedUse
	dbErr := svc.GDB.Find(&useRec, req.IntendedUseID).Error
	if dbErr != nil {
		log.Printf("ERROR: unable to get intended use %d: %s", req.IntendedUseID, dbErr.Error())
		c.String(http.StatusInternalServerError, dbErr.Error())
		return
	}

	log.Printf("INFO: lookup customer %d", customerID)
	var user customer
	dbErr = svc.GDB.Preload("AcademicStatus").Preload("Addresses").Find(&user, customerID).Error
	if dbErr != nil {
		log.Printf("ERROR: unable to get customer %d: %s", req.IntendedUseID, dbErr.Error())
		c.String(http.StatusInternalServerError, dbErr.Error())
		return
	}

	log.Printf("INFO: create order and order items")
	newOrder := order{CustomerID: customerID, DateDue: req.DateDue, SpecialInstructions: req.SpecialInstructions,
		DateRequestSubmitted: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(), OrderStatus: "requested"}
	addErr := svc.GDB.Create(&newOrder).Error
	if addErr != nil {
		log.Printf("ERROR: unable to create order: %s", addErr.Error())
		c.String(http.StatusInternalServerError, addErr.Error())
		return
	}

	log.Printf("INFO: add %d items to order", len(req.Items))
	for _, item := range req.Items {
		item.OrderID = newOrder.ID
		item.IntendedUseID = useRec.ID
		itemErr := svc.GDB.Create(&item).Error
		if itemErr != nil {
			log.Printf("ERROR: unable to add items to order: %s", itemErr.Error())
			c.String(http.StatusInternalServerError, itemErr.Error())
			return
		}
	}
	log.Printf("INFO: all items created")

	log.Printf("INFO: generate confirmation email for %s", user.Email)
	var emailData struct {
		OrderID             int64
		DueDate             time.Time
		SpecialInstructions string
		User                customer
		IntendedUse         intendedUse
		Items               []orderItem
	}
	emailData.OrderID = newOrder.ID
	emailData.User = user
	emailData.DueDate = req.DateDue
	emailData.IntendedUse = useRec
	emailData.SpecialInstructions = req.SpecialInstructions
	emailData.Items = req.Items
	var renderedEmail bytes.Buffer
	err = svc.ConfirmTemplate.Execute(&renderedEmail, emailData)
	if err != nil {
		log.Printf("ERROR: unable to generate confirmation email for %s: %s", user.Email, err.Error())
	} else {
		to := make([]string, 0)
		to = append(to, user.Email)
		sub := fmt.Sprintf("UVA Digital Production Group - Request #%d Confirmation", newOrder.ID)
		eRequest := emailRequest{Subject: sub, To: to, From: svc.SMTP.Sender, CC: svc.SMTP.Sender, Body: renderedEmail.String()}
		sendErr := svc.SendEmail(&eRequest)
		if sendErr != nil {
			log.Printf("ERROR: Unable to send confirm email to %s: %s", user.Email, sendErr.Error())
		}
	}

	c.String(http.StatusOK, fmt.Sprintf("%d", newOrder.ID))
}
