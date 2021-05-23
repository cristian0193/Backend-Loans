package service

import (
	"Backend-Loans/domain/dto"
)

type ClientService interface {
	FindAll() ([]dto.ClientsDto, dto.Response)
	FindById(id int32) (dto.ClientsDto, dto.Response)
	Create(create dto.ClientsDto, headers dto.Headers) dto.Response
	Update(clientDto dto.ClientsDto, headers dto.Headers) dto.Response
}
