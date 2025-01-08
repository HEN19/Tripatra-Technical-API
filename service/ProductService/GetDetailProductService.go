package service

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/gin-gonic/gin"
)

func (input productService) GetDetailProductService(c *gin.Context) (err error) {

	id := c.Param("id")
	db := config.Connect()
	defer db.Disconnect(nil)

	product, err := dao.ProductDAO.GetDetailProduct(id)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	c.JSON(constanta.CodeSuccessResponse, product)

	return

}
