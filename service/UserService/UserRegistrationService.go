package UserService

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/utils"
	"github.com/gin-gonic/gin"

	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
)

func UserRegistration(c *gin.Context) {
	var reqBody model.User

	userRequest, err := utils.GetUserBody(c)
	if err != nil {
		c.JSON(constanta.CodeBadRequestResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	errValidation := userRequest.ValidationRegistration(c)
	if errValidation.Code != constanta.CodeSuccessResponse {
		c.JSON(errValidation.Code, errValidation)
		return
	}

	reqBody = mapToUser(userRequest)

	db := config.Connect()
	defer db.Disconnect(nil)

	err = dao.UserDAO.InsertUser(reqBody)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	out.ResponseOut(c, nil, true, constanta.CodeSuccessResponse, constanta.SuccessRegistrationData)
}

func mapToUser(reqBody in.UserRequest) model.User {
	return model.User{
		ID:        reqBody.Id,
		Username:  reqBody.Username,
		Password:  reqBody.Password,
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Gender:    reqBody.Gender,
		Phone:     reqBody.Telephone,
		Email:     reqBody.Email,
		Address:   reqBody.Address,
	}
}
