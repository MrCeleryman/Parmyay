package parmyay

import (
	"testing"
)

func TestGetVenues(t *testing.T) {

	// Test get all when the db is empty
	GetFunc(t, "/api/v1/venues/", 200, []Venue{})

	var v1 = Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House", Longitude: 22.5, Latitude: 50.44, Created: getNow(), Updated: getNow()}
	var v2 = Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House", Longitude: 15.5, Latitude: 10.44, Created: getNow(), Updated: getNow()}
	DB.Create(v1)
	DB.Create(v2)

	// Test get all when there are achievements
	GetFunc(t, "/api/v1/venues/", 200, []Venue{v1, v2})

	// Test cases
	cases := []Case{
		{"/api/v1/venues/1", []byte(``), 200, v1},
		{"/api/v1/venues/2", []byte(``), 200, v2},

		{"/api/v1/venues/0", []byte(``), 404, ErrorResult{"Venue #0 not found"}},
		{"/api/v1/venues/20", []byte(``), 404, ErrorResult{"Venue #20 not found"}},
	}
	for _, c := range cases {
		GetFunc(t, c.query, c.expectedCode, c.expected)
	}
}

func TestDeleteVenue(t *testing.T) {
	var v1 = Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House", Longitude: 22.5, Latitude: 50.44, Created: getNow(), Updated: getNow()}
	var v2 = Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House", Longitude: 15.5, Latitude: 10.44, Created: getNow(), Updated: getNow()}
	DB.Create(v1)
	DB.Create(v2)

	// Test cases
	cases := []Case{
		{"/api/v1/venues/1", []byte(``), 200, SuccessResult{
			"success": Venue{ID: 1, Address: "30 Willy Wonka Way", VenueName: "Dans House", Longitude: 22.5, Latitude: 50.44, Created: getNow(), Updated: getNow(), Deleted: NullTime{Time: getNow(), Valid: true}},
		}},
		{"/api/v1/venues/2", []byte(``), 200, SuccessResult{
			"success": Venue{ID: 2, Address: "42 Wallaby Way, Sydney", VenueName: "Dans Old House", Longitude: 15.5, Latitude: 10.44, Created: getNow(), Updated: getNow(), Deleted: NullTime{Time: getNow(), Valid: true}},
		}},

		{"/api/v1/venues/0", []byte(``), 404, ErrorResult{"Venue #0 not found"}},
		{"/api/v1/venues/20", []byte(``), 404, ErrorResult{"Venue #20 not found"}},
	}
	for _, c := range cases {
		SoftDeleteFunc(t, c.query, c.json, c.expectedCode, c.expected)
	}
}
