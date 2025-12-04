package services

import (
	"UploadFileInSql/internal/models"
	"UploadFileInSql/internal/repo"
)

type DiscomService struct {
	Repo repo.DiscomRepository
}

func (s DiscomService) FindByPinCode(pincode string) (*models.OfficeData, error) {
	return s.Repo.FindByPinCode(pincode)
}
