package handlers

import (
	"eCommerce/internal/dto"
	"eCommerce/internal/service"
	"eCommerce/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = service.UserService{}

func SignUP(c *gin.Context) {
	// var signUpBody struct{
	// 	Name string `json:"name" binding:"required"`
	// 	Email    string `json:"email" binding:"required"`
	// 	Password string `json:"password" binding:"required"`
	// 	Role string `json:"role" default:"user"`
	// }
	fmt.Println("inside signUp")
	var signUpBody = dto.SignInDto{}
	if err := c.BindJSON(&signUpBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "Invalid input",
			"real error": err.Error(),
		})
		return
	}
	fmt.Println("reached 30")
	if err := userService.CreateUser(&signUpBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(200, "user created")
}

func Login(c *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Output",
		})
		return
	}
	// token, err := utils.GenerateJWT(body.Email)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
	// 	return
	// }

	user, err := userService.FindByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":     "no user found",
			"realError": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":     "Token error",
			"realError": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token":   token,
		"message": "logiedIn success",
	})
}
