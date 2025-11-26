package service

import (
	"eCommerce/internal/models"
	"eCommerce/internal/repository"
	"eCommerce/internal/utils"
)

type OrderService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
	orderRepo   repository.OrderRepository
	userRepo    repository.UserRepository
}

func (s OrderService) OrderCreate(userId int, address string) (*models.Order, error) {

	cart, _ := s.cartRepo.GetCartByUserID(userId)
	//extract cart.items
	total_price := 0
	//for _,val:=range cart.Items{}

	var order = models.Order{
		UserID:     uint(userId),
		TotalPrice: total_price,
		Address:    address,
	}
	s.orderRepo.CreateOrder(&order)
	// now orderidem
	// s.orderRepo.CreateOrderItem()
	for _, val := range cart.Items {
		product, _ := s.productRepo.FindByID(val.ProductID)

		total_price = total_price + product.Price*val.Quantity

		var orderItem = &models.OrderItem{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  val.Quantity,
		}
		s.orderRepo.CreateOrderItem(orderItem)
	}
	order.TotalPrice = total_price
	s.orderRepo.UpdateOrder(&order)
	//get by orderid
	//fmt.Println(order, 46)
	user, _ := s.userRepo.FindByID(userId)
	defer utils.OrderConfirmationMail(user.Email, &order)
	return s.orderRepo.FindOderbyId(int(order.ID))
	//delete element from cart(defer) and send email for confirmation

}

func (s *OrderService) GetAllOrder(userId int) (*[]models.Order, error) {
	return s.orderRepo.FindOrdersbyUserId(userId)
}
