package models

import (
	"app/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" binding:"required,max=35"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// NOT USING THESE FUNCTIONS BELOW FOR NOW

func GetUsers() []User {
	var users []User
	db.DB.Find(&users)
	return users
}

func GetUser(id int) User {
	var user User
	db.DB.First(&user, id)
	return user
}

func PostUser(user *User) {
	db.DB.Create(user)
}
