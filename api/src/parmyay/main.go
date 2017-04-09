package parmyay

import (
	"database/sql/driver"
	"strconv"
	"time"
    "fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
    DB *gorm.DB
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

// InitDb initializes the Database
func InitDb() {
    var err error
    
	if DB, err = gorm.Open("sqlite3", os.Getenv("DB_NAME")); err != nil {
		panic(fmt.Sprintf("Error when connecting to production DB: err=%+v", err))
		DB.LogMode(true)
	}
    
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

//func main() {
//    InitDb();
	//router := SetupRouter(false, true)
	//router.Run(":8900")
//}
