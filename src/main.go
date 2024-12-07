package main

import (
	_ "github.com/lib/pq"
)

type Venues struct {
	ID                int     `json:"id"`
	VenueName         string  `json:"venue_name"`
	VenueType         string  `json:"venue_type"`
	MainBus           string  `json:"main_bus"`
	VenueAddOne       string  `json:"venue_add_one"`
	VenueAddTwo       string  `json:"venue_add_two"`
	VenuePcode        string  `json:"venue_pcode"`
	VenueWeb          string  `json:"venue_web"`
	VenuePhone        string  `json:"venue_phone"`
	VenueEmail        string  `json:"venue_email"`
	VenueFB           string  `json:"venue_fb"`
	VenueTW           string  `json:"venue_tw"`
	OrgOwn            string  `json:"org_own"`
	Manager           string  `json:"mger"`
	ManagerEmail      string  `json:"mger_email"`
	ManagerPhone      string  `json:"mger_phone"`
	OnlineContactForm bool    `json:"online_contact_form"`
	FollowUp          bool    `json:"follow_up"`
	VenueCity         string  `json:"venue_city"`
	VenueCountry      string  `json:"venue_country"`
	Notes             string  `json:"notes"`
	StdSignOff        bool    `json:"std_sign_off"`
	ResSignOff        bool    `json:"res_sign_off"`
	SKickID           string  `json:"s_kick_id"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	Constituency      string  `json:"constituency"`
	WdayOpening       string  `json:"wday_opening"`
	WdayClosing       string  `json:"wday_closing"`
	WkdOpening        string  `json:"wkd_opening"`
	WkdClosing        string  `json:"wkd_closing"`
	ShowOpening       string  `json:"show_opening"`
	ShowClosing       string  `json:"show_closing"`
	Variations        string  `json:"variations"`
	Promoters         string  `json:"promoters"`
	VenueCapacity     int     `json:"venue_capacity"`
	OutdoorSpace      bool    `json:"outdoor_space"`
	TypeOutdoorSpace  string  `json:"type_outdoor_space"`
	Instagram         string  `json:"instagram"`
	PRSType           string  `json:"prs_type"`
	PRSVenueCat       string  `json:"prs_venue_cat"`
	YearVenueOpened   int     `json:"year_venue_opened"`
	YearVenueClosed   int     `json:"year_venue_closed"`
	AdditionalNotes   string  `json:"notes_additional"`
	Quartiere         string  `json:"quartiere"`
	VenueCapacityAll  string  `json:"venue_capacity_all"`
	Nr                int     `json:"nr"`
	Key               string  `json:"key"`
}
