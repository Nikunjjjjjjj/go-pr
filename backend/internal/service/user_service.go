package service

import (
	"eCommerce/internal/dto"
	"eCommerce/internal/models"
	"eCommerce/internal/repository"
	"eCommerce/internal/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

// func (s UserService) CreateUser(name, email, role, password string) error {
func (s UserService) CreateUser(signUpBody *dto.SignInDto) error {
	hash, _ := utils.HashPassword(signUpBody.Password)
	user := models.User{Name: signUpBody.Name, Email: signUpBody.Email, Password: string(hash), Role: signUpBody.Role}
	return s.Repo.Create(&user)
}

func (s UserService) FindByEmail(email string) (*models.User, error) {
	return s.Repo.FindByEmail(email)
}

// func (s UserService)
