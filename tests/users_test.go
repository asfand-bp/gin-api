package tests

import (
	"app/models"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// User data jsonified
	userData := models.User{
		Username:  "test_user",
		FirstName: "Test",
		LastName:  "User",
	}

	// Make the api request
	w := MakePostRequest(t, R, "/users", userData)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	// Make the api request
	w := MakeGetRequest(t, R, "/users")

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}

func TestGetUser(t *testing.T) {
	// Make the GET api request
	w := MakeGetRequest(t, R, "/users/1")

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
}
