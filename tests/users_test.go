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
	DeleteUser(t, R, fmt.Sprintf("%v", user.ID))
}

func TestGetUsers(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make GET api request
	w := MakeGetRequest(t, R, "/users")

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestGetUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make GET api request
	w := MakeGetRequest(t, R, fmt.Sprintf("/users/%v", user_id))

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestUpdateUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	NEW_FIRSTNAME := "test_user_new"
	body := models.UserUpdate{
		FirstName: NEW_FIRSTNAME,
	}

	// Make PUT api request
	w := MakePutRequest(t, R, fmt.Sprintf("/users/%v", user_id), body)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.FirstName, NEW_FIRSTNAME)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestDeleteUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make PUT api request
	w := MakeDeleteRequest(t, R, fmt.Sprintf("/users/%v", user_id))

	assert.Equal(t, http.StatusOK, w.Code)
}
