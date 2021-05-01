package service

import (
	"Backend-Loans/domain/component/errorException"
	"Backend-Loans/domain/dto"
	persistence "Backend-Loans/infrastructure/persistence/db"
	"Backend-Loans/infrastructure/repository"
	"Backend-Loans/utils"
	"log"
	"net/http"
)

type LoansServiceImpl struct {
	loansRepository repository.LoansRepository
	errorResponse   errorException.ErrorResponse
}

func InitLoansServiceImpl() *LoansServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &LoansServiceImpl{
		loansRepository: dbHelper.LoansRepository,
	}
}

func (a *LoansServiceImpl) CreateLoan(loansDto dto.LoansDto, headers dto.Headers) dto.Response {
	var responseDto = dto.Response{}

	err := a.loansRepository.Insert(loansDto, headers)

	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto
	}

	responseDto.Status = http.StatusCreated
	responseDto.Description = utils.StatusText(http.StatusCreated)
	responseDto.Message = utils.Lenguage(headers.Lenguage, "CREATED")
	return responseDto
}

/*  func (a *ProductServiceImpl) UpdateProduct(productDto dto.ProductDto, headers dto.Headers) dto.Response {
	var responseDto = dto.Response{}

	err := a.productRepository.Update(productDto, headers)

	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto
	}

	responseDto.Status = http.StatusCreated
	responseDto.Description = utils.StatusText(http.StatusCreated)
	responseDto.Message = utils.Lenguage(headers.Lenguage, "UPDATED")
	return responseDto
}

func (a *ProductServiceImpl) GetByIdMarketProduct(headers dto.Headers) (dto.Response, []dto.ProductDto) {
	var responseDto = dto.Response{}
	var listProductDto = make([]dto.ProductDto, 0)

	products, err := a.productRepository.GetByIdMarket(headers)
	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto, listProductDto
	}

	if len(products) == 0 || headers.IdMarket == "" {
		responseDto.Status = http.StatusNotFound
		responseDto.Description = utils.StatusText(http.StatusNotFound)
		responseDto.Message = utils.Lenguage(headers.Lenguage, "NOT_FOUND_PRODUCT")
		return responseDto, listProductDto
	}

	for _, product := range products {
		var productDto = dto.ProductDto{}
		model.Copy(&productDto, product)
		listProductDto = append(listProductDto, productDto)
	}

	responseDto.Status = http.StatusOK
	return responseDto, listProductDto
}

func (a *ProductServiceImpl) DeleteProduct(idProduct string, headers dto.Headers) dto.Response {
	var responseDto = dto.Response{}

	err := a.productRepository.Delete(idProduct, headers)

	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto
	}

	responseDto.Status = http.StatusOK
	responseDto.Description = utils.StatusText(http.StatusOK)
	responseDto.Message = utils.Lenguage(headers.Lenguage, "DELETE")
	return responseDto
}

func (a *ProductServiceImpl) GetByIdProduct(idProduct string, headers dto.Headers) (dto.Response, dto.ProductDto) {
	var responseDto = dto.Response{}
	var productDto = dto.ProductDto{}

	products, err := a.productRepository.GetByIdProduct(idProduct, headers)
	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto, productDto
	}

	if products.Id == 0 || headers.IdMarket == "" {
		responseDto.Status = http.StatusNotFound
		responseDto.Description = utils.StatusText(http.StatusNotFound)
		responseDto.Message = utils.Lenguage(headers.Lenguage, "NOT_FOUND_PRODUCT")
		return responseDto, productDto
	}

	model.Copy(&productDto, products)
	responseDto.Status = http.StatusOK
	return responseDto, productDto
}

func (a *ProductServiceImpl) GetByQueryParameters(queryParameters dto.QueryParameters, headers dto.Headers) (dto.Response, []dto.ProductDto) {
	var responseDto = dto.Response{}
	var listProductDto = make([]dto.ProductDto, 0)
	var result string = ""

	if queryParameters.IdCategory == "" {
		result = "NOT_FOUND_QUERY_IDCATEGORY"
	}

	if queryParameters.Price == "" {
		result = "NOT_FOUND_QUERY_PRICE"
	}

	if queryParameters.Name == "" {
		result = "NOT_FOUND_QUERY_NAME"
	}

	if result != "" {
		responseDto.Status = http.StatusNotFound
		responseDto.Description = utils.StatusText(http.StatusNotFound)
		responseDto.Message = utils.Lenguage(headers.Lenguage, result)
		return responseDto, listProductDto
	}

	products, err := a.productRepository.GetByQueryParameters(queryParameters, headers)
	if err != nil {
		responseDto.Status = http.StatusBadRequest
		responseDto.Description = utils.StatusText(http.StatusBadRequest)
		responseDto.Message = err.Error()
		return responseDto, listProductDto
	}

	if len(products) == 0 || headers.IdMarket == "" {
		responseDto.Status = http.StatusNotFound
		responseDto.Description = utils.StatusText(http.StatusNotFound)
		responseDto.Message = utils.Lenguage(headers.Lenguage, "NOT_FOUND_PRODUCT")
		return responseDto, listProductDto
	}

	for _, product := range products {
		var productDto = dto.ProductDto{}
		model.Copy(&productDto, product)
		listProductDto = append(listProductDto, productDto)
	}

	responseDto.Status = http.StatusOK
	return responseDto, listProductDto
}
*/
