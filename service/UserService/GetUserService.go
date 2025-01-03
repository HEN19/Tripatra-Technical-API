package UserService

import (
	"fmt"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) (err error) {
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

	// Connect to the database
	db := config.Connect()
	defer db.Disconnect(nil)

	// Get user profile from the DAO layer using the decoded user ID
	user, err := dao.UserDAO.GetUserProfile(claims.Id)
	if err != nil {
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	// Convert the repository model to DTO for the response
	outUser := convertRepoToDTOID(user)

	// Send the response
	out.ResponseOut(c, outUser, true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
	return nil
}

// convertRepoToDTOID converts the UserModel to a UserRequest DTO for response.
func convertRepoToDTOID(datas model.User) out.UserRequest {
	return out.UserRequest{
		FirstName: datas.FirstName,
		LastName:  datas.LastName,
		Email:     datas.Email,
		Address:   datas.Address,
		Gender:    datas.Gender,
		Telephone: datas.Phone,
		CreatedAt: datas.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: datas.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
