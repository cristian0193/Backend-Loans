package application

import (
	service "Backend-Loans/business/service/clients"
	"Backend-Loans/domain/dto"
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
