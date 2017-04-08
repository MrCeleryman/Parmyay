package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Reviews DB Model
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

	if review.Rating != 0 && review.Notes != "" && review.UserID != 0 && review.VenueID != 0 {

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
		review.Created = time.Now()
		review.Updated = time.Now()
		DB.Create(&review)
		c.JSON(201, gin.H{"success": review})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
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
	var review Review
	DB.First(&review, id)

	if review.Rating != 0 && review.Notes != "" {
		if review.ID != 0 {
			var newReview Review
			c.Bind(&newReview)

			result := Review{
				ID:      review.ID,
				Rating:  newReview.Rating,
				Notes:   newReview.Notes,
				Updated: time.Now(),
			}

			DB.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Review #" + id + " not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

// DeleteReview soft deletes a review by setting the deleted date
func DeleteReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review Review
	DB.First(&review, id)

	if review.ID != 0 {
		var newReview Review
		c.Bind(&newReview)

		result := Review{
			ID:      review.ID,
			ValidTo: NullTime{Time: time.Now(), Valid: true},
		}

		DB.Save(&result)
		c.JSON(200, gin.H{"success": result})
	} else {
		c.JSON(404, gin.H{"error": "Review #" + id + " not found"})
	}
}

// OptionsReview allows DELETE, POST and PUT to come through
func OptionsReview(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
