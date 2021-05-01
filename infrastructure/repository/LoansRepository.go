package repository

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
)

type LoansRepository interface {
	Insert(loansDto dto.LoansDto) (int32, error)
	FindUserByUser(identification int32) (bool, error)
	FindUserById(id int32) (entity.Loans, error)
}
