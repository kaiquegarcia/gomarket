package productcli

import (
	"fmt"
	"gomarket/internal/errs"
	"gomarket/pkg/util"
	"strconv"
)

func (u *cliUsecases) Delete() {
	var code int
	var err error
	util.Try(3, func() (bool, string) {
		codeStr := util.AskCLI("what's the code of the product to be deleted?")
		code, err = strconv.Atoi(codeStr)
		if err != nil {
			return false, errs.ProductCodeDecodingErr(err)
		}

		err = u.repository.Delete(code)
		if err == errs.RegistryNotFoundErr {
			return false, errs.ProductNotFoundErr
		}

		if err != nil {
			return false, errs.ProductDeleteErr(code, err)
		}

		return true, ""
	})

	fmt.Printf("product #%d deleted successfully\n", code)
}
