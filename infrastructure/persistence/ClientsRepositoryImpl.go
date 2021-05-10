package persistence

import (
	"Backend-Loans/domain/entity"

	"github.com/jinzhu/gorm"
)

type ClientsRepositoryImpl struct {
	db *gorm.DB
}

func InitClientsRepositoryImpl(db *gorm.DB) *ClientsRepositoryImpl {
	return &ClientsRepositoryImpl{db}
}

func (repo *ClientsRepositoryImpl) FindAll() ([]entity.Clients, error) {

	var client = []entity.Clients{}

	err := repo.db.Find(&client).Error
	if err != nil {
		return client, err
	}
	return client, nil
}
