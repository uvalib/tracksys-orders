package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type addressRec struct {
	ID        int64          `db:"id"`
	UserID    int64          `db:"addressable_id"`
	Type      string         `db:"address_type"` // primary or billable_address
	Address1  sql.NullString `db:"address_1"`
	Address2  sql.NullString `db:"address_2"`
	City      sql.NullString `db:"city"`
	State     sql.NullString `db:"state"`
	Zip       sql.NullString `db:"post_code"`
	Country   sql.NullString `db:"country"`
	Phone     sql.NullString `db:"phone"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u addressRec) TableName() string {
	return "addresses"
}

type addressJSON struct {
	ID       int64  `json:"id"`
	Type     string `json:"addressType"` // primary or billable_address
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
}

func (svc serviceContext) getUserAddress(c *gin.Context) {
	uid := c.Param("id")
	aType := c.Query("type")
	if aType == "" {
		aType = "primary"
	} else if aType == "billing" {
		aType = "billable_address"
	}
	log.Printf("INFO: get %s address for user id %s", aType, uid)

	q := svc.DB.NewQuery("select * from addresses where addressable_id={:id} and addressable_type={:at} and address_type={:t}")
	q.Bind(dbx.Params{"id": uid})
	q.Bind(dbx.Params{"at": "Customer"})
	q.Bind(dbx.Params{"t": aType})
	var addr addressRec
	err := q.One(&addr)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("INFO: no %s address found for user %s", aType, uid)
			c.String(http.StatusNotFound, "no address")
		} else {
			log.Printf("ERROR: unable to get %s address user %s: %s", aType, uid, err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	out := addressJSON{ID: addr.ID, Type: addr.Type}
	if addr.Address1.Valid {
		out.Address1 = addr.Address1.String
	}
	if addr.Address2.Valid {
		out.Address2 = addr.Address2.String
	}
	if addr.City.Valid {
		out.City = addr.City.String
	}
	if addr.State.Valid {
		out.State = addr.State.String
	}
	if addr.Zip.Valid {
		out.Zip = addr.Zip.String
	}
	if addr.Country.Valid {
		out.Country = addr.Country.String
	}
	if addr.Phone.Valid {
		out.Phone = addr.Phone.String
	}
	c.JSON(http.StatusOK, out)
}
