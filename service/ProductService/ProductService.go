package service

import (
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"
)

type productService struct {
}

var ProductService = productService{}.New()

func (input productService) New() (output productService) {
	return
}

func (input productService) ReadBody(c *gin.Context) (in.ProductRequest, error) {
	var userRequest in.ProductRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		return userRequest, err
	}
	return userRequest, nil
}

func mapToProduct(inputStruct in.ProductRequest) model.Product {
	return model.Product{
		ID:    inputStruct.ID,
		Name:  inputStruct.Name,
		Price: inputStruct.Price,
		Stock: inputStruct.Stock,
	}
}
