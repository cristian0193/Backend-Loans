package application

import (
	service "Backend-Loans/business/service"
	"Backend-Loans/domain/dto"
	"Backend-Loans/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func InitUserController(router *gin.Engine) {
	usersRepository := UserController{
		userService: service.InitUserServiceImpl(),
	}
	router.POST("/users", usersRepository.ValidationUserAndPasswordHandler)
}

func HeadersParamUser(c *gin.Context) dto.Headers {
	var headerLoans = dto.Headers{}
	headerLoans.Lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
	return headerLoans
}

func (a *UserController) ValidationUserAndPasswordHandler(c *gin.Context) {
	var headers = HeadersParamLoans(c)
	var usersDto dto.UsersDto

	if err := c.ShouldBindJSON(&usersDto); err != nil {
		responseDto := utils.ResponseValidation(http.StatusUnprocessableEntity, headers, "BODY_INVALID")
		c.JSON(http.StatusUnprocessableEntity, responseDto)
		return
	}

	response := a.userService.FindByUserAndPassword(usersDto, headers)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, response)
}
