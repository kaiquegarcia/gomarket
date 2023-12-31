package http

import (
	"gomarket/cmd/http/product"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	server *gin.Engine,
	productHandlers product.Handlers,
) *gin.Engine {
	// Products CRUD
	server.GET("/products", productHandlers.List)
	server.GET("/products/:code", productHandlers.Get)
	server.POST("/products/", productHandlers.Create)
	server.PUT("/products/:code", productHandlers.Update)
	server.DELETE("/products/:code", productHandlers.Delete)

	return server
}
