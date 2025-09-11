package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type academicStatus struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type intendedUse struct {
	ID                    int64
	Description           string
	DeliverableFormat     string
	DeliverableResolution string
}

type constantInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (svc *serviceContext) getConstants(c *gin.Context) {
	log.Printf("request for constant values")
	var academicStatuses []academicStatus
	dbResp := svc.GDB.Order("id asc").Find(&academicStatuses)
	if dbResp.Error != nil {
		log.Printf("ERROR: unable to get academic statuses: %s", dbResp.Error.Error())
		c.String(http.StatusInternalServerError, dbResp.Error.Error())
		return
	}

	var uses []intendedUse
	dbResp = svc.GDB.Where("is_approved = ? and id != ? and id != ?", 1, 101, 110).Order("description asc").Find(&uses)
	if dbResp.Error != nil {
		log.Printf("ERROR: unable to get intended uses: %s", dbResp.Error.Error())
		c.String(http.StatusInternalServerError, dbResp.Error.Error())
		return
	}

	out := make(map[string]interface{})
	out["academicStatus"] = academicStatuses
	iu := make([]constantInfo, 0)
	for _, use := range uses {
		u := constantInfo{ID: use.ID}
		switch use.DeliverableResolution {
		case "":
			u.Name = fmt.Sprintf("%s : %s", use.Description, use.DeliverableFormat)
		case "Highest Possible":
			u.Name = fmt.Sprintf("%s : %s resolution %s", use.Description, use.DeliverableResolution, use.DeliverableFormat)
		default:
			u.Name = fmt.Sprintf("%s : %s dpi %s", use.Description, use.DeliverableResolution, use.DeliverableFormat)
		}
		iu = append(iu, u)
	}
	out["intendedUse"] = iu

	c.JSON(http.StatusOK, out)
}
