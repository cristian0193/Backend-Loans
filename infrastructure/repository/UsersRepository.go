package repository

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
)

type UsersRepository interface {
	FindUserAndPassword(user dto.UsersDto) ([]entity.Users, error)
}
