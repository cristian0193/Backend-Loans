package service

import (
	"Backend-Loans/domain/dto"
)

type UserService interface {
	FindByUserAndPassword(user dto.UsersDto, headers dto.Headers) dto.Response
}
