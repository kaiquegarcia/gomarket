package product

import (
	"gomarket/internal/usecases/producthttp"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type handlers struct {
	usecases producthttp.HTTP
}

func NewHandlers(usecases producthttp.HTTP) Handlers {
	return &handlers{
		usecases: usecases,
	}
}
