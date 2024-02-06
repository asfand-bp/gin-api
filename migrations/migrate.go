package main

import (
	"app/db"
	"app/models"
)

func init() {
	db.StartDB() // Initialize the database connection
}

func main() {
	// Migrate the schema(s)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Grocery{})
}
