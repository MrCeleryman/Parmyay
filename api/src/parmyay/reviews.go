package parmyay

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Review DB Model
type Review struct {
	ID      int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserID  int       `gorm:"index" form:"userId" json:"userId"`
	VenueID int       `gorm:"index" form:"venueId" json:"venueId"`
	Rating  float64   `gorm:"not null" form:"rating" json:"rating"`
	Notes   string    `gorm:"not null" form:"notes" json:"notes"`
	Created time.Time `gorm:"not null" form:"created" json:"created"`
	Updated time.Time `gorm:"not null" form:"updated" json:"updated"`
	ValidTo NullTime  `form:"deleted" json:"deleted"`
}

// PostReview creates a review
func PostReview(c *gin.Context) {
	var review Review
	c.Bind(&review)

	if review.Notes != "" && review.UserID != 0 && review.VenueID != 0 {

		userID := review.UserID
		venueID := review.VenueID
		var user User
		DB.First(&user, userID)
		if user.ID == 0 {
			c.JSON(404, gin.H{"error": "User #" + strconv.Itoa(userID) + " not found"})
			return
		}
		var venue Venue
		DB.First(&venue, venueID)
		if venue.ID == 0 {
			c.JSON(404, gin.H{"error": "Venue #" + strconv.Itoa(venueID) + " not found"})
			return
		}
		review.Created = getNow()
		review.Updated = getNow()
		DB.Create(&review)
		c.JSON(201, gin.H{"success": review})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty or UserID/VenueID are 0"})
	}
}

// GetReviews gets all reviews
func GetReviews(c *gin.Context) {
	var reviews []Review
	DB.Find(&reviews)
	c.JSON(200, reviews)
}

// GetReview gets a review
func GetReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review Review
	DB.First(&review, id)

	if review.ID != 0 {
		c.JSON(200, review)
	} else {
		c.JSON(404, gin.H{"error": "Review #" + id + " not found"})
	}
}

// UpdateReview updates a review
func UpdateReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var existingReview Review
	DB.First(&existingReview, id)

	var newReview Review
	c.Bind(&newReview)

	if newReview.Notes != "" {
		if existingReview.ID != 0 {
			newReview.Updated = getNow()
			newReview.ID = existingReview.ID
			newReview.ValidTo = existingReview.ValidTo
			newReview.Created = existingReview.Created
			newReview.VenueID = existingReview.VenueID
			newReview.UserID = existingReview.UserID

			DB.Model(&existingReview).Updates(newReview)

			c.JSON(200, gin.H{"success": newReview})
		} else {
			c.JSON(404, gin.H{"error": "Review #" + id + " not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "Notes are empty"})
	}
}

// DeleteReview soft deletes a review by setting the deleted date
func DeleteReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review Review
	DB.First(&review, id)

	if review.ID != 0 {
		var newReview = review
		c.Bind(&newReview)
		newReview.ValidTo = NullTime{Time: getNow(), Valid: true}

		DB.Save(&newReview)
		c.JSON(200, gin.H{"success": newReview})
	} else {
		c.JSON(404, gin.H{"error": "Review #" + id + " not found"})
	}
}
