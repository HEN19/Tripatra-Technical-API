package UserService

import (
	"fmt"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"

	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/utils"
)

func LoginService(c *gin.Context) (err error) {
	var user model.User

	now := time.Now()
	fmt.Println("HIT -> LoginService.go On ", now.Format("2006-01-02 15:04:05"))

	// Get user request body from Gin context
	userBody, err := utils.GetUserBody(c)

	errValidation := userBody.ValidationRegistration(c)
	if errValidation.Code != constanta.CodeSuccessResponse {
		c.JSON(errValidation.Code, errValidation)
		return
	}

	userRepo := mapToUser(userBody)

	db := config.Connect()
	defer db.Disconnect(nil)

	user, err = dao.UserDAO.LoginCheck(userRepo)
	if err != nil {
		c.JSON(constanta.CodeBadRequestResponse, constanta.ErrorInternalDB)
		return
	}

	if user.ID == 0 {
		c.JSON(constanta.CodeBadRequestResponse, constanta.ErrorDataUnknown)
		return
	}

	token, err := config.GenerateToken(user)
	if err != nil {
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	// Respond with the generated token
	out.ResponseOut(c, token, true, constanta.CodeSuccessResponse, "Login berhasil")
	return nil
}
