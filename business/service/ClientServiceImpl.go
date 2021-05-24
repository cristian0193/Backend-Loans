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

func (a *ClientServiceImpl) FindById(id int32) (dto.ClientsDto, dto.Response) {

	var client = dto.ClientsDto{}
	var responseDto = dto.Response{}

	clients, err := a.clientsRepository.FindById(id)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return client, response
	}

	client.Identification = clients.Identification
	client.FullName = clients.FullName
	client.Address = clients.Address
	client.Mobile = clients.Mobile

	responseDto.Status = http.StatusOK
	return client, responseDto
}

func (a *ClientServiceImpl) Create(clientDto dto.ClientsDto, headers dto.Headers) dto.Response {

	if clientDto.Identification == 0 {
		return utils.ResponseValidation(http.StatusNotFound, headers, "CLIENT_NOT_EXIST")
	}

	err := a.clientsRepository.Create(clientDto)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	return utils.ResponseValidation(http.StatusCreated, headers, "CREATED")
}

func (a *ClientServiceImpl) Update(clientDto dto.ClientsDto, headers dto.Headers) dto.Response {

	err := a.clientsRepository.Update(clientDto)
	if response := utils.ResponseError(http.StatusBadRequest, err); response.Status != http.StatusOK {
		return response
	}

	return utils.ResponseValidation(http.StatusCreated, headers, "UPDATED")
}
