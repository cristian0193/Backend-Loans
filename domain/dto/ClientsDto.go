package dto

type ClientsDto struct {
	Identification int32  `json:"identification"`
	FullName       string `json:"fullname"`
	Address        string `json:"address"`
	Mobile         string `json:"mobile"`
}
