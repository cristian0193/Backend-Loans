package repository

import (
	"Backend-Loans/domain/entity"
)

type ClientsRepository interface {
	FindAll() ([]entity.Clients, error)
}
