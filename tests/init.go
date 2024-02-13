package tests

import (
	"app/db"
	"app/models"

	"github.com/gin-gonic/gin"
)

func startdb() {
	db.StartDB() // Initialize the database connection
	// Migrate the schema(s)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Grocery{})
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func init() {
	startdb()
}
