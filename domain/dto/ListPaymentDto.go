package dto

import "time"

type ListPaymentDto struct {
	Id          int32     `json:"id"`
	Capital     float32   `json:"capital"`
	Interest    float32   `json:"interest"`
	Balance     float32   `json:"balance"`
	PaymentDate time.Time `json:"paymentDate"`
	IdType      int32     `json:"idType"`
	Type        string    `json:"type"`
}
