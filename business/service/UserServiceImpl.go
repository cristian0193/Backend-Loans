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

type UserServiceImpl struct {
	loansRepository     repository.LoansRepository
	paymentsRepository  repository.PaymentsRepository
	interestsRepository repository.InterestsRepository
	usersRepository     repository.UsersRepository
	errorResponse       errorException.ErrorResponse
}

func InitUserServiceImpl() *UserServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &UserServiceImpl{
		loansRepository:     dbHelper.LoansRepository,
		paymentsRepository:  dbHelper.PaymentsRepository,
		interestsRepository: dbHelper.InterestsRepository,
		usersRepository:     dbHelper.UsersRepository,
	}
}

func (a *UserServiceImpl) FindByUserAndPassword(user dto.UsersDto, headers dto.Headers) dto.Response {

	var responseDto = dto.Response{}

	users, err := a.usersRepository.FindUserAndPassword(user)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	if len(users) == 0 {
		return utils.ResponseValidation(http.StatusNotFound, headers, "ERROR_LOGIN")
	}

	responseDto.Status = http.StatusOK
	return responseDto
}
