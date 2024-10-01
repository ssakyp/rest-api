package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssakyp/rest-api/models"
)

func main() {
	//http service that already has some basic functionality
	//returns a handle pointer
	server := gin.Default()

	//handler for GET request (GET, POST, PATCH, DELETE)
	// if a "/events" is requested then the function (can be anonymous or named) will be executed
	server.GET("/events", getEvents)

	server.POST("/events", createEvent)	// in order to create a new event and should have some data for us 

	// listens on some incoming requests with domain locally 8080
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	//the function sends back response

	events := models.GetAllEvents()

	// arguments for below function are status code and data (structs, maps and slices (gin.H{"message": "Hello!"}) to be sent
	context.JSON(http.StatusOK, events) // HTML function can also be used but it is uncommon in REST API
}


func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)		//works like a fmt.Scan => stores the data into event => make sure that JSON corresponds our Event struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the data!"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message":"Event created!", "event": event})
}