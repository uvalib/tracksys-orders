package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// customer is the model representing the customers table
type customer struct {
	ID               int64          `json:"id"`
	FirstName        string         `json:"firstName"`
	LastName         string         `json:"lastName"`
	Email            string         `json:"email"`
	AcademicStatusID int64          `json:"academicStatusID"`
	AcademicStatus   academicStatus `gorm:"foreignKey:AcademicStatusID" json:"academicStatus"`
	Addresses        []address      `json:"addresses,omitempty" gorm:"polymorphic:Addressable;polymorphicValue:Customer"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
}

func (svc *serviceContext) getUser(c *gin.Context) {
	cid := c.Param("id")
	log.Printf("INFO: get user data for computing id %s", cid)

	email := fmt.Sprintf("%s@virginia.edu", cid)
	var user customer
	err := svc.GDB.Preload("AcademicStatus").Where("email=?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("INFO: customer %s does not exist", email)
			c.String(http.StatusNotFound, fmt.Sprintf("%s not found", email))
		} else {
			log.Printf("ERROR: query for customer %s failed: %s", email, err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (svc *serviceContext) updateUser(c *gin.Context) {
	log.Printf("INFO: update or create customer requested")
	var user customer
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

	// First, see if there is an ID with this user. If so, just updateeeee
	if user.ID > 0 {
		log.Printf("INFO: update existing customer with ID %d", user.ID)
		user.UpdatedAt = time.Now()
		upErr := svc.GDB.Model(&user).Select("first_name", "last_name", "email", "academic_stattus_id", "updated_at").Updates(user).Error
		if upErr != nil {
			log.Printf("ERROR: unable to update customer %d: %s", user.ID, upErr.Error())
			c.String(http.StatusInternalServerError, upErr.Error())
		} else {
			c.JSON(http.StatusOK, user)
		}
		return
	}

	// See if there is an existing customer with the matching email
	var existUser customer
	err = svc.GDB.Select("id").Where("email=?", user.Email).First(&existUser).Error
	if err == nil {
		user.ID = existUser.ID
		log.Printf("INFO: update existing customer with email %s id %d", user.Email, existUser.ID)
		upErr := svc.GDB.Model(&user).Select("first_name", "last_name", "email", "academic_stattus_id", "updated_at").Updates(user).Error
		if upErr != nil {
			log.Printf("ERROR: unable to update customer %d: %s", user.ID, upErr.Error())
			c.String(http.StatusInternalServerError, upErr.Error())
		} else {
			c.JSON(http.StatusOK, user)
		}
		return
	}

	log.Printf("INFO: create new customer %+v", user)
	user.CreatedAt = time.Now()
	addErr := svc.GDB.Create(&user).Error
	if addErr != nil {
		log.Printf("ERROR: unable to create customer %s: %s", user.Email, addErr.Error())
		c.String(http.StatusInternalServerError, addErr.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}
