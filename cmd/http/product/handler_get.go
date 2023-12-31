package product

import (
	"fmt"
	"gomarket/internal/errs"
	"gomarket/internal/usecases/producthttp"
	"gomarket/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlers) Get(ctx *gin.Context) {
	var input producthttp.GetInput
	var err error
	if input.Code, err = strconv.Atoi(ctx.Param("code")); err != nil {
		fmt.Printf("could not decode the code param. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid code on path, please inform only numbers",
		})
		return
	}

	if err = ctx.BindQuery(&input); err != nil {
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

	getOutput, err := h.usecases.Get(ctx.Request.Context(), input)
	if err == errs.RegistryNotFoundErr {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "product not found",
		})
		return
	}

	if err != nil {
		fmt.Printf("could not get the product. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not get the product, please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, getOutput)
}
