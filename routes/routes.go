package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ssakyp/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//handler for GET request (GET, POST, PATCH, DELETE)
	// if a "/events" is requested then the function (can be anonymous or named) will be executed
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // /events/5, /events/1

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)    // in order to create a new event and should have some data for us
	authenticated.PUT("/events/:id", updateEvent) // to update an event
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup) // signup to create new users
	server.POST("/login", login)   //handle user login
}
