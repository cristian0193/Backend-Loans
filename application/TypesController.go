package application

import (
	service "Backend-Loans/business/service"
	"Backend-Loans/domain/dto"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TypesController struct {
	typeService service.TypeService
}

func InitTypesController(router *gin.Engine) {
	typesRepository := TypesController{
		typeService: service.InitTypeServiceImpl(),
	}
	router.GET("/types", typesRepository.FindAllHandler)
}

func HeadersParamType(c *gin.Context) dto.Headers {
	var headerLoans = dto.Headers{}
	headerLoans.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	return headerLoans
}

func (a *TypesController) FindAllHandler(c *gin.Context) {

	types, response := a.typeService.FindAll()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, types)
}
