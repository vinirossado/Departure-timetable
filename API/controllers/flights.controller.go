package controllers

import (
	"departure_time_table/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetAllFlights(ctx *gin.Context) {
	var flightResponse models.FlightResponse

	response, err := http.Get(os.Getenv("API_FLIGHT_URL"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &flightResponse); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": flightResponse.Response,
	})
}

func (c *Controller) GetFlight(ctx *gin.Context) {
	var flight models.Flight

	response, err := http.Get(os.Getenv("API_FLIGHT_URL") + "/flight/" + ctx.Params.ByName("id"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &flight); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": flight,
	})
}
