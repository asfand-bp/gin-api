package tests

import (
	"app/db"
	"app/handlers"
	"app/models"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func StartDBAndMigrate() {
	db.StartDB() // Initialize the database connection
	// Migrate the schema(s)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Grocery{})
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	// Routes
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.PostUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	return router
}

func DestroyDB() {
	databaseFilePath := "./app.db"
	err := os.Remove(databaseFilePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}
}

func CreateUser(t *testing.T, R *gin.Engine) models.User {
	// Define the Attendance data
	userData := models.User{
		Username:  "new_user",
		FirstName: "New",
		LastName:  "User",
		Email:     "newuser@gmail.com",
	}

	// Make POST api request
	w := MakePostRequest(t, R, "/users", userData)

	// Parse the response body
	var user models.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	return user
}

func DeleteUser(t *testing.T, R *gin.Engine, id string) {
	// Make Delete api request
	MakeDeleteRequest(t, R, "/users/"+id)
}
