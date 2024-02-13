package tests

import (
	"app/handlers"
	"app/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup the router and routes
	r := SetUpRouter()
	r.POST("/users", handlers.PostUser)

	// User data jsonified
	userData, _ := json.Marshal(
		models.User{
			Username:  "test_user",
			FirstName: "Test",
			LastName:  "User",
		},
	)

	// Make the post request with user data
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer((userData)))

	// Mock the api call
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users", handlers.GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}
