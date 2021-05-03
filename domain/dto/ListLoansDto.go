package dto

type ListLoansDto struct {
	Id            int32   `json:"id"`
	Client        int32   `json:"client"`
	Name          string  `json:"name"`
	BorrowedValue float32 `json:"borrowedValue"`
	PaidValue     float32 `json:"paidValue"`
	PendingValue  float32 `json:"pendingValue"`
	InterestPaid  float32 `json:"interestPaid"`
	IdStatus      int     `json:"idStatus"`
}
