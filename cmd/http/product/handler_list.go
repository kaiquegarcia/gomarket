package product

import (
	"fmt"
	"gomarket/internal/usecases/producthttp"
	"gomarket/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlers) List(ctx *gin.Context) {
	var input producthttp.ListInput
	input.Page = 1
	if err := ctx.BindQuery(&input); err != nil {
		fmt.Printf("could not decode the query parameters. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "could not decode the query parameters",
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

	products, err := h.usecases.List(ctx.Request.Context(), input)
	if err != nil {
		fmt.Printf("could not list the products. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not list the products, please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
