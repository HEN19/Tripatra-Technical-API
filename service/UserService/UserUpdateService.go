package UserService

import (
	"fmt"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/utils"
	"github.com/gin-gonic/gin"
)

func UserProfileUpdate(c *gin.Context) (err error) {

	now := time.Now()
	fmt.Println("HIT -> GetUserProfileService.go On ", now.Format("2006-01-02 15:04:05"))

	// Extract the Authorization header
	tokenString := c.GetHeader("Authorization")

	// Decode the token to get claims
	claims, err := config.DecodeToken(tokenString)
	if err != nil {
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	userRequest, err := utils.GetUserBody(c)
	if err != nil {
		c.JSON(constanta.CodeBadRequestResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	modelRequest := mapToUser(userRequest)
	modelRequest.ID = claims.Id

	// Connect to the database
	db := config.Connect()
	defer db.Disconnect(nil)

	_, err = dao.UserDAO.UpdateUserProfile(modelRequest)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	out.ResponseOut(c, nil, true, constanta.CodeSuccessResponse, constanta.SuccessEditData)
	return
}
