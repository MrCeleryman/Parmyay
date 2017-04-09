package parmyay

import (
	"github.com/gin-gonic/gin"
 	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Achievement DB Model
type Achievement struct {
	ID          int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Achievement string `gorm:"not null" form:"achievement" json:"achievement"`
}

// PostAchievement Create an Achievement
func PostAchievement(c *gin.Context) {
	var achievement Achievement
	c.Bind(&achievement)
	if IsInt(achievement.Achievement) {
		c.JSON(400, gin.H{"error": "No Numbers allowed"})
	} else if achievement.Achievement != "" {
		DB.Create(&achievement)
		c.JSON(201, gin.H{"success": achievement})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// GetAchievement Gets all Achievements
func GetAchievements(c *gin.Context) {
	var Achievement []Achievement
	DB.Find(&Achievement)
	c.JSON(200, Achievement)
}

// GetAchievement Gets an achievement
func GetAchievement(c *gin.Context) {
	id := c.Params.ByName("id")
	var achievement Achievement
	DB.First(&achievement, id)

	if achievement.ID != 0 {
		c.JSON(200, achievement)
	} else {
		c.JSON(404, gin.H{"error": "Achievement #" + id + " not found"})
	}
}

// UpdateAchievement updates an Achievement
func UpdateAchievement(c *gin.Context) {
	id := c.Params.ByName("id")
	var achievement Achievement
	DB.First(&achievement, id)

	if achievement.ID != 0 {
		var newAchievement Achievement
		c.Bind(&newAchievement)
		if newAchievement.Achievement != "" {
			result := Achievement{
				ID:          achievement.ID,
				Achievement: newAchievement.Achievement,
			}
			DB.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(422, gin.H{"error": "The Achievement field is empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "Achievement not found"})
	}
}

// DeleteAchievement deletes an achievement
func DeleteAchievement(c *gin.Context) {

	id := c.Params.ByName("id")
	var achievement Achievement
	DB.First(&achievement, id)

	if achievement.ID != 0 {
		DB.Delete(&achievement)
		c.Status(204)
	} else {
		c.JSON(404, gin.H{"error": "Achievement not found"})
	}
}

// OptionsAchievement allows DELETE, POST and PUT to come through
func OptionsAchievement(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
