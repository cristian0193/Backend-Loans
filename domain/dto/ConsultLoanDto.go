package dto

type ConsultLoanDto struct {
	Pages int            `json:"pages"`
	Loans []ListLoansDto `json:"loans"`
}
