package service

import (
	"Backend-Loans/domain/dto"
)

type LoansService interface {
	CreateLoan(loansDto dto.LoansDto, headers dto.Headers) dto.Response
	CreatePayment(paymentDto dto.PaymentDto, headers dto.Headers) dto.Response
	FindAllLoans(headers dto.Headers) ([]dto.ListLoansDto, dto.Response)
	FindByIdLoan(idLoan int32, headers dto.Headers) ([]dto.ListPaymentDto, dto.Response)
	FindInformacionByLoan(idLoan int32, headers dto.Headers) (dto.InformacionUserDto, dto.Response)
}
