package repository

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
)

type LoansRepository interface {
	Insert(loansDto dto.LoansDto) (int32, error)
	FindUserByUser(identification int32) (bool, error)
	FindUserById(id int32) (entity.Loans, error)
	UpdateCalculateById(id int32) error
	FindAllLoans(query dto.QueryParameters) ([]entity.Loans, error)
	FindInformacionByIdLoan(idLoan int32) (entity.Loans, error)
	CountAllLoans() (int, error)
}
