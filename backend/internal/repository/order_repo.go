package repository

import (
	"eCommerce/internal/database"
	"eCommerce/internal/models"
)

type OrderRepository struct {
}

//func checkOrderby userid
//func create order
//func add order_item

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}

func (r *OrderRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	return database.DB.Create(orderItem).Error
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	return database.DB.Save(order).Error
}

func (r *OrderRepository) FindOrdersbyUserId(userId int) (*[]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("Items").Where("user_id=?", userId).Find(&orders).Error
	return &orders, err
} //mistake

func (r *OrderRepository) FindOderbyId(orderId int) (*models.Order, error) {
	var order models.Order
	err := database.DB.Preload("Items").Where("id=?", orderId).First(&order).Error
	return &order, err
}
