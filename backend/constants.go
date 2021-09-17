package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type constantInfo struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (svc *serviceContext) getConstants(c *gin.Context) {
	log.Printf("request for constant values")
	q := svc.DB.NewQuery("select id,name from academic_statuses order by id asc")
	var academicStatuses []constantInfo
	err := q.All(&academicStatuses)
	if err != nil {
		log.Printf("ERROR: unable to get academic statuses: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	q = svc.DB.NewQuery("select id,description as name from intended_uses order by id asc")
	var uses []constantInfo
	err = q.All(&uses)
	if err != nil {
		log.Printf("ERROR: unable to get intended uses: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	out := make(map[string][]constantInfo)
	out["academicStatus"] = academicStatuses
	out["intendedUse"] = uses

	c.JSON(http.StatusOK, out)
}
