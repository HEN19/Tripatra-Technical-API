package service

import (
	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"
)

func (input productService) GetListProductService(c *gin.Context) (err error) {

	db := config.Connect()
	defer db.Disconnect(nil)

	products, err := dao.ProductDAO.GetListProduct()
	if err != nil {
		out.ResponseOut(c, err, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	outProducts := convertToRepoDTOOut(products)

	out.ResponseOut(c, outProducts, true, constanta.CodeSuccessResponse, constanta.SuccessGetData)

	return

}

func convertToRepoDTOOut(products []model.Product) (productsOut []out.ProductResponse) {

	for _, product := range products {
		productsOut = append(productsOut, out.ProductResponse{
			ID:           product.ID,
			ProductName:  product.Name,
			Price:        product.Price,
			ProductStock: product.Stock,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		})
	}
	return

}
