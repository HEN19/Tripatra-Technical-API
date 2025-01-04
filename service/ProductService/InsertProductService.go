package service

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (input productService) InsertProductService(c *gin.Context) (err error) {

	var (
		reqBody     model.Product
		inputStruct in.ProductRequest
		result      *mongo.InsertOneResult
	)

	inputStruct, err = input.ReadBody(c)
	if err != nil {
		return err
	}

	reqBody = mapToProduct(inputStruct)

	db := config.Connect()
	defer db.Disconnect(nil)

	result, err = dao.ProductDAO.InsertProduct(reqBody)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	} else {
		c.JSON(constanta.CodeSuccessResponse, result)
	}

	return

}

func mapToProduct(inputStruct in.ProductRequest) model.Product {
	return model.Product{
		ID:    inputStruct.ID,
		Name:  inputStruct.Name,
		Price: inputStruct.Price,
		Stock: inputStruct.Stock,
	}
}
