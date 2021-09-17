package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// customerRec is the model representing the customers table. It handles null values.
type customerRec struct {
	ID               int64          `db:"id"`
	FirstName        sql.NullString `db:"first_name"`
	LastName         sql.NullString `db:"last_name"`
	Email            sql.NullString `db:"email"`
	AcademicStatusID int64          `db:"academic_status_id"`
	CreatedAt        time.Time      `db:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u customerRec) TableName() string {
	return "customers"
}

// customerJSON is the JSON data passed to/from the API
type customerJSON struct {
	ID               int64  `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	AcademicStatusID int64  `json:"academicStatusID"`
}

func (svc *serviceContext) getUser(c *gin.Context) {
	cid := c.Param("id")
	log.Printf("INFO: get user data for computing id %s", cid)

	email := fmt.Sprintf("%s@virginia.edu", cid)
	var user customerRec
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

	out := customerJSON{ID: user.ID, AcademicStatusID: user.AcademicStatusID}
	if user.FirstName.Valid {
		out.FirstName = user.FirstName.String
	}
	if user.LastName.Valid {
		out.LastName = user.LastName.String
	}
	if user.Email.Valid {
		out.Email = user.Email.String
	}
	c.JSON(http.StatusOK, out)
}

func (svc *serviceContext) updateUser(c *gin.Context) {
	log.Printf("INFO: update or create customer requested")
	var user customerJSON
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("ERROR: unable to parse user update payload: %s", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if user.Email == "" {
		log.Printf("ERROR: email is required for customer update")
		c.String(http.StatusBadRequest, "email is required")
		return
	}

	log.Printf("INFO: customer update payload: %+v", user)
	custRec := customerRec{ID: user.ID, AcademicStatusID: user.AcademicStatusID, UpdatedAt: time.Now()}
	if user.FirstName != "" {
		custRec.FirstName.String = user.FirstName
		custRec.FirstName.Valid = true
	}
	if user.LastName != "" {
		custRec.LastName.String = user.LastName
		custRec.LastName.Valid = true
	}
	if user.Email != "" {
		custRec.Email.String = user.Email
		custRec.Email.Valid = true
	}

	// First, see if there is an ID with this user. If so, just update
	if user.ID > 0 {
		log.Printf("INFO: update existing customer with ID %d", user.ID)
		upErr := svc.DB.Model(&custRec).Exclude("CreatedAt").Update()
		if upErr != nil {
			log.Printf("ERROR: unable to update customer %d: %s", user.ID, upErr.Error())
			c.String(http.StatusInternalServerError, upErr.Error())
		} else {
			c.JSON(http.StatusOK, user)
		}
		return
	}

	// See if there is an existing customer with the matching email
	q := svc.DB.NewQuery("select id from customers where email={:email}")
	q.Bind(dbx.Params{"email": user.Email})
	var uid int64
	err = q.Row(&uid)
	if err == nil {
		custRec.ID = uid
		log.Printf("INFO: update existing customer with email %s id %d", user.Email, uid)
		upErr := svc.DB.Model(&custRec).Exclude("CreatedAt").Update()
		if upErr != nil {
			log.Printf("ERROR: unable to update customer %d: %s", user.ID, upErr.Error())
			c.String(http.StatusInternalServerError, upErr.Error())
		} else {
			user.ID = uid
			c.JSON(http.StatusOK, user)
		}
		return
	}

	log.Printf("INFO: create new customer %+v", user)
	custRec.CreatedAt = time.Now()
	custRec.ID = 0
	addErr := svc.DB.Model(&custRec).Insert()
	if addErr != nil {
		log.Printf("ERROR: unable to create customer %s: %s", user.Email, addErr.Error())
		c.String(http.StatusInternalServerError, addErr.Error())
		return
	}
	user.ID = custRec.ID
	c.JSON(http.StatusOK, user)

}
