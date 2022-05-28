
export class TimeTableModelObject{


  //Filtering by ICAO24 Hex address.
  hex!: string;

  //Filtering by Airline ICAO code.
  aircraft_icao!: string;

  // Aircraft elevation for now (meters).
  alt!: number;

  // Aircraft head direction for now;
  dir!: number;

  // ISO 2 country code from Countries DB;
  flag!: string;

  // Aircraft Geo-Latitude for now;
  lat!: number;

  // Aircraft Geo-Longitude for now;
  lng!: number;

  // Aircraft Registration Number
  reg_number!: string;

  // Aircraft horizontal speed (km) for now.
  speed!: number;

  // Current flight status - scheduled, en-route, landed.
  status!: string;

  // UNIX timestamp of last aircraft signal.
  updated!: number;

  // Aircraft vertical speed (km) for now.
  v_speed!: number;

  // Aircraft squawk signal code.
  squawk!: string;

  // Flight number only.
  flight_number!: string;

  // Flight ICAO code-number.
  flight_icao!: string;

	// Flight IATA code-number.
  flight_iata!: string;

  // Departure Airport ICAO code.
  dep_icao!: string;

  // Departure Airport IATA code.
  dep_iata!: string;

  // Arrival Airport ICAO code.
  arr_icao!: string;

  // Arrival Airport IATA code.
  arr_iata!: string;

  // Airline ICAO code.
  airline_icao!: string;

  // Airline IATA code.
  airline_iata!: string;
}

// hex
// reg_number	optional	Filtering by aircraft Registration number.
// airline_icao	optional	Filtering by Airline ICAO code.
// airline_iata	optional	Filtering by Airline IATA code.
// flag	optional	Filtering by Airline Country ISO 2 code from Countries DB.
// flight_icao	optional	Filtering by Flight ICAO code-number.
// flight_iata	optional	Filtering by Flight IATA code-number.
// flight_number	optional	Filtering by Flight number only.
// dep_icao	optional	Filtering by departure Airport ICAO code.
// dep_iata	optional	Filtering by departure Airport IATA code.
// arr_icao	optional	Filtering by arrival Airport ICAO code.
// arr_iata	optional	Filtering by arrival Airport IATA code.
// _fields	optional	Fields to return (comma separated, e.g.: hex,airline_iata,lat,lng)
// _view	optional	View format - object (default) or array (good for browser)
