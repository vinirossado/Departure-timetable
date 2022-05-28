package models

type FlightResponse struct {
	Request  interface{} `json:"request"`
	Response []Flight    `json:"response"`
	Terms    string      `json:"terms"`
}

type Flight struct {
	Hex          string  `json:"hex"`
	AircraftICAO string  `json:"aircraft_icao"`
	Alt          float32 `json:"alt"`
	Dir          float32 `json:"dir"`
	Flag         string  `json:"flag"`
	Lat          float32 `json:"lat"`
	Lng          float32 `json:"lng"`
	RegNumber    string  `json:"reg_number"`
	Speed        float32 `json:"speed"`
	Status       string  `json:"status"`
	Updated      float32 `json:"updated"`
	VSpeed       float32 `json:"v_speed"`
	Squawk       string  `json:"squawk"`
	FlightNumber string  `json:"flight_number"`
	FlightICAO   string  `json:"flight_icao"`
	FlightIATA   string  `json:"flight_iata"`
	DepICAO      string  `json:"dep_icao"`
	DepIata      string  `json:"dep_iata"`
	ArrICAO      string  `json:"arr_icao"`
	ArrIATA      string  `json:"arr_iata"`
	AirlineICAO  string  `json:"airline_icao"`
	AirlineIATA  string  `json:"airline_iata"`
}

// type CityRepository interface {
// 	GetAllCity() ([]models.City, error)
// 	GetCity(uint64) (*models.City, error)
// }
