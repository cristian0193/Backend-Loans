package repository

import (
	"Backend-Loans/domain/entity"
)

type TypesRepository interface {
	FindAll() ([]entity.Types, error)
}
