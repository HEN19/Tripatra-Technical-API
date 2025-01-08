package endpoint

import (
	"github.com/api-skeleton/service/UserService"
	"github.com/gin-gonic/gin"
)

// func RegistrationEndpoint(response http.ResponseWriter, request *http.Request) {
// 	switch request.Method {
// 	case "POST":
// 		UserService.UserRegistration(response, request)
// 		break
// 	}
// }

func RegistrationEndpoint(c *gin.Context) {
	//  add some middleware logic here, if needed
	// Call UserRegistration service function
	c.Header("Allow", "POST, OPTIONS")
	switch c.Request.Method {
	case "POST":
		UserService.UserRegistration(c)
		break
	}

}
func UserWithParamEndpoint(c *gin.Context) {
	c.Header("Allow", "GET, PUT, OPTIONS")
	switch c.Request.Method {
	case "PUT":
		UserService.UserProfileUpdate(c)
		break
	case "GET":
		UserService.GetUserProfile(c)
		break
	}
}

func LoginEndpoint(c *gin.Context) {
	c.Header("Allow", "POST, OPTIONS")
	switch c.Request.Method {
	case "POST":
		UserService.LoginService(c)
		break
	}
}
