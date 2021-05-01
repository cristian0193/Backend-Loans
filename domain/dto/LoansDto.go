package dto

type LoansDto struct {
	IdentificationClient int32   `json:"identificationClient"`
	BorrowedValue        float32 `json:"borrowedValue"`
	InterestPercentage   float32 `json:"interestPercentage"`
}
