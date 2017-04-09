package parmyay

import (
	"testing"
	"os"
)
func TestMain(m *testing.M) {
    os.Setenv("DB_NAME", ":memory:")
    InitDb();
    DB.LogMode(false)
    
    DB.Create(Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House"})
    DB.Create(Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House"})

    DB.Create(Achievement{ID: 1, Achievement: "Reviewed first Parmy!"})
    DB.Create(Achievement{ID: 2, Achievement: "Ate first Parmy!"})

    DB.Create(Review{ID: 1, Notes: "It was prety good", UserID: 1, VenueID: 1, Rating: 10})
    DB.Create(Review{ID: 2, Notes: "Too much water", UserID: 1, VenueID: 1, Rating: 7.8})

    DB.Create(User{ID: 1, FirstName: "Daniel", LastName: "Mitchell", UserName: "DMitch", Email: "d@gmail.com", Password: []byte("Daniel")})
    DB.Create(User{ID: 2, FirstName: "Jerry", LastName: "Seinfeld", UserName: "DarkLordDD", Email: "bee@movie.com", Password: []byte("dfdfgdfgdfg")})

    os.Exit(m.Run());
}