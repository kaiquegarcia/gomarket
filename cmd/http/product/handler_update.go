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

func (h *handlers) Update(ctx *gin.Context) {
	var input producthttp.UpdateInput
	var err error
	if input.Code, err = strconv.Atoi(ctx.Param("code")); err != nil {
		fmt.Printf("could not decode the code param. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid code on path, please inform only numbers",
		})
		return
	}

	if err = ctx.BindJSON(&input); err != nil {
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

	product, err := h.usecases.Update(ctx.Request.Context(), input)
	if err == errs.RegistryNotFoundErr {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "product not found",
		})
		return
	}

	if err != nil {
		fmt.Printf("could not update the product. check the error:\n%s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not update the product, please try again later",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Product '%s' (#%d) updated successfully", product.Name, product.Code),
	})
}
