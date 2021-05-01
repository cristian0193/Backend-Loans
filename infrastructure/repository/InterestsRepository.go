package repository

import (
	"Backend-Loans/domain/dto"
)

type InterestsRepository interface {
	Insert(interestsDto dto.InterestsDto) error
	UpdateStateAllInterestByIdLoan(idLoan int32) error
}
