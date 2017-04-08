package main

import (
	"database/sql/driver"
	"strconv"
	"time"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IsInt checkes whether a given string is an integer
func IsInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

// Cors sets up Cors to allow Cross Origin requests
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// InitDb initializes the Database depending on the test environment variable
func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true)
	if os.Getenv("TEST") == "1" {
		db, err = gorm.Open("sqlite3", ":memory:")
		db.DropTableIfExists(&Achievements{}, &Reviews{}, &Users{}, &Venues{})
		db.LogMode(false)
	}
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&Achievement{}) {
		db.CreateTable(&Achievement{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Achievement{})
	}
	if !db.HasTable(&Reviews{}) {
		db.CreateTable(&Reviews{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Reviews{})
	}
	if !db.HasTable(&Users{}) {
		db.Model(&Users{}).Related(&Achievement{}, "Achievements")
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}
	if !db.HasTable(&Venues{}) {
		db.CreateTable(&Venues{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Venues{})
	}

	if os.Getenv("TEST") == "1" {
		var achievement Achievement
		achievement.ID = 1
		achievement.Achievement = "Reviewed first Parmy!"
		db.Create(&achievement)
		achievement.ID = 2
		achievement.Achievement = "Ate first Parmy!"
		db.Create(&achievement)

		var venue Venue
		venue.ID = 1
		venue.Address = "30 Willy Wonka Way"
		venue.VenueName = "Dans House"
		db.Create(&venue)
		venue.ID = 2
		venue.Address = "42 Wallaby Way, Sydney"
		venue.VenueName = "Dans Old House"
		db.Create(&venue)

		var user User
		user.ID = 1
		user.FirstName = "Daniel"
		user.LastName = "Mitchell"
		user.UserName = "DMitch"
		user.Email = "d@gmail.com"
		user.Password = []byte("Daniel")
		db.Create(&user)
		user.ID = 2
		user.FirstName = "Jerry"
		user.LastName = "Seinfield"
		user.UserName = "Beemovie4lyf"
		user.Email = "beemovie@gmail.com"
		user.Password = []byte("Daniel")
		db.Create(&user)

		var review Review
		review.ID = 1
		review.Notes = "It was pretty good"
		review.UserID = 1
		review.VenueID = 1
		review.Rating = 10
		db.Create(&review)
		review.ID = 2
		review.Notes = "Too much water"
		review.UserID = 1
		review.VenueID = 1
		review.Rating = 7.8
		db.Create(&review)
	}
	return db
}

// NullTime is a model to allow for a nullable time
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

// SetupRouter sets up the Router to route requests to the functions
func SetupRouter(release bool, log bool) *gin.Engine {
	if release == true {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	if log == false {
		router = gin.New()
	}
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
	return router
}

func main() {
	router := SetupRouter(false, true)
	router.Run(":8900")
}
