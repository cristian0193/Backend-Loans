package dto

type ProductDto struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	IdCategory int     `json:"id_category"`
}
