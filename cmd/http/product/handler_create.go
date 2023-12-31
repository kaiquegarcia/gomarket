package product

import (
	"fmt"
	"gomarket/internal/usecases/producthttp"
	"gomarket/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlers) Create(ctx *gin.Context) {
	var input producthttp.CreateInput
	if err := ctx.BindJSON(&input); err != nil {
		fmt.Printf("could not decode the POST body. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "could not decode the POST body",
		})
		return
	}

	if errorList, err := util.Validate(input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "something is not right, check the errors",
			"errors":  errorList,
		})
		return
	}

	product, err := h.usecases.Create(ctx.Request.Context(), input)
	if err != nil {
		fmt.Printf("could not create the product. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not create the product, please try again later",
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("Product '%s' created successfully on the code #%d", product.Name, product.Code),
	})
}
