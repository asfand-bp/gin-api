package tests

import (
	"app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// User data
	userData := models.User{
		Username:  "test_user",
		FirstName: "Test",
		LastName:  "User",
	}

	// Make POST api request
	w := MakePostRequest(t, R, "/users", userData)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, user)

	// Assigning the user id for use in later tests
	USER_ID = user.ID
}

func TestGetUsers(t *testing.T) {
	// Make GET api request
	w := MakeGetRequest(t, R, "/users")

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)
}

func TestGetUser(t *testing.T) {
	// Make GET api request
	w := MakeGetRequest(t, R, fmt.Sprintf("/users/%v", USER_ID))

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
}

func TestUpdateUser(t *testing.T) {
	NEW_FIRSTNAME := "test_user_new"

	body := models.UserUpdate{
		FirstName: NEW_FIRSTNAME,
	}

	// Make PUT api request
	w := MakePutRequest(t, R, fmt.Sprintf("/users/%v", USER_ID), body)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.FirstName, NEW_FIRSTNAME)
}

func TestDeleteUser(t *testing.T) {
	// Make PUT api request
	w := MakeDeleteRequest(t, R, fmt.Sprintf("/users/%v", USER_ID))

	assert.Equal(t, http.StatusOK, w.Code)
}
