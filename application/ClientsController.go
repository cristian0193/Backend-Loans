package application

import (
	service "Backend-Loans/business/service/clients"
	"Backend-Loans/domain/dto"
	"Backend-Loans/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ClientsController struct {
	clientService service.ClientService
}

func InitClientsController(router *gin.Engine) {
	clientRepository := ClientsController{
		clientService: service.InitClientServiceImpl(),
	}
	router.GET("/clients", clientRepository.FindAllHandler)
	router.GET("/clients/:id", clientRepository.FindByIdHandler)
	router.POST("/clients", clientRepository.CreateClientHandler)
	router.PUT("/clients", clientRepository.UpdateClientHandler)
}

func HeadersParamClient(c *gin.Context) dto.Headers {
	var headerLoans = dto.Headers{}
	headerLoans.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	return headerLoans
}

func (a *ClientsController) FindAllHandler(c *gin.Context) {

	clients, response := a.clientService.FindAll()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, clients)
}

func (a *ClientsController) FindByIdHandler(c *gin.Context) {
	var id = utils.ConvertInt32(c.Param("id"))

	clients, response := a.clientService.FindById(id)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, clients)
}

func (a *ClientsController) CreateClientHandler(c *gin.Context) {

	var headers = HeadersParamLoans(c)
	var clientsDto dto.ClientsDto

	if err := c.ShouldBindJSON(&clientsDto); err != nil {
		responseDto := utils.ResponseValidation(http.StatusUnprocessableEntity, headers, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.clientService.Create(clientsDto, headers)

	if response.Status != http.StatusCreated {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, response)
}

func (a *ClientsController) UpdateClientHandler(c *gin.Context) {

	var headers = HeadersParamLoans(c)
	var clientsDto dto.ClientsDto

	if err := c.ShouldBindJSON(&clientsDto); err != nil {
		responseDto := utils.ResponseValidation(http.StatusUnprocessableEntity, headers, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.clientService.Update(clientsDto, headers)

	if response.Status != http.StatusCreated {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, response)
}
