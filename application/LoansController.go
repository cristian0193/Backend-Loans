package application

import (
	service "Backend-Loans/business/service/loans"
	"Backend-Loans/domain/dto"
	"Backend-Loans/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoansController struct {
	loansService service.LoansService
}

func InitLoansController(router *gin.Engine) {
	loansRepository := LoansController{
		loansService: service.InitLoansServiceImpl(),
	}
	router.POST("/loans", loansRepository.CreateLoansHandler)
	router.POST("/loans/payment", loansRepository.CreatePaymentHandler)
	router.GET("/loans", loansRepository.FindAllHandler)
	router.GET("/loans/historial/:idLoan", loansRepository.FindByIdLoanHandler)
	router.GET("/loans/information/:idLoan", loansRepository.FindInformationLoanHandler)
}

func HeadersParamLoans(c *gin.Context) dto.Headers {
	var headerLoans = dto.Headers{}
	headerLoans.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	return headerLoans
}

func QueryParamLoans(c *gin.Context) dto.QueryParameters {
	var queryParameters = dto.QueryParameters{}
	queryParameters.Pages = uint(utils.GetInt(c.Request.URL.Query().Get("page")))
	queryParameters.Identification = c.Request.URL.Query().Get("identification")
	return queryParameters
}

func (a *LoansController) CreateLoansHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var loansDto dto.LoansDto

	if err := c.ShouldBindJSON(&loansDto); err != nil {
		responseDto := utils.ResponseValidation(http.StatusUnprocessableEntity, headers, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.loansService.CreateLoan(loansDto, headers)

	if response.Status != http.StatusCreated {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, response)
}

func (a *LoansController) CreatePaymentHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var paymentDto dto.PaymentDto

	if err := c.ShouldBindJSON(&paymentDto); err != nil {
		responseDto := utils.ResponseValidation(http.StatusUnprocessableEntity, headers, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.loansService.CreatePayment(paymentDto, headers)

	if response.Status != http.StatusCreated {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, response)
}

func (a *LoansController) FindAllHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var query = QueryParamLoans(c)

	loans, response := a.loansService.FindAllLoans(query, headers)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, loans)
}

func (a *LoansController) FindByIdLoanHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var idLoan = utils.ConvertInt32(c.Param("idLoan"))

	payments, response := a.loansService.FindByIdLoan(idLoan, headers)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, payments)
}

func (a *LoansController) FindInformationLoanHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var idLoan = utils.ConvertInt32(c.Param("idLoan"))

	information, response := a.loansService.FindInformacionByLoan(idLoan, headers)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, information)
}
