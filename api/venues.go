package main

import (
	"time"

	"math"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Venues struct {
	ID        int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	VenueName string    `gorm:"not null;size:255" form:"venueName" json:"venueName"`
	Address   string    `gorm:"not null" form:"address" json:"address"`
	Latitude  float64   `gorm:"not null" form:"latitude" json:"latitude"`
	Longitude float64   `gorm:"not null" form:"longitude" json:"longitude"`
	Created   time.Time `gorm:"not null" form:"created" json:"created"`
	Updated   time.Time `gorm:"not null" form:"updated" json:"updated"`
	Deleted   NullTime  `form:"deleted" json:"deleted"`
	Reviews   []Reviews `form:"reviews" json:"reviews"`
}

func PostVenue(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var venue Venues
	c.Bind(&venue)

	if venue.VenueName != "" && venue.Address != "" && !math.IsNaN(venue.Latitude) && !math.IsNaN(venue.Longitude) {
		venue.Created = time.Now()
		venue.Updated = time.Now()
		db.Create(&venue)
		c.JSON(201, gin.H{"success": venue})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func GetVenues(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var venues []Venues
	db.Find(&venues)
	c.JSON(200, venues)
}

func GetVenue(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var venue Venues
	db.First(&venue, id)

	if venue.ID != 0 {
		c.JSON(200, venue)
	} else {
		c.JSON(404, gin.H{"error": "Venue #" + id + " not found"})
	}
}

func UpdateVenue(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var venue Venues
	db.First(&venue, id)

	if venue.VenueName != "" && venue.Address != "" && !math.IsNaN(venue.Latitude) && !math.IsNaN(venue.Longitude) {
		if venue.ID != 0 {
			var newVenue Venues
			c.Bind(&newVenue)

			result := Venues{
				ID:        venue.ID,
				VenueName: newVenue.VenueName,
				Address:   newVenue.Address,
				Latitude:  newVenue.Latitude,
				Longitude: newVenue.Longitude,
				Updated:   time.Now(),
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Venue not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

func DeleteVenue(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var venue Venues
	db.First(&venue, id)

	if venue.ID != 0 {
		var newVenue Venues
		c.Bind(&newVenue)

		result := Venues{
			ID:      venue.ID,
			Deleted: NullTime{Time: time.Now(), Valid: true},
		}

		db.Save(&result)
		c.JSON(200, gin.H{"success": result})
	} else {
		c.JSON(404, gin.H{"error": "Venue #" + id + " not found"})
	}
}

func OptionsVenue(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
