package repository

import (
	"eCommerce/internal/database"
	"eCommerce/internal/models"
)

type CartRepository struct{}

func (r CartRepository) GetCartByUserID(userId int) (*models.Cart, error) {
	var cart models.Cart
	err := database.DB.Preload("Items").Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		return nil, err
	}
	//err:= database.DB.Where("user_id=?",userId).First(&cart).Error
	return &cart, err
}

func (r CartRepository) CreateCart(cart *models.Cart) error {
	err := database.DB.Create(cart)
	return err.Error
}

func (r CartRepository) AddItemToCartItem(item *models.CartItem) error {
	err := database.DB.Create(item).Error
	return err
}

func (r CartRepository) UpdateItemToCartItem(item *models.CartItem) error {
	err := database.DB.Save(item).Error
	return err
}

func (r CartRepository) FindItemInCartItem(cartId, productId int) (*models.CartItem, error) {
	var item models.CartItem
	err := database.DB.Where("cart_id=? AND product_id=?", cartId, productId).First(&item).Error
	return &item, err
}

//empty cart
// 🚀 DELETE ENTIRE CART FOR USER (future use)
//////////////////////////////////////////////////////

func (r CartRepository) DeleteCart(userId int) error {
	// Step 1 — find cart
	var cart models.Cart
	err := database.DB.Where("user_id = ?", userId).First(&cart).Error
	if err != nil {
		return err
	}

	// Step 2 — delete all items in that cart
	err = database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
	if err != nil {
		return err
	}

	// Step 3 — delete cart itself
	return database.DB.Delete(&cart).Error
}

// delete item from cart
func (r CartRepository) DeleteCartItems(cartId, productId int) error {
	return database.DB.Where("cart_id=? AND product_id=?", cartId, productId).
		Delete(&models.CartItem{}).Error
}
