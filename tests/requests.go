package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// MakeGetRequest performs a GET request and returns the response recorder
func MakeGetRequest(t *testing.T, router *gin.Engine, url string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	return w
}

// MakePostRequest performs a POST request with the given data and returns the response recorder
func MakePostRequest(t *testing.T, router *gin.Engine, url string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal JSON data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "applicaton/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

// MakePutRequest performs a PUT request with the given data and returns the response recorder
func MakePutRequest(t *testing.T, router *gin.Engine, url string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal JSON data: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "applicaton/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

// MakeDeleteRequest performs a Delete request and returns the response recorder
func MakeDeleteRequest(t *testing.T, router *gin.Engine, url string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
