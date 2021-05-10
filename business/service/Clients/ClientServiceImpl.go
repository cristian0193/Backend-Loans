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

type ClientServiceImpl struct {
	clientsRepository repository.ClientsRepository
	errorResponse     errorException.ErrorResponse
}

func InitClientServiceImpl() *ClientServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ClientServiceImpl{
		clientsRepository: dbHelper.ClientsRepository,
	}
}

func (a *ClientServiceImpl) FindAll() ([]dto.ClientsDto, dto.Response) {

	var listClient = make([]dto.ClientsDto, 0)
	var responseDto = dto.Response{}

	clients, err := a.clientsRepository.FindAll()
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return listClient, response
	}

	for _, client := range clients {
		var listClientDto = dto.ClientsDto{
			Identification: client.Identification,
			FullName:       client.FullName,
			Address:        client.Address,
			Mobile:         client.Mobile,
		}

		listClient = append(listClient, listClientDto)
	}

	responseDto.Status = http.StatusOK
	return listClient, responseDto
}
