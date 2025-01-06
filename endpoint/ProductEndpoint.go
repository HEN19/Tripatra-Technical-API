package endpoint

import (
	service "github.com/api-skeleton/service/ProductService"
	"github.com/gin-gonic/gin"
)

func ProductEndpointWithoutParam(c *gin.Context) {
	c.Header("Allow", "POST, GET, OPTIONS")
	switch c.Request.Method {
	case "POST":
		service.ProductService.InsertProductService(c)
		break
	case "GET":
		service.ProductService.GetListProductService(c)
		break
	}

}

func ProductEndpointWithParam(c *gin.Context) {
	c.Header("Allow", "PUT, GET, DELETE, OPTIONS")
	switch c.Request.Method {
	case "PUT":
		service.ProductService.UpdateProductService(c)
		break
	case "GET":
		service.ProductService.GetDetailProductService(c)
		break
	case "DELETE":
		service.ProductService.DeleteProductService(c)
		break
	}

}
