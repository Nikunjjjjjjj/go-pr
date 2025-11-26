package database

import (
	"eCommerce/internal/config"
	"eCommerce/internal/models"
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	conf := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DBHOST, conf.DBUSER, conf.DBPassword, conf.DBName, conf.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect db:", err)
	} else {
		//db.AutoMigrate(&models.User{}, &models.Products{}, &models.Cart{}, &models.CartItem{}, &models.Orders{})

		db.AutoMigrate(
			&models.User{},
			&models.Product{},
			&models.Cart{},
			&models.CartItem{},
			&models.Order{},
			&models.OrderItem{},
		)

		DB = db
		fmt.Println("connected to db")
	}

}
