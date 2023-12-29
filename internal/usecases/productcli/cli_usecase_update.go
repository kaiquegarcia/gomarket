package productcli

import (
	"fmt"
	"gomarket/internal/entity"
	"gomarket/internal/errs"
	"gomarket/internal/usecases/dto"
	"gomarket/pkg/util"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func (u *cliUsecases) Update() {
	var product *entity.Product = nil
	util.Try(3, func() (bool, string) {
		codeStr := util.AskCLI("what's the code of the product to be updated?")
		code, err := strconv.Atoi(codeStr)
		if err != nil {
			return false, errs.ProductCodeDecodingErr(err)
		}

		if code <= 0 {
			return false, errs.NumberShouldBeHigherThanZeroErr
		}

		product, err = u.repository.Get(code)
		if err == errs.RegistryNotFoundErr {
			return false, errs.ProductNotFoundErr
		}

		if err != nil {
			return false, errs.ProductGetErr(code, err)
		}

		return true, ""
	})

	changeOptions := []string{"done", "name", "materials", "quantityPerLot", "sellingPrice"}
	changedOptions := make(map[string]bool)
	productDTO := dto.ProductDTO{
		Name:              product.Name,
		Materials:         product.Materials,
		QuantityPerLot:    product.QuantityPerLot,
		SellingPriceCents: product.SellingPriceCents,
	}
	askedToKeepAsking := false
	for {
		availableOptions := make([]string, 0)
		for _, changeOption := range changeOptions {
			if isDifferent, isMapped := changedOptions[changeOption]; !isMapped || !isDifferent {
				availableOptions = append(availableOptions, changeOption)
			}
		}

		if len(availableOptions) == 1 && availableOptions[0] == "done" {
			availableOptions = changeOptions
			if !askedToKeepAsking {
				fmt.Println("you've changed everything that's possible")
				keepAsking := util.AskBoolCLI("do you still want to change something?")
				if !keepAsking {
					break
				}

				fmt.Println("ok, I'll show you all the options. send 'done' when you finish")
				askedToKeepAsking = true
			}
		}

		var opt string
		util.Try(3, func() (bool, string) {
			opt = strings.ToUpper(util.AskCLI(
				fmt.Sprintf(
					"what do you want to change in '%s' registry? (%s)",
					product.Name,
					strings.Join(availableOptions, "/"),
				),
			))

			if !slices.Contains([]string{"NAME", "MATERIALS", "QUANTITYPERLOT", "SELLINGPRICE", "DONE"}, opt) {
				return false, "this option is not valid, please try again"
			}

			return true, ""
		})

		switch opt {
		case "NAME":
			productDTO.Name = util.AskCLI(
				fmt.Sprintf("what's the new name of the product '%s' (#%d)?", productDTO.Name, product.Code),
			)
			changedOptions[opt] = productDTO.Name != product.Name
		case "QUANTITYPERLOT":
			util.Try(3, func() (bool, string) {
				quantityStr := util.AskCLI(
					fmt.Sprintf("what's the new quantity per fabrication lot of the product '%s' (#%d)?", productDTO.Name, product.Code),
				)
				quantity, err := strconv.Atoi(quantityStr)
				if err != nil {
					return false, errs.NumberDecodingErr(err)
				}

				if quantity < 0 {
					return false, errs.NumberShouldBeHigherOrEqualThanZeroErr
				}

				productDTO.QuantityPerLot = quantity
				changedOptions[opt] = productDTO.QuantityPerLot != product.QuantityPerLot
				return true, ""
			})
		case "SELLINGPRICE":
			util.Try(3, func() (bool, string) {
				priceStr := util.AskCLI(
					fmt.Sprintf("what's the new selling price (in cents) of the product '%s' (#%d)?", productDTO.Name, product.Code),
				)
				price, err := strconv.Atoi(priceStr)
				if err != nil {
					return false, errs.NumberDecodingErr(err)
				}

				if price < 0 {
					return false, errs.NumberShouldBeHigherOrEqualThanZeroErr
				}

				productDTO.SellingPriceCents = price
				changedOptions[opt] = productDTO.SellingPriceCents != product.SellingPriceCents
				return true, ""
			})
		case "MATERIALS":
			// TODO
			fmt.Println("this option is not implemented yet, please try other options")
			changedOptions[opt] = true
		case "DONE":
			fmt.Println("ok! let's check the changes")
			break
		}
	}

	hasChanges := false
	for _, isDifferent := range changedOptions {
		if isDifferent {
			hasChanges = true
			break
		}
	}

	if !hasChanges {
		fmt.Println("you didn't change anything, aborting the update command")
		return
	}

	fmt.Println("changes detected, updating the storage")
	updatedProduct, err := u.repository.Update(product.Code, productDTO)
	if err != nil {
		fmt.Println(errs.ProductUpdateErr(product.Code, err))
		return
	}

	fmt.Printf("product '%s' (#%d) updated successfully\n", updatedProduct.Name, product.Code)
}
