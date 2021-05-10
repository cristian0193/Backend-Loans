package persistence

import (
	"Backend-Loans/domain/entity"

	"github.com/jinzhu/gorm"
)

type TypesRepositoryImpl struct {
	db *gorm.DB
}

func InitTypesRepositoryImpl(db *gorm.DB) *TypesRepositoryImpl {
	return &TypesRepositoryImpl{db}
}

func (repo *TypesRepositoryImpl) FindAll() ([]entity.Types, error) {

	var types = []entity.Types{}

	err := repo.db.Find(&types).Error
	if err != nil {
		return types, err
	}
	return types, nil
}
