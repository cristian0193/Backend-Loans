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

type ProductController struct {
	productService service.ProductService
}

func InitProductController(router *gin.Engine) {
	productRepository := ProductController{
		productService: service.InitProductServiceImpl(),
	}
	router.GET("/product", productRepository.GetAllProductHandler)
	router.POST("/product", productRepository.CreateProductHandler)
	router.PUT("/product", productRepository.UpdateProductHandler)
	router.DELETE("/product/:idProduct", productRepository.DeleteProductHandler)
	router.GET("/product/:idProduct", productRepository.GetByIdProductHandler)
	router.GET("/products/parameters", productRepository.GetByQueryParametersHandler)

	InitServiceName()
}

func InitServiceName() {
	serviceName = utils.GetStrEnv("SERVICE_NAME")
}

func HeadersParamProduct(c *gin.Context) dto.Headers {
	var headerIncident = dto.Headers{}
	headerIncident.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	headerIncident.IdMarket = c.Request.Header.Get(os.Getenv("ID_MARKET_HEADER"))
	return headerIncident
}

func QueryParamProduct(c *gin.Context) dto.QueryParameters {
	var queryParameters = dto.QueryParameters{}
	queryParameters.Name = c.Request.URL.Query().Get("name")
	queryParameters.Price = c.Request.URL.Query().Get("price")
	queryParameters.IdCategory = c.Request.URL.Query().Get("id_category")
	return queryParameters
}

func (p *ProductController) GetAllProductHandler(c *gin.Context) {
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

func (a *ProductController) CreateProductHandler(c *gin.Context) {
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

	response := a.productService.CreateProduct(productDto, headers)

	if response.Status != http.StatusCreated {
		utils.Trace(serviceName, c, response.Status, response)
		c.JSON(response.Status, response)
		return
	}
	utils.Trace(serviceName, c, http.StatusOK, response)
	c.JSON(http.StatusAccepted, response)
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
