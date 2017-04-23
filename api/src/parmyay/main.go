package parmyay

import (
	"database/sql/driver"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Register sqlite3 driver for gorm
	_ "github.com/mattn/go-sqlite3"
)

// ErrorResult Model for a server error message
type ErrorResult struct {
	Error string `json:"error"`
}

// SuccessResult Model for a server success message
type SuccessResult map[string]interface{}

var (
	// DB global to handle the DB connection
	DB *gorm.DB
)

var getNow = time.Now

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

// InitDb initializes the Database
func InitDb() {
	var err error

	if DB, err = gorm.Open("sqlite3", os.Getenv("DB_NAME")); err != nil {
		panic(fmt.Sprintf("Error when connecting to %s: err=%+v", os.Getenv("DB_NAME"), err))
	}
	DB.LogMode(false)
	if !DB.HasTable(&Achievement{}) {
		DB.CreateTable(&Achievement{})
	}
	if !DB.HasTable(&Review{}) {
		DB.CreateTable(&Review{})
	}
	if !DB.HasTable(&User{}) {
		DB.Model(&User{}).Related(&User{})
		DB.CreateTable(&User{})
	}
	if !DB.HasTable(&Venue{}) {
		DB.CreateTable(&Venue{})
	}
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
	}

	venues := v1.Group("venues")
	{
		venues.POST("/", PostVenue)
		venues.GET("/", GetVenues)
		venues.GET("/:id", GetVenue)
		venues.PUT("/:id", UpdateVenue)
		venues.PATCH("/:id", DeleteVenue)
	}

	reviews := v1.Group("reviews")
	{
		reviews.POST("/", PostReview)
		reviews.GET("/", GetReviews)
		reviews.GET("/:id", GetReview)
		reviews.PUT("/:id", UpdateReview)
		reviews.PATCH("/:id", DeleteReview)
	}

	achievements := v1.Group("achievements")
	{
		achievements.POST("/", PostAchievement)
		achievements.GET("/", GetAchievements)
		achievements.GET("/:id", GetAchievement)
		achievements.PUT("/:id", UpdateAchievement)
		achievements.DELETE("/:id", DeleteAchievement)
	}
	return router
}
