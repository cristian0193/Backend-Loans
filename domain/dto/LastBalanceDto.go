package dto

type LastBalanceDto struct {
	Id                 int32   `json:"id"`
	Balance            float32 `json:"balance"`
	InterestPercentage float32 `json:"interest_percentage"`
}
