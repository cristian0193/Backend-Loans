package repository

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
)

type ClientsRepository interface {
	FindAll() ([]entity.Clients, error)
	FindById(id int32) (entity.Clients, error)
	Create(clienDto dto.ClientsDto) error
	Update(clienDto dto.ClientsDto) error
}
