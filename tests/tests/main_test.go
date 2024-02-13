package tests

import (
	"app/db"
	"app/models"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var BASE_URL string
var API_URL string

func startdb() {
	db.StartDB() // Initialize the database connection
	// Migrate the schema(s)
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Grocery{})
}

func setup_router() {
	router := gin.Default()

	go func() {
		router.Run(":8080")
	}()

	time.Sleep(10 * time.Second)
}

func setup() {
	// Add your setup code here.
	// This code will run before the tests.
	// You can initialize resources, set up configurations, etc.
	BASE_URL = "http://localhost:8080/"
	API_URL = fmt.Sprintf("%s/api", BASE_URL)

	startdb()
	setup_router()
}

func teardown() {
	// Add your teardown code here.
	// This code will run after the tests.
	// You can clean up resources, close connections, etc.
}

func TestMain(m *testing.M) {
	// Run setup code before running the tests
	setup()

	// Run the tests
	exitCode := m.Run()

	// Run teardown code after the tests
	teardown()

	// Exit with the exit code from the tests
	os.Exit(exitCode)
}
