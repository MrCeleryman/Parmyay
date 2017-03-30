package main

import (
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	ID           int            `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserName     string         `gorm:"not null;size:64" form:"userName" json:"userName"`
	Password     []byte         `gorm:"not null" form:"passWord" json:"passWord"`
	Email        string         `gorm:"not null;size:255" form:"email" json:"email"`
	FirstName    string         `gorm:"not null;size:64" form:"firstName" json:"firstName"`
	LastName     string         `gorm:"not null;size:64" form:"lastName" json:"lastName"`
	Created      time.Time      `gorm:"not null" form:"created" json:"created"`
	Updated      time.Time      `gorm:"not null" form:"updated" json:"updated"`
	Deleted      NullTime       `form:"deleted" json:"deleted"`
	Achievements []Achievements `gorm:"many2many:user_achievements;" form:"achievements" json:"achievements"`
	Reviews      []Reviews      `form:"reviews" json:"reviews"`
}

func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user Users
	c.Bind(&user)
	fmt.Println(user.Password)
	if user.FirstName != "" && user.LastName != "" && user.UserName != "" && user.Password != nil && user.Email != "" {
		user.Created = time.Now()
		user.Updated = time.Now()
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func GetUsers(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var users []Users
	db.Find(&users)

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users

	db.First(&user, id)
	if user.ID != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "User #" + id + " not found"})
	}
}

func UpdateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	db.First(&user, id)

	if user.FirstName != "" && user.LastName != "" && user.UserName != "" &&
		user.Password != nil && user.Email != "" {
		if user.ID != 0 {
			var newUser Users
			c.Bind(&newUser)

			result := Users{
				ID:        user.ID,
				UserName:  newUser.UserName,
				Password:  newUser.Password,
				Email:     newUser.Email,
				FirstName: newUser.FirstName,
				LastName:  newUser.LastName,
				Updated:   time.Now(),
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "User #" + id + " not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	db.First(&user, id)

	if user.ID != 0 {
		var newUser Users
		c.Bind(&newUser)

		result := Users{
			ID:      user.ID,
			Deleted: NullTime{Time: time.Now(), Valid: true},
		}

		db.Save(&result)
		c.JSON(200, gin.H{"success": result})
	} else {
		c.JSON(404, gin.H{"error": "User #" + id + " not found"})
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
