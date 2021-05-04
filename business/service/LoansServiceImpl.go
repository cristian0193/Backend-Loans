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
	loansRepository     repository.LoansRepository
	paymentsRepository  repository.PaymentsRepository
	interestsRepository repository.InterestsRepository
	errorResponse       errorException.ErrorResponse
}

func InitLoansServiceImpl() *LoansServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &LoansServiceImpl{
		loansRepository:     dbHelper.LoansRepository,
		paymentsRepository:  dbHelper.PaymentsRepository,
		interestsRepository: dbHelper.InterestsRepository,
	}
}

func (a *LoansServiceImpl) CreateLoan(loansDto dto.LoansDto, headers dto.Headers) dto.Response {

	result, err := a.loansRepository.FindUserByUser(loansDto.IdentificationClient)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	if result != false {
		return utils.ResponseValidation(http.StatusNotFound, headers, "USER_EXIST")
	}

	id, err := a.loansRepository.Insert(loansDto)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	var payment = dto.PaymentDto{IdLoan: id, IdType: 1, Balance: loansDto.BorrowedValue}

	err = a.paymentsRepository.Insert(payment)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	var calculate = (loansDto.BorrowedValue) * (loansDto.InterestPercentage / 100)
	var interest = dto.InterestsDto{IdLoan: id, Status: "ACT", Share: calculate}

	err = a.interestsRepository.Insert(interest)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	return utils.ResponseValidation(http.StatusCreated, headers, "CREATED")
}

func (a *LoansServiceImpl) CreatePayment(paymentDto dto.PaymentDto, headers dto.Headers) dto.Response {

	balance, err := a.paymentsRepository.FindLastBalance(paymentDto.IdLoan)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	err = a.paymentsRepository.Insert(paymentDto)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	lastBalance, err := a.paymentsRepository.FindLastBalance(paymentDto.IdLoan)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	if paymentDto.Interest > 0 {
		err = a.paymentsRepository.UpdateBalance(lastBalance.Id, balance.Balance)
		if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
			return response
		}
	}

	if paymentDto.Capital > 0 {
		var calculteBalance = (balance.Balance - paymentDto.Capital)
		err = a.paymentsRepository.UpdateBalance(lastBalance.Id, calculteBalance)
		if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
			return response
		}

		err = a.interestsRepository.UpdateStateAllInterestByIdLoan(paymentDto.IdLoan)
		if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
			return response
		}

		var newShare = calculteBalance * (lastBalance.InterestPercentage / 100)
		var interest = dto.InterestsDto{IdLoan: paymentDto.IdLoan, Status: "ACT", Share: newShare}

		err = a.interestsRepository.Insert(interest)
		if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
			return response
		}
	}

	err = a.loansRepository.UpdateCalculateById(paymentDto.IdLoan)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	return utils.ResponseValidation(http.StatusCreated, headers, "CREATED")
}

func (a *LoansServiceImpl) FindAllLoans(headers dto.Headers) ([]dto.ListLoansDto, dto.Response) {

	var listLoans = make([]dto.ListLoansDto, 0)
	var responseDto = dto.Response{}

	loans, err := a.loansRepository.FindAllLoans()
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return listLoans, response
	}

	for _, loan := range loans {
		var listLoansDto = dto.ListLoansDto{
			Id:            loan.Id,
			Client:        loan.Client.Identification,
			Name:          loan.Client.FullName,
			BorrowedValue: loan.BorrowedValue,
			PaidValue:     loan.PaidValue,
			PendingValue:  loan.PendingValue,
			InterestPaid:  loan.InterestPaid,
			IdStatus:      loan.IdStatus,
		}

		listLoans = append(listLoans, listLoansDto)
	}

	responseDto.Status = http.StatusOK
	return listLoans, responseDto
}

func (a *LoansServiceImpl) FindByIdLoan(idLoan int32, headers dto.Headers) ([]dto.ListPaymentDto, dto.Response) {

	var listPayment = make([]dto.ListPaymentDto, 0)
	var responseDto = dto.Response{}

	payments, err := a.paymentsRepository.FindByIdLoan(idLoan)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return listPayment, response
	}

	for _, payment := range payments {
		var listPaymentDto = dto.ListPaymentDto{
			Id:          payment.Id,
			Capital:     payment.Capital,
			Interest:    payment.Interest,
			Balance:     payment.Balance,
			PaymentDate: payment.PaymentDate,
			IdType:      payment.Type.Id,
			Type:        payment.Type.Name,
		}

		listPayment = append(listPayment, listPaymentDto)
	}

	responseDto.Status = http.StatusOK
	return listPayment, responseDto
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
