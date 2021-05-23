package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"

	"github.com/jinzhu/gorm"
	"gopkg.in/jeevatkm/go-model.v1"
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

func (repo *ClientsRepositoryImpl) FindById(id int32) (entity.Clients, error) {

	var client = entity.Clients{}

	err := repo.db.Where("identification = ?", id).Find(&client).Error
	if err != nil {
		return client, err
	}
	return client, nil
}

func (repo *ClientsRepositoryImpl) Create(clienDto dto.ClientsDto) error {

	var client = entity.Clients{}
	model.Copy(&client, clienDto)

	err := repo.db.Create(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ClientsRepositoryImpl) Update(clienDto dto.ClientsDto) error {

	var client = entity.Clients{}
	model.Copy(&client, clienDto)

	err := repo.db.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}
