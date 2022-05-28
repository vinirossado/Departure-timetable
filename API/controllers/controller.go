package controllers

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	mongoCollection *mongo.Collection
}

func NewBaseController(collection *mongo.Collection) *Controller {
	return &Controller{
		mongoCollection: collection,
	}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
