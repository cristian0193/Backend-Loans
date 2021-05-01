package repository

import (
	"Backend-Loans/domain/dto"
)

type LoansRepository interface {
	Insert(loansDto dto.LoansDto, headers dto.Headers) error
}
