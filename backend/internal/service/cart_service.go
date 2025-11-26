package service

import (
	"eCommerce/internal/models"
	"eCommerce/internal/repository"
)

type CartService struct {
	cartRepo repository.CartRepository
	//productRepo repository.ProductRepository
}

func (s *CartService) AddToCart(userId, productId, quantity int) (*models.Cart, error) {
	//var cart models.Cart
	//check cart is there? create one
	cart, err := s.cartRepo.GetCartByUserID(userId)
	if err != nil {
		cart = &models.Cart{UserId: userId}
		s.cartRepo.CreateCart(cart)
	}
	//check pid is valid
	// product,err:=s.productRepo.FindByID(uint(productId))
	// if err!=nil{
	// 	errors.New("invalid product Id")
	// 	return
	// }
	//check item -add update
	item, err := s.cartRepo.FindItemInCartItem(int(cart.ID), productId)
	if err == nil {
		//update
		item.Quantity += quantity
		s.cartRepo.UpdateItemToCartItem(item)
	} else {
		newitem := models.CartItem{
			CartID:    cart.ID,
			Quantity:  quantity,
			ProductID: uint(productId),
		}
		s.cartRepo.AddItemToCartItem(&newitem)
	}
	//get back updated cart
	return s.cartRepo.GetCartByUserID(userId)
}
