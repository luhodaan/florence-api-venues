package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Venue struct {
	ID                int     `db:"id" json:"id"`
	VenueName         string  `db:"venue_name" json:"venue_name"`
	VenueType         string  `db:"venue_type" json:"venue_type"`
	MainBus           string  `db:"main_bus" json:"main_bus"`
	VenueAddOne       string  `db:"venue_add_one" json:"venue_add_one"`
	VenueAddTwo       string  `db:"venue_add_two" json:"venue_add_two"`
	VenuePcode        string  `db:"venue_pcode" json:"venue_pcode"`
	VenueWeb          string  `db:"venue_web" json:"venue_web"`
	VenuePhone        string  `db:"venue_phone" json:"venue_phone"`
	VenueEmail        string  `db:"venue_email" json:"venue_email"`
	VenueFB           string  `db:"venue_fb" json:"venue_fb"`
	VenueTW           string  `db:"venue_tw" json:"venue_tw"`
	OrgOwn            string  `db:"org_own" json:"org_own"`
	Manager           string  `db:"mger" json:"mger"`
	ManagerEmail      string  `db:"mger_email" json:"mger_email"`
	ManagerPhone      string  `db:"mger_phone" json:"mger_phone"`
	OnlineContactForm bool    `db:"online_contact_form" json:"online_contact_form"`
	FollowUp          bool    `db:"follow_up" json:"follow_up"`
	VenueCity         string  `db:"venue_city" json:"venue_city"`
	VenueCountry      string  `db:"venue_country" json:"venue_country"`
	Notes             string  `db:"notes" json:"notes"`
	StdSignOff        bool    `db:"std_sign_off" json:"std_sign_off"`
	ResSignOff        bool    `db:"res_sign_off" json:"res_sign_off"`
	SKickID           string  `db:"s_kick_id" json:"s_kick_id"`
	Latitude          float64 `db:"latitude" json:"latitude"`
	Longitude         float64 `db:"longitude" json:"longitude"`
	Constituency      string  `db:"constituency" json:"constituency"`
	WdayOpening       string  `db:"wday_opening" json:"wday_opening"`
	WdayClosing       string  `db:"wday_closing" json:"wday_closing"`
	WkdOpening        string  `db:"wkd_opening" json:"wkd_opening"`
	WkdClosing        string  `db:"wkd_closing" json:"wkd_closing"`
	ShowOpening       string  `db:"show_opening" json:"show_opening"`
	ShowClosing       string  `db:"show_closing" json:"show_closing"`
	Variations        string  `db:"variations" json:"variations"`
	Promoters         string  `db:"promoters" json:"promoters"`
	VenueCapacity     int     `db:"venue_capacity" json:"venue_capacity"`
	OutdoorSpace      bool    `db:"outdoor_space" json:"outdoor_space"`
	TypeOutdoorSpace  string  `db:"type_outdoor_space" json:"type_outdoor_space"`
	Instagram         string  `db:"instagram" json:"instagram"`
	PRSType           string  `db:"prs_type" json:"prs_type"`
	PRSVenueCat       string  `db:"prs_venue_cat" json:"prs_venue_cat"`
	YearVenueOpened   int     `db:"year_venue_opened" json:"year_venue_opened"`
	YearVenueClosed   int     `db:"year_venue_closed" json:"year_venue_closed"`
	AdditionalNotes   string  `db:"notes_additional" json:"notes_additional"`
	Quartiere         string  `db:"quartiere" json:"quartiere"`
	VenueCapacityAll  string  `db:"venue_capacity_all" json:"venue_capacity_all"`
	Nr                int     `db:"nr" json:"nr"`
	Key               string  `db:"key" json:"key"`
}

func main() {

	var db *sqlx.DB
	var err error

	connsStr := "postgresql://luho:luho@localhost:5432/myapp?sslmode=disable"

	db, err = sqlx.Open("postgres", connsStr)
	//defer db.Close()

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")

	var venues []Venue
	venues, err = getVenuesFromDB(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(venues)

}

func getVenuesFromDB(db *sqlx.DB) ([]Venue, error) {
	var venues []Venue
	err := db.Select(&venues, "SELECT * FROM florence_open_data")
	return venues, err
}
