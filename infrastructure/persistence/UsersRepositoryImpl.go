package persistence

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"

	"github.com/jinzhu/gorm"
)

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func InitUsersRepositoryImpl(db *gorm.DB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{db}
}

func (repo *UsersRepositoryImpl) FindUserAndPassword(user dto.UsersDto) ([]entity.Users, error) {

	var users = []entity.Users{}

	err := repo.db.Where("username = ? and password = ?", user.UserName, user.Password).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
