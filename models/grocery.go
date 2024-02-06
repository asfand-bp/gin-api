package models

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
