package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type customerInfo struct {
	ID               int64   `db:"id" json:"id"`
	FirstName        *string `db:"first_name" json:"firstName"`
	LastName         *string `db:"last_name" json:"lastName"`
	Email            *string `db:"email" json:"email"`
	AcademicStatusID int64   `db:"academic_status_id" json:"academicStatusID"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u customerInfo) TableName() string {
	return "customers"
}

func (svc *serviceContext) getUser(c *gin.Context) {
	cid := c.Param("id")
	log.Printf("INFO: get user data for computing id %s", cid)

	email := fmt.Sprintf("%s@virginia.edu", cid)
	var user customerInfo
	q := svc.DB.NewQuery("select id,first_name,last_name,email,academic_status_id from customers where email={:email}")
	q.Bind(dbx.Params{"email": email})
	err := q.One(&user)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("ERROR: query for customer %s failed: %s", email, err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			log.Printf("INFO: customer %s does not exist", email)
			c.String(http.StatusNotFound, fmt.Sprintf("%s not found", email))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
