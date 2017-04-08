package main

import (
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Users DB Model
type User struct {
	ID           int            `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserName     string         `gorm:"not null;size:64" form:"userName" json:"userName"`
	Password     []byte         `gorm:"not null" form:"passWord" json:"passWord"`
	Email        string         `gorm:"not null;size:255" form:"email" json:"email"`
	FirstName    string         `gorm:"not null;size:64" form:"firstName" json:"firstName"`
	LastName     string         `gorm:"not null;size:64" form:"lastName" json:"lastName"`
	Created      time.Time      `gorm:"not null" form:"created" json:"created"`
	Updated      time.Time      `gorm:"not null" form:"updated" json:"updated"`
	Deleted      NullTime       `form:"deleted" json:"deleted"`
	Achievements []Achievement `gorm:"many2many:user_achievements;" form:"achievements" json:"achievements"`
	Reviews      []Review      `form:"reviews" json:"reviews"`
}

// PostUser creates a User
func PostUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	if user.FirstName != "" && user.LastName != "" && user.UserName != "" && user.Password != nil && user.Email != "" {
		user.Created = time.Now()
		user.Updated = time.Now()
		DB.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// GetUsers gets all Users
func GetUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)

	c.JSON(200, users)
}

// GetUser gets a User
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User

	DB.First(&user, id)
	if user.ID != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "User #" + id + " not found"})
	}
}

// UpdateUser updates a User
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	DB.First(&user, id)

	if user.FirstName != "" && user.LastName != "" && user.UserName != "" &&
		user.Password != nil && user.Email != "" {
		if user.ID != 0 {
			var newUser User
			c.Bind(&newUser)

			result := User{
				ID:        user.ID,
				UserName:  newUser.UserName,
				Password:  newUser.Password,
				Email:     newUser.Email,
				FirstName: newUser.FirstName,
				LastName:  newUser.LastName,
				Updated:   time.Now(),
			}

			DB.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "User #" + id + " not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

// DeleteUser soft deletes a user by setting the deleted date
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	DB.First(&user, id)

	if user.ID != 0 {
		var newUser User = user;
		c.Bind(&newUser)
		newUser.Deleted = NullTime{Time: time.Now(), Valid: true}

		DB.Save(&newUser)
		c.JSON(200, gin.H{"success": User{ID: user.ID, Deleted: user.Deleted}})
	} else {
		c.JSON(404, gin.H{"error": "User #" + id + " not found"})
	}
}

// OptionsUser allows DELETE, POST and PUT to come through
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
