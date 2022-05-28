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

func (c *Controller) GetAirportByIATACode(ctx *gin.Context) {
	var airport models.Airport

	response, err := http.Get(os.Getenv("API_AIRPORT_URL") + "/airport/" + ctx.Params.ByName("id"))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(responseData, &airport); err != nil {
		panic(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": airport,
	})
}
