package repository

import (
	"Backend-Loans/domain/dto"
)

type PaymentsRepository interface {
	Insert(paymentDto dto.PaymentDto) error
	UpdateBalance(id int32, balance float32) error
	FindLastBalance(idLoan int32) (dto.LastBalanceDto, error)
}
