package service

import (
	"eCommerce/internal/dto"
	"eCommerce/internal/models"
	"eCommerce/internal/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

func (s ProductService) Create(newProd *dto.CreateProductDTO) error {
	product := models.Product{
		Name:        newProd.Name,
		Description: newProd.Description,
		Price:       newProd.Price,
		Url:         newProd.Url,
	}
	return s.Repo.Create(&product)
}

func (s ProductService) GetAll(limit, offset int) (*[]models.Product, int64, error) {
	//filteration and pagination to be added soon
	return s.Repo.GetAll(offset, limit)
}
