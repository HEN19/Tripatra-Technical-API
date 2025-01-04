package service

import (
	"github.com/api-skeleton/dto/in"
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
