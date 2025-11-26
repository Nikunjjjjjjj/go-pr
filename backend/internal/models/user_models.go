// package models

// import "gorm.io/gorm"

// type User struct {
// 	gorm.Model
// 	Name     string `json:"name"`
// 	Email    string `json:"email" gorm:"unique"`
// 	Password string `json:"password"`
// 	Role     string `json:"role" gorm:"default:user"`
// }

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`

	// Relations
	Orders []Order `gorm:"foreignKey:UserID"`
	Cart   Cart    `gorm:"foreignKey:UserID"`
}
