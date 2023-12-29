package errs

import "fmt"

const (
	anErrorOccurredWhile string = "an error occurred while trying to %s, please try again. check the error:\n%s"
)

func NumberDecodingErr(err error) string {
	return fmt.Sprintf(anErrorOccurredWhile, "decode this number", err.Error())
}

func ProductCodeDecodingErr(err error) string {
	return fmt.Sprintf(anErrorOccurredWhile, "decode the product code", err.Error())
}

func ProductGetErr(productCode int, err error) string {
	return fmt.Sprintf(
		anErrorOccurredWhile,
		fmt.Sprintf("retrieve the product #%d from the storage", productCode),
		err.Error(),
	)
}

func ProductListErr(err error) string {
	return fmt.Sprintf(anErrorOccurredWhile, "list products from the storage", err.Error())
}

func ProductDeleteErr(productCode int, err error) string {
	return fmt.Sprintf(
		anErrorOccurredWhile,
		fmt.Sprintf("delete the product #%d", productCode),
		err.Error(),
	)
}
