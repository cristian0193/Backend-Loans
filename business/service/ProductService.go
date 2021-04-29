package service

import (
	"Backend-Loans/domain/dto"
)

type ProductService interface {
	GetByIdMarketProduct(headers dto.Headers) (dto.Response, []dto.ProductDto)
	CreateProduct(productDto dto.ProductDto, headers dto.Headers) dto.Response
	UpdateProduct(productDto dto.ProductDto, headers dto.Headers) dto.Response
	DeleteProduct(idProduct string, headers dto.Headers) dto.Response
	GetByIdProduct(idProduct string, headers dto.Headers) (dto.Response, dto.ProductDto)
	GetByQueryParameters(queryParameters dto.QueryParameters, headers dto.Headers) (dto.Response, []dto.ProductDto)
}
