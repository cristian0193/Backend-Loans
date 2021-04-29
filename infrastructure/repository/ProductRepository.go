package repository

import (
	"Backend-Loans/domain/dto"
	"Backend-Loans/domain/entity"
)

type ProductRepository interface {
	GetByIdMarket(headers dto.Headers) ([]entity.Product, error)
	Insert(productDto dto.ProductDto, headers dto.Headers) error
	Update(productDto dto.ProductDto, headers dto.Headers) error
	Delete(idProduct string, headers dto.Headers) error
	GetByIdProduct(idProduct string, headers dto.Headers) (entity.Product, error)
	GetByQueryParameters(queryParameters dto.QueryParameters, headers dto.Headers) ([]entity.Product, error)
}
