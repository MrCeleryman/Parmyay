package main

import (
	"database/sql/driver"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&Achievements{}) {
		db.CreateTable(&Achievements{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Achievements{})
	}
	if !db.HasTable(&Reviews{}) {
		db.CreateTable(&Reviews{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Reviews{})
	}
	if !db.HasTable(&Users{}) {
		db.Model(&Users{}).Related(&Achievements{}, "Achievements")
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}
	if !db.HasTable(&Venues{}) {
		db.CreateTable(&Venues{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Venues{})
	}
	return db
}

type NullTime struct {
	Time  time.Time `form:"time" json:"time"`
	Valid bool      `form:"valid" json:"valid"` // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func main() {
	//gin.SetMode(gin.ReleaseMode) //For Release
	router := gin.Default()
	router.Use(Cors())

	v1 := router.Group("api/v1")

	users := v1.Group("users")
	{
		users.POST("/", PostUser)
		users.GET("/", GetUsers)
		users.GET("/:id", GetUser)
		users.PUT("/:id", UpdateUser)
		users.PATCH("/:id", DeleteUser)
		users.OPTIONS("/", OptionsUser)
		users.OPTIONS("/:id", OptionsUser)
	}

	venues := v1.Group("venues")
	{
		venues.POST("/", PostVenue)
		venues.GET("/", GetVenues)
		venues.GET("/:id", GetVenue)
		venues.PUT("/:id", UpdateVenue)
		venues.PATCH("/:id", DeleteVenue)
		venues.OPTIONS("/", OptionsVenue)
		venues.OPTIONS("/:id", OptionsVenue)
	}

	reviews := v1.Group("reviews")
	{
		reviews.POST("/", PostReview)
		reviews.GET("/", GetReviews)
		reviews.GET("/:id", GetReview)
		reviews.PUT("/:id", UpdateReview)
		reviews.PATCH("/:id", DeleteReview)
		reviews.OPTIONS("/", OptionsReview)
		reviews.OPTIONS("/:id", OptionsReview)
	}

	achievements := v1.Group("achievements")
	{
		achievements.POST("/", PostAchievement)
		achievements.GET("/", GetAchievements)
		achievements.GET("/:id", GetAchievement)
		achievements.PUT("/:id", UpdateAchievement)
		achievements.DELETE("/:id", DeleteAchievement)
		achievements.OPTIONS("/", OptionsAchievement)
		achievements.OPTIONS("/:id", OptionsAchievement)
	}

	router.Run(":8900")
}
