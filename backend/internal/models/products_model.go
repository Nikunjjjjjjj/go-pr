package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Url         string `json:"url" gorm:"default:https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTU073gNaeN3kAc7dWjS4ifOurPfBxK9X8rEwTTVIwp_WIq2WBKQHjot5r95ZL0GLlV_eZmAMXccU8Keg-8uOU9x2AcOlIMLAv-g8bHgH6z&s=10"`
}
