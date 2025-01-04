package service

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dao"
	"github.com/gin-gonic/gin"
)

func (input productService) GetListProductService(c *gin.Context) (err error) {

	db := config.Connect()
	defer db.Disconnect(nil)

	products, err := dao.ProductDAO.GetListProduct()
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	c.JSON(constanta.CodeSuccessResponse, products)

	return

}
