package application

import (
	"Backend-Loans/business/service"
	"Backend-Loans/domain/dto"
	"Backend-Loans/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var serviceName string

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
}

func HeadersParamLoans(c *gin.Context) dto.Headers {
	var headerLoans = dto.Headers{}
	headerLoans.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	return headerLoans
}

func QueryParamLoans(c *gin.Context) dto.QueryParameters {
	var queryParameters = dto.QueryParameters{}
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
	c.JSON(http.StatusAccepted, response)
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
	c.JSON(http.StatusAccepted, response)
}

func (a *LoansController) FindAllHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)

	loans, response := a.loansService.FindAllLoans(headers)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusAccepted, loans)
}

/* func (p *LoansController) GetAllProductHandler(c *gin.Context) {
	var headers = HeadersParamProduct(c)

	response, product := p.productService.GetByIdMarketProduct(headers)
	if response.Status != http.StatusOK {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusOK, product)
}

func (a *ProductController) UpdateProductHandler(c *gin.Context) {
	var headers = HeadersParamProduct(c)
	var productDto dto.ProductDto
	var responseDto = dto.Response{}

	if err := c.ShouldBindJSON(&productDto); err != nil {
		responseDto.Status = http.StatusUnprocessableEntity
		responseDto.Description = utils.StatusText(http.StatusUnprocessableEntity)
		responseDto.Message = utils.Lenguage(headers.Lenguage, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.productService.UpdateProduct(productDto, headers)

	if response.Status != http.StatusCreated {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusAccepted, response)
}

func (a *ProductController) DeleteProductHandler(c *gin.Context) {
	var headers = HeadersParamProduct(c)
	var idProduct = c.Param("idProduct")

	response := a.productService.DeleteProduct(idProduct, headers)

	if response.Status != http.StatusOK {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusAccepted, response)
}

func (a *ProductController) GetByIdProductHandler(c *gin.Context) {
	var headers = HeadersParamProduct(c)
	var idProduct = c.Param("idProduct")

	response, product := a.productService.GetByIdProduct(idProduct, headers)

	if response.Status != http.StatusOK {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusAccepted, product)
}

func (a *ProductController) GetByQueryParametersHandler(c *gin.Context) {
	var headers = HeadersParamProduct(c)
	var queryParameters = QueryParamProduct(c)

	response, product := a.productService.GetByQueryParameters(queryParameters, headers)

	if response.Status != http.StatusOK {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusAccepted, product)
}
*/
