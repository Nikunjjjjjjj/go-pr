package handlers

import (
	"eCommerce/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var cartService = service.CartService{}
// var productService=service.ProductService{}
type CartHandler struct {
	cartService    service.CartService
	productService service.ProductService
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	//prodid,quantity,userid
	var userId, _ = c.Get("user_id")
	// var productId, _ = strconv.Atoi(c.Query("pid"))
	// var quantity, _ = strconv.Atoi(c.Query("q"))
	productIdStr := c.Query("pid")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or missing productId",
		})
		return
	}

	quantityStr := c.Query("q")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid or missing quantity",
		})
		return
	}
	product, err := h.productService.Repo.FindByID(uint(productId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "worong pid",
		})
		fmt.Println(product)
		return
	}
	// fmt.Println(userId, productId, quantity)
	// fmt.Printf("type1:%T and type2:%T", userId, quantity)

	cart, err := h.cartService.AddToCart(int(userId.(float64)), productId, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "under development",
		"cart":    cart,
	})

}
