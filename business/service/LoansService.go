package service

import (
	"Backend-Loans/domain/dto"
)

type LoansService interface {
	CreateLoan(loansDto dto.LoansDto, headers dto.Headers) dto.Response
	CreatePayment(paymentDto dto.PaymentDto, headers dto.Headers) dto.Response
}
