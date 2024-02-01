package main

import "github.com/gin-gonic/gin"

func index(c *gin.Context) { // c.JSON serializes the map given in 2nd argument
	c.JSON(200, gin.H{ // H is shortcut for map[string]any
		"message": "Welcome to homepage",
	})
}

func main() {
	r := gin.Default() // Gin engine with default stuff

	// API End points
	r.GET("/", index)

	r.Run() // listens and serves on 0.0.0.0:8080 (on lan and local)
}
