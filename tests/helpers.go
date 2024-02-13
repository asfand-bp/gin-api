package tests

import (
	"app/db"
	"app/handlers"
	"app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// MakeGetRequest performs a GET request and returns the response recorder
func MakeGetRequest(t *testing.T, router *gin.Engine, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	return w
}

// MakePostRequest performs a POST request with the given data and returns the response recorder
func MakePostRequest(t *testing.T, router *gin.Engine, path string, data interface{}) *httptest.ResponseRecorder {
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal JSON data: %v", err)
	}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
