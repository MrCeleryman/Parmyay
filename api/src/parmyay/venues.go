package parmyay

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

// Venue DB Model
type Venue struct {
	ID        int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	VenueName string    `gorm:"not null;size:255" form:"venueName" json:"venueName"`
	Address   string    `gorm:"not null" form:"address" json:"address"`
	Latitude  float64   `gorm:"not null" form:"latitude" json:"latitude"`
	Longitude float64   `gorm:"not null" form:"longitude" json:"longitude"`
	Created   time.Time `gorm:"not null" form:"created" json:"created"`
	Updated   time.Time `gorm:"not null" form:"updated" json:"updated"`
	Deleted   NullTime  `form:"deleted" json:"deleted"`
	Reviews   []Review  `form:"reviews" json:"reviews"`
}

// PostVenue creates a Venue
func PostVenue(c *gin.Context) {
	var venue Venue
	c.Bind(&venue)

	if venue.VenueName != "" && venue.Address != "" && !math.IsNaN(venue.Latitude) && !math.IsNaN(venue.Longitude) {
		venue.Created = time.Now()
		venue.Updated = time.Now()
		DB.Create(&venue)
		c.JSON(201, gin.H{"success": venue})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// GetVenues gets all Venues
func GetVenues(c *gin.Context) {
	var venues []Venue
	DB.Find(&venues)
	c.JSON(200, venues)
}

// GetVenue gets a Venue
func GetVenue(c *gin.Context) {
	id := c.Params.ByName("id")
	var venue Venue
	DB.First(&venue, id)

	if venue.ID != 0 {
		c.JSON(200, venue)
	} else {
		c.JSON(404, gin.H{"error": "Venue #" + id + " not found"})
	}
}

// UpdateVenue updates a Venue
func UpdateVenue(c *gin.Context) {
	id := c.Params.ByName("id")
	var venue Venue
	DB.First(&venue, id)

	if venue.VenueName != "" && venue.Address != "" && !math.IsNaN(venue.Latitude) && !math.IsNaN(venue.Longitude) {
		if venue.ID != 0 {
			var newVenue Venue
			c.Bind(&newVenue)

			result := Venue{
				ID:        venue.ID,
				VenueName: newVenue.VenueName,
				Address:   newVenue.Address,
				Latitude:  newVenue.Latitude,
				Longitude: newVenue.Longitude,
				Updated:   time.Now(),
			}

			DB.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Venue not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

// DeleteVenue soft deletes a venue by setting the deleted date
func DeleteVenue(c *gin.Context) {
	id := c.Params.ByName("id")
	var venue Venue
	DB.First(&venue, id)

	if venue.ID != 0 {
		var newVenue Venue
		c.Bind(&newVenue)
		venue.Deleted = NullTime{Time: getNow(), Valid: true}

		DB.Save(&venue)
		c.JSON(200, gin.H{"success": venue})
	} else {
		c.JSON(404, gin.H{"error": "Venue #" + id + " not found"})
	}
}

// OptionsVenue allows DELETE, POST and PUT to come through
func OptionsVenue(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
