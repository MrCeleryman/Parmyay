package parmyay

import (
    "fmt"
	"testing"
	"os"
	"github.com/jinzhu/gorm"
)
func TestMain(m *testing.M) {
    var err error

	if DB, err = gorm.Open("sqlite3", ":memory:"); err != nil {
        panic(fmt.Sprintf("Error when connecting to test DB: err=%+v", err))
    }

    DB.CreateTable(&Venue{})
    DB.Create(Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House"})
    DB.Create(Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House"})

    DB.CreateTable(&Achievement{})
    DB.Create(Achievement{ID: 1, Achievement: "Reviewed first Parmy!"})
    DB.Create(Achievement{ID: 2, Achievement: "Ate first Parmy!"})

    DB.CreateTable(&Review{})
    DB.Create(Review{ID: 1, Notes: "It was prety good", UserID: 1, VenueID: 1, Rating: 10})
    DB.Create(Review{ID: 2, Notes: "Too much water", UserID: 1, VenueID: 1, Rating: 7.8})

    DB.CreateTable(&User{})
    DB.Create(User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: []byte("Daniel")})
    DB.Create(User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: []byte("dfdfgdfgdfg")})

    os.Exit(m.Run());
}