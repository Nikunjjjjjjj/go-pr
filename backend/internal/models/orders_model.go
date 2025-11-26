// package models

// import "gorm.io/gorm"

// type Orders struct {
// 	gorm.Model
// 	UserId          int
// 	ProductsOrdered []Cart `json:"productsordered" gorm:"foreignKey:UserId"`
// 	TotalPrice      int
// 	Address         string
// 	IsDelivered     bool `json:"isdelivered" gorm:"default:false"`
// }

package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint        `json:"user_id"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	TotalPrice  int         `json:"total_price" gorm:"default:0"`
	Address     string      `json:"address"`
	IsDelivered bool        `json:"is_delivered" gorm:"default:false"`
}
