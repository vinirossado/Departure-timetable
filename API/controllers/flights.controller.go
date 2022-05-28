package controllers

import (
	"context"
	"departure_time_table/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) GetAllFlights(ctx *gin.Context) {

	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// retrieve all the documents that match the filter
	cursor, err := c.mongoCollection.Find(context.TODO(), filter)
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}

	// retrieving the first document that match the filter
	var result bson.M
	// check for errors in the finding
	if err = c.mongoCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		panic(err)
	}

	// display the document retrieved
	fmt.Println("displaying the first result from the search filter")
	fmt.Println(result)

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
