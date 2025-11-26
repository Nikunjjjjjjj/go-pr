package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId int        `json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}

// package models

// import "gorm.io/gorm"

// type Cart struct {
// 	gorm.Model
// 	UserID uint       `json:"user_id"`
// 	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
// }
