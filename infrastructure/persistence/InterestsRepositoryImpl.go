package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"

	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
)

type InterestsRepositoryImpl struct {
	db *gorm.DB
}

func InitInterestsRepositoryImpl(db *gorm.DB) *InterestsRepositoryImpl {
	return &InterestsRepositoryImpl{db}
}

func (repo *InterestsRepositoryImpl) Insert(interestsDto dto.InterestsDto) error {

	var interest = entity.Interests{}
	model.Copy(&interest, interestsDto)

	err := repo.db.Create(&interest).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *InterestsRepositoryImpl) UpdateStateAllInterestByIdLoan(idLoan int32) error {

	err := repo.db.Table("Interests").Where("id_loan = ?", idLoan).
		Updates(map[string]interface{}{"status": "INA"}).Error

	if err != nil {
		return err
	}
	return nil
}
