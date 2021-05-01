package dto

type PaymentDto struct {
	IdLoan   int32   `json:"idLoan"`
	Capital  float32 `json:"capital"`
	Interest float32 `json:"interest"`
	Balance  float32 `json:"balance"`
	IdType   int     `json:"idType"`
}
