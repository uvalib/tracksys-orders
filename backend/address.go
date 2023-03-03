package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func (svc *serviceContext) updateUserAddress(c *gin.Context) {
	uidStr := c.Param("id")
	var req struct {
		Primary          address `json:"primary"`
		DifferentBilling bool    `json:"differentBilling"`
		Billing          address `json:"billing"`
	}
	err := c.ShouldBindJSON(&req)
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

	log.Printf("INFO: user id %d requests address update: %+v", uid, req)
	var addresses []address
	addresses = append(addresses, req.Primary)
	if req.DifferentBilling {
		addresses = append(addresses, req.Billing)
	} else {
		log.Printf("INFO: customer has same billing and primary addres; remove prior billing address")
		svc.GDB.Where("address_type=? AND addressable_id=?", "billable_address", uid).Delete(&address{})
	}

	for _, addr := range addresses {
		if addr.ID > 0 {
			log.Printf("INFO: updating user %s %s address %d", uidStr, addr.AddressType, addr.ID)
			addr.UpdatedAt = time.Now()
			upErr := svc.GDB.Model(&addr).Omit("created_at", "addressable_id", "addressable_type").Updates(addr).Error
			if upErr != nil {
				log.Printf("ERROR: unable to update address for customer %s: %s", uidStr, upErr.Error())
				c.String(http.StatusInternalServerError, upErr.Error())
			}
		} else {
			log.Printf("INFO: create new %s address for user %s", addr.AddressType, uidStr)
			addr.CreatedAt = time.Now()
			addr.AddressableType = "Customer"
			addr.AddressableID = uid
			addErr := svc.GDB.Create(&addr).Error
			if addErr != nil {
				log.Printf("ERROR: unable to add address for customer %s: %s", uidStr, addErr.Error())
				c.String(http.StatusInternalServerError, addErr.Error())
				return
			}
		}
	}

	c.JSON(http.StatusOK, svc.getCustomerAddresses(uid))
}
