package service

import (
	"Backend-Loans/domain/component/errorException"
	"Backend-Loans/domain/dto"
	persistence "Backend-Loans/infrastructure/persistence/db"
	"Backend-Loans/infrastructure/repository"
	"Backend-Loans/utils"
	"log"
	"math"
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

func (a *LoansServiceImpl) FindAllLoans(query dto.QueryParameters, headers dto.Headers) (dto.ConsultLoanDto, dto.Response) {

	var listLoans = make([]dto.ListLoansDto, 0)
	var consultLoan = dto.ConsultLoanDto{}
	var responseDto = dto.Response{}

	loans, err := a.loansRepository.FindAllLoans(query)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return consultLoan, response
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
			IdStatus:      int(loan.State.Id),
			Status:        loan.State.Description,
		}

		listLoans = append(listLoans, listLoansDto)
	}

	count, err := a.loansRepository.CountAllLoans()
	totalDouble := float64(count) / float64(6)

	consultLoan.Pages = int(math.Ceil(totalDouble))

	if query.Identification != "" {
		consultLoan.Pages = 1
	}

	consultLoan.Loans = listLoans
	responseDto.Status = http.StatusOK
	return consultLoan, responseDto
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

func (a *LoansServiceImpl) FindInformacionByLoan(idLoan int32, headers dto.Headers) (dto.InformacionUserDto, dto.Response) {

	var listInformacionUserDto = dto.InformacionUserDto{}
	var responseDto = dto.Response{}

	userInformation, err := a.loansRepository.FindInformacionByIdLoan(idLoan)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return listInformacionUserDto, response
	}

	arrears := utils.ResultCalculateMonthsArrears(userInformation.CreationDate, userInformation.State.Id)

	listInformacionUserDto = dto.InformacionUserDto{
		Id:             userInformation.Id,
		Identification: userInformation.IdentificationClient,
		FullName:       userInformation.Client.FullName,
		Address:        userInformation.Client.Address,
		Mobile:         userInformation.Client.Mobile,
		BorrowedValue:  userInformation.BorrowedValue,
		Interest:       userInformation.InterestPercentage,
		MonthlyFee:     userInformation.Interest.Share,
		LoanDate:       userInformation.CreationDate,
		MonthsArrears:  arrears,
		State:          userInformation.State.Description,
	}

	responseDto.Status = http.StatusOK
	return listInformacionUserDto, responseDto
}
