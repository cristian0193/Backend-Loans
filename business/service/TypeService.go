package service

import (
	"Backend-Loans/domain/dto"
)

type TypeService interface {
	FindAll() ([]dto.TypesDto, dto.Response)
}
