package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ssakyp/rest-api/db"
	"github.com/ssakyp/rest-api/routes"
)

func main() {
	db.InitDB()

	//http service that already has some basic functionality
	//returns a handle pointer
	server := gin.Default()

	routes.RegisterRoutes(server)
	// listens on some incoming requests with domain locally 8080
	server.Run(":8080") //localhost:8080
}
