package service

import (
	"Backend-Loans/domain/dto"
)

type ClientService interface {
	FindAll() ([]dto.ClientsDto, dto.Response)
}
