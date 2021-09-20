package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
}

type submission struct {
	UserID  int64       `json:"userID"`
	Request requestInfo `json:"request"`
	Items   []itemInfo  `json:"items"`
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

	order := orderRec{CustomeerID: req.UserID, DateDue: req.Request.DueDate, Instructions: req.Request.Instructions,
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

	// TODO send email

	log.Printf("INFO: all items created")
	c.String(http.StatusOK, "ok")
}
