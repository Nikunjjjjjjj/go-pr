package repository

import (
	"eCommerce/internal/database"
	"eCommerce/internal/models"
)

type UserRepository struct {
}

func (r UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r UserRepository) FindByID(userId int) (*models.User, error) {
	var user models.User
	err := database.DB.Where("ID=?", userId).First(&user).Error
	return &user, err
}

//	func (r UserRepository) Create(name, email, hashedPassword, role string) (string, error) {
//		var user = models.User{
//			Name:     name,
//			Email:    email,
//			Password: hashedPassword,
//			Role:     role,
//		}
//		result := database.DB.Create(&user)
//		if result.Error != nil {
//			// if pgErr,ok:=result.Error.(*pgconn.PgError);ok{
//			// 	//"error":"email already exists ",
//			// 	c.J
//			// }
//			return "failed", result.Error
//		}
//		return "success", result.Error
//	}
func (r UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
