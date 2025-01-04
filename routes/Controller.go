package routes

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/api-skeleton/endpoint"
	"github.com/api-skeleton/graph"
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	// routes := mux.NewRouter()

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")
	routes := gin.Default()

	user := routes.Group("v1/user")
	{
		user.POST("/register", endpoint.RegistrationEndpoint)
		user.POST("/login", endpoint.LoginEndpoint)
		user.GET("/profile", endpoint.UserWithParamEndpoint)
	}

	product := routes.Group("v1/product")
	{
		product.POST("/", endpoint.ProductEndpointWithoutParam)
		product.GET("/", endpoint.ProductEndpointWithoutParam)
		product.GET("/{id}", endpoint.ProductEndpointWithParam)
		product.PUT("/{id}", endpoint.ProductEndpointWithParam)
		product.DELETE("/{id}", endpoint.ProductEndpointWithParam)
	}

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")

	return routes
}
