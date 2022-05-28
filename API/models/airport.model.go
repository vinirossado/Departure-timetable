package models

type Airport struct {
	Name            string `json:"name"`
	IATACode        string `json:"iata_code"`
	ICAOCode        string `json:"icao_code"`
	Lat             string `json:"lat"`
	Lng             string `json:"lng"`
	Alt             string `json:"alt"`
	City            string `json:"city"`
	CityCode        string `json:"city_code"`
	UnLocode        string `json:"un_locode"`
	Timezone        string `json:"timezone"`
	CountryCode     string `json:"country_code"`
	Phone           string `json:"phone"`
	Website         string `json:"website"`
	Facebook        string `json:"facebook"`
	Instagram       string `json:"instagram"`
	Linkedin        string `json:"linkedin"`
	Twitter         string `json:"twitter"`
	Runways         string `json:"runways"`
	Departures      string `json:"departures"`
	Connections     string `json:"connections"`
	IsMajor         string `json:"is_major"`
	IsInternational string `json:"is_international"`
	Slug            string `json:"slug"`
}
