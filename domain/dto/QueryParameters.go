package dto

type QueryParameters struct {
	Pages    uint   `json:"pages"`
	Fullname string `json:"fullname"`
	Status   uint   `json:"status"`
}
