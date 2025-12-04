package repo

import (
	"UploadFileInSql/internal/database"
	"UploadFileInSql/internal/models"
)

type DiscomRepository struct {
}

func (r DiscomRepository) FindByPinCode(pincode string) (*models.OfficeData, error) {
	var data models.OfficeData
	err := database.DB.Where("pincode=?", pincode).First(&data).Error
	return &data, err
}
