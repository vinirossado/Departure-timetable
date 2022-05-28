package main

import (
	"departure_time_table/controllers"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	r := gin.New()
	r.Use(cors.Default())

	gin.SetMode(gin.ReleaseMode)

	c := controllers.NewController()

	v1 := r.Group("/api/v1")
	{
		flights := v1.Group("/flights")
		{
			flights.GET(":id", c.GetFlight)
			flights.GET("", c.GetAllFlights)
		}

		airports := v1.Group("/airports")
		{
			airports.GET(":id", c.GetAirportByIATACode)
		}

	}

	http.ListenAndServe(":8080", r)
}
