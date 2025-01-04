package service

import (
	"strconv"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/gin-gonic/gin"
)

func (input productService) DeleteProductService(c *gin.Context) (err error) {

	id := c.Param("id")
	db := config.Connect()
	defer db.Disconnect(nil)

	idProduct, errParse := strconv.ParseInt(id, 10, 64)
	if errParse != nil {
		return errParse
	}

	product, err := dao.ProductDAO.DeleteProduct(idProduct)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	c.JSON(constanta.CodeSuccessResponse, product)

	return

}
