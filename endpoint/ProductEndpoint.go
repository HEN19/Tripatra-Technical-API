package endpoint

import (
	"github.com/gin-gonic/gin"
)

func ProductEndpointWithoutParam(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		// UserService.InsertProduct(response, request)
		break
	case "GET":
		// UserService.GetListProduct(response, request)
		break
	}
	
}


func ProductEndpointWithParam(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		// UserService.InsertProduct(response, request)
		break
	case "GET":
		// UserService.GetListProduct(response, request)
		break
	case "DELETE":
		// UserService.DeleteProduct(response, request)
		break
	}
	
}