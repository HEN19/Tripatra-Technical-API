package routes

import (
	"net/http"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/endpoint"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Controller() *gin.Engine {
	// routes := mux.NewRouter()

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")
	routes := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}                   // Allow requests from your frontend domain
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Allow methods
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Allow headers

	// Use the CORS middleware with the specified config
	routes.Use(cors.New(corsConfig))
	routes.Use(removeTrailingSlash())

	user := routes.Group("/v1/user")
	{
		user.POST("/register", endpoint.RegistrationEndpoint)
		user.POST("/login", endpoint.LoginEndpoint)
		user.GET("/profile", endpoint.UserWithParamEndpoint)
		user.PUT("/profile", endpoint.UserWithParamEndpoint)

		//CORS
		user.OPTIONS("/register", func(c *gin.Context) {
			c.Header("Allow", "POST, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "OPTIONS request handled"})
		})
		user.OPTIONS("/login", func(c *gin.Context) {
			c.Header("Allow", "POST, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "OPTIONS request handled"})
		})
		user.OPTIONS("/profile", func(c *gin.Context) {
			c.Header("Allow", "GET, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "OPTIONS request handled"})
		})
	}

	product := routes.Group("/v1/product")
	{
		product.POST("", config.AuthMiddleware(), endpoint.ProductEndpointWithoutParam)
		product.GET("", config.AuthMiddleware(), endpoint.ProductEndpointWithoutParam)
		product.GET("/:id", config.AuthMiddleware(), endpoint.ProductEndpointWithParam)
		product.PUT("/:id", config.AuthMiddleware(), endpoint.ProductEndpointWithParam)
		product.DELETE("/:id", config.AuthMiddleware(), endpoint.ProductEndpointWithParam)

		//CORS
		product.OPTIONS("", func(c *gin.Context) {
			c.Header("Allow", "POST, GET, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "OPTIONS request handled"})
		})
		product.OPTIONS("/:id", func(c *gin.Context) {
			c.Header("Allow", "PUT, GET, DELETE, OPTIONS")
			c.JSON(http.StatusOK, gin.H{"message": "OPTIONS request handled"})
		})
	}

	// user := routes.PathPrefix("/user").Subrouter()
	// user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	// user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")

	return routes
}

func removeTrailingSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' {
			c.Request.URL.Path = path[:len(path)-1]
		}
		c.Next()
	}
}

// Add this middleware to your routes
