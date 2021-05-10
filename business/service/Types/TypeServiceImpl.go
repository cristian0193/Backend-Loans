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

type TypeServiceImpl struct {
	typesRepository repository.TypesRepository
	errorResponse   errorException.ErrorResponse
}

func InitTypeServiceImpl() *TypeServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &TypeServiceImpl{
		typesRepository: dbHelper.TypesRepository,
	}
}

func (a *TypeServiceImpl) FindAll() ([]dto.TypesDto, dto.Response) {

	var listType = make([]dto.TypesDto, 0)
	var responseDto = dto.Response{}

	types, err := a.typesRepository.FindAll()
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return listType, response
	}

	for _, typ := range types {
		var listTypeDto = dto.TypesDto{
			Id:   typ.Id,
			Name: typ.Name,
		}

		listType = append(listType, listTypeDto)
	}

	responseDto.Status = http.StatusOK
	return listType, responseDto
}
