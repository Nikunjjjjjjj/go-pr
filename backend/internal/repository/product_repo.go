package repository

import (
	"eCommerce/internal/database"
	"eCommerce/internal/models"
)

type ProductRepository struct{}

func (r ProductRepository) Create(product *models.Product) error {

	return database.DB.Create(product).Error
}

func (r ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := database.DB.Where("ID=?", id).First(&product).Error
	return &product, err
}

func (r ProductRepository) GetAll(offset, limit int) (*[]models.Product, int64, error) {
	var products []models.Product
	var total int64
	err := database.DB.Offset(offset).Limit(limit).Find(&products).Error
	//i want to return total number of products
	if err := database.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return &products, total, err
}

// // Optional search
// if search != "" {
// 	query = query.Where("name ILIKE ?", "%"+search+"%")
// }

// // Optional filter example (customize as needed)
// if filter != "" {
// 	query = query.Where("category = ?", filter)
// }

//upgrade,delete
