package handlers

import (
	"eCommerce/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var orderService = service.OrderService{}

type AddressS struct {
	Address string `json:"address" `
}

func CreateOrder(c *gin.Context) {
	userId, _ := c.Get("user_id")
	//check cart exists
	//order
	var addBody AddressS
	if err := c.ShouldBindJSON(&addBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "pls enter your address",
		})
		return
	}
	res, err := orderService.OrderCreate(int(userId.(float64)), addBody.Address)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"order":   res,
	})
}

func GetAllOrder(c *gin.Context) {
	userId, _ := c.Get("user_id")
	res, err := orderService.GetAllOrder(int(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": res,
	})
}
