package dto

type InterestsDto struct {
	Id     int32   `json:"id"`
	IdLoan int32   `json:"id_loan"`
	Share  float32 `json:"share"`
	Status string  `json:"status"`
}
