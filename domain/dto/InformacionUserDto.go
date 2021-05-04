package dto

import "time"

type InformacionUserDto struct {
	Id             int32     `json:"id"`
	Identification int32     `json:"identification"`
	FullName       string    `json:"fullname"`
	Mobile         string    `json:"mobile"`
	Address        string    `json:"address"`
	BorrowedValue  float32   `json:"borrowedValue"`
	Interest       float32   `json:"interest"`
	MonthlyFee     float32   `json:"monthlyFee"`
	LoanDate       time.Time `json:"loanDate"`
	MonthsArrears  int       `json:"monthsArrears"`
	State          string    `json:"state"`
}
