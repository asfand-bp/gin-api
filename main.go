package main

import (
	"app/db"
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) { // c.JSON serializes the map given in 2nd argument
	c.JSON(200, gin.H{ // H is shortcut for map[string]any
		"message": "Welcome to homepage",
	})
}

func init() {
	db.StartDB() // Initialize the database connection
}

func main() {
	r := gin.Default() // Gin engine with default stuff

	// API End points
	r.GET("/", index)
	// User endpoints
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.PostUser)
	r.GET("/users/:id", handlers.GetUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.Run() // listens and serves on 0.0.0.0:8080 (on lan and local)
}
