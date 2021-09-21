package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type constantInfo struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type intendedUseRec struct {
	ID         int64          `db:"id"`
	Name       string         `db:"description"`
	Format     string         `db:"deliverable_format"`
	Resolution sql.NullString `db:"deliverable_resolution"`
}

// TableName sets the name of the table in the DB that this struct binds to
func (u intendedUseRec) TableName() string {
	return "intended_uses"
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

	q = svc.DB.NewQuery("select id,description, deliverable_format, deliverable_resolution from intended_uses where id != 110 and is_approved=1 order by description asc")
	var uses []intendedUseRec
	err = q.All(&uses)
	if err != nil {
		log.Printf("ERROR: unable to get intended uses: %s", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	out := make(map[string][]constantInfo)
	out["academicStatus"] = academicStatuses
	out["intendedUse"] = make([]constantInfo, 0)
	for _, use := range uses {
		u := constantInfo{ID: use.ID}
		if use.Resolution.Valid {
			if use.Resolution.String == "Highest Possible" {
				u.Name = fmt.Sprintf("%s : %s resolution %s", use.Name, use.Resolution.String, use.Format)
			} else {
				u.Name = fmt.Sprintf("%s : %s dpi %s", use.Name, use.Resolution.String, use.Format)
			}
		} else {
			u.Name = fmt.Sprintf("%s : %s", use.Name, use.Format)
		}
		out["intendedUse"] = append(out["intendedUse"], u)
	}

	c.JSON(http.StatusOK, out)
}
