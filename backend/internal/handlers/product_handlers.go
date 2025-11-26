package handlers

import (
	"eCommerce/internal/dto"
	"eCommerce/internal/service"
	"eCommerce/internal/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var productService = service.ProductService{}

func ADDNew(c *gin.Context) {
	//fmt.Println("insssjij")
	userID, _ := c.Get("user_id")
	fmt.Println(userID)
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "plls enter proper format",
		})
	}

	timestamp := time.Now().Unix() // e.g. 1732558800

	newFileName := fmt.Sprintf("%d_%s", timestamp, file.Filename)
	c.SaveUploadedFile(file, "./files/"+newFileName)
	destination := "./files/" + newFileName
	defer func() {
		if err := os.Remove(destination); err != nil {
			fmt.Println("error deleting file", err)
		}
	}()
	//added by userID
	var createProd = dto.CreateProductDTO{}
	if err := c.ShouldBind(&createProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      "invalid body",
			"real Error": err.Error(),
		})
		return
	}
	//fmt.Println(createProd, 43)

	awsUrl, err := utils.AwsUpload(destination, newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in uploading on aws",
		})
	}
	createProd.Url = awsUrl
	//fmt.Println(createProd, 53)
	if err := productService.Create(&createProd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal erroe",
		})
		return
	}
	//defer delete filepath
	c.JSON(http.StatusAccepted, gin.H{
		"message": "accepted",
	})
}

func GetAll(c *gin.Context) {
	// offset,limit,searching,totalitem,
	// limit, _ := strconv.Atoi(c.Query("limit"))
	// Offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, total, err := productService.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	//fmt.Println(products)

	c.JSON(http.StatusOK, gin.H{
		"totalItems": total, //change karna hai
		"products":   *products,
	})
}
