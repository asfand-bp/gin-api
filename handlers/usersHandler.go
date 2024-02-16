package handlers

import (
	"app/db"
	"app/models"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func GetUsers(c *gin.Context) {
	// Query all users
	var users []models.User
	db.DB.Find(&users)

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	// Extract id from url and convert it to int
	id, _ := strconv.Atoi(c.Param("id"))

	// Query user with provided id
	var user models.User
	res := db.DB.First(&user, id)

	// If any error
	if res.Error != nil {
		// If record is not found with provided id
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		} else { // Else any other error
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}
	}
	// Everything's good and return the user object
	c.JSON(200, user)
}

func PostUser(c *gin.Context) {
	var u models.User

	// Bind JSON from the request body
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&u) // Create new user (insert new row)

	c.JSON(201, u)
}

func UpdateUser(c *gin.Context) {
	// Extract id from url and convert it to int
	id, _ := strconv.Atoi(c.Param("id")) // NO NEED TO TYPE CAST to INT

	var user models.User
	res := db.DB.First(&user, id) // id can be either in string format or int doesnt matter

	// If any error
	if res.Error != nil {
		// If record is not found with provided id
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		} else { // Else any other error
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}
	}

	// Bind JSON from the request body
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate data
	err := validate.Struct(user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Validation error"})
		return
	}

	// Update the user in the database
	if err := db.DB.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	// Extract id from url and convert it to int
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.User
	res := db.DB.First(&user, id)

	// If any error
	if res.Error != nil {
		// If record is not found with provided id
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		} else { // Else any other error
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}
	}
	// All good, delete the user
	res = db.DB.Delete(&user, id)
	if res.Error != nil {
		c.JSON(500, "Failed to delete user")
	}
	c.JSON(200, gin.H{
		"detail": "Sucessfully deleted the user",
	})
}
