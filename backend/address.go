package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type address struct {
	ID              int64     `json:"id"`
	AddressableID   int64     `json:"-"`
	AddressableType string    `json:"-"`
	AddressType     string    `json:"addressType"` // primary or billable_address
	Address1        string    `json:"address1" gorm:"column:address_1"`
	Address2        string    `json:"address2" gorm:"column:address_2"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	PostCode        string    `json:"zip"`
	Country         string    `json:"country"`
	Phone           string    `json:"phone"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (svc serviceContext) getUserAddress(c *gin.Context) {
	uid := c.Param("id")
	log.Printf("INFO: get addresses for user id %s", uid)
	var addresses []address
	err := svc.GDB.Order("address_type desc").Where("addressable_id=? and addressable_type=?", uid, "Customer").Find(&addresses).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("INFO: no address found for user %s", uid)
			c.String(http.StatusNotFound, "no address")
		} else {
			log.Printf("ERROR: unable to get address user %s: %s", uid, err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, addresses)
}

func (svc *serviceContext) updateUserAddress(c *gin.Context) {
	uidStr := c.Param("id")
	var addr address
	err := c.ShouldBindJSON(&addr)
	if err != nil {
		log.Printf("ERROR: unable to parse address update payload: %s", err.Error())
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		log.Printf("ERROR: invalid user id %s: %s", uidStr, err.Error())
		c.String(http.StatusBadRequest, "invlalid user id")
		return
	}

	log.Printf("INFO: user id %d requests address update: %+v", uid, addr)
	if addr.ID > 0 {
		log.Printf("INFO: updating user %s address %d", uidStr, addr.ID)
		addr.UpdatedAt = time.Now()
		upErr := svc.GDB.Model(&addr).Omit("created_at", "addressable_id", "addressable_type").Updates(addr).Error
		if upErr != nil {
			log.Printf("ERROR: unable to update address for customer %s: %s", uidStr, upErr.Error())
			c.String(http.StatusInternalServerError, upErr.Error())
		} else {
			c.JSON(http.StatusOK, addr)
		}
		return
	}

	log.Printf("INFO: create new address for user %s", uidStr)
	addr.CreatedAt = time.Now()
	addr.AddressableType = "Customer"
	addr.AddressableID = uid
	addErr := svc.GDB.Create(&addr).Error
	if addErr != nil {
		log.Printf("ERROR: unable to add address for customer %s: %s", uidStr, addErr.Error())
		c.String(http.StatusInternalServerError, addErr.Error())
		return
	}

	c.JSON(http.StatusOK, addr)
}
