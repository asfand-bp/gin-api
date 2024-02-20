package tests

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine // Variale to hold Reference to Router Engine

type BadBodyType struct {
	BadField string
}

func init() {
	StartDBAndMigrate()
	R = SetUpRouter()
}

func setup() {
	// Add your setup code here.
	// This code will run before the tests.
	// You can initialize resources, set up configurations, etc.
}

func teardown() {
	// Add your teardown code here.
	// This code will run after the tests.
	// You can clean up resources, close connections, etc.

	// Destroy or delete database
	DestroyDB()
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
