package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//handler for GET request (GET, POST, PATCH, DELETE)
	// if a "/events" is requested then the function (can be anonymous or named) will be executed
	server.GET("/events", getEvents)

	server.GET("/events/:id", getEvent)    // /events/5, /events/1
	server.POST("/events", createEvent)    // in order to create a new event and should have some data for us
	server.PUT("/events/:id", updateEvent) // to update an event
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup) // signup to create new users
}
