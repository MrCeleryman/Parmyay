package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Achievements struct {
	ID          int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Achievement string `gorm:"not null" form:"achievement" json:"achievement"`
}

func PostAchievement(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var achievement Achievements
	c.Bind(&achievement)
	if achievement.Achievement != "" {
		db.Create(&achievement)
		c.JSON(201, gin.H{"success": achievement})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func GetAchievements(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var achievements []Achievements
	db.Find(&achievements)
	c.JSON(200, achievements)
}

func GetAchievement(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var achievement Achievements
	db.First(&achievement, id)

	if achievement.ID != 0 {
		c.JSON(200, achievement)
	} else {
		c.JSON(404, gin.H{"error": "Achievement #" + id + " not found"})
	}
}

func UpdateAchievement(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var achievement Achievements
	db.First(&achievement, id)

	if achievement.Achievement != "" {
		if achievement.ID != 0 {
			var newAchievement Achievements
			c.Bind(&newAchievement)

			result := Achievements{
				ID:          achievement.ID,
				Achievement: newAchievement.Achievement,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": "Achievement not found"})
		}

	} else {
		c.JSON(422, gin.H{"error": "One or more of the fields are empty"})
	}
}

func DeleteAchievement(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var achievement Achievements
	db.First(&achievement, id)

	if achievement.ID != 0 {
		db.Delete(&achievement)
		c.JSON(200, gin.H{"success": "Achievement #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Achievement not found"})
	}
}

func OptionsAchievement(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
