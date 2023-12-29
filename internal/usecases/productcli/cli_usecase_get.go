package productcli

import (
	"fmt"
	"gomarket/internal/entity"
	"gomarket/internal/errs"
	"gomarket/pkg/util"
	"strconv"
)

func (u *cliUsecases) Get() {
	var product *entity.Product
	util.Try(3, func() (bool, string) {
		codeStr := util.AskCLI("what's the product code?")
		code, err := strconv.Atoi(codeStr)
		if err != nil {
			return false, fmt.Sprintf("an error ocurred while trying to decode the product error, please try again. check the error:\n%s", err.Error())
		}

		if code <= 0 {
			return false, "please inform a number higher than zero"
		}

		product, err = u.repository.Get(code)
		if err == errs.RegistryNotFoundErr {
			return false, "could not find any product with this code, please change it and try again"
		}

		if err != nil {
			return false, fmt.Sprintf("an error ocurred while trying to retrieve the product from repository, please try again. check the error:\n%s", err.Error())
		}

		return true, ""
	})

	util.PrintLineSeparator()
	fmt.Printf("Product: %s (#%d)\n", product.Name, product.Code)
	if len(product.Materials) == 0 {
		fmt.Println("Materials: <empty>")
	} else {
		fmt.Printf("Materials: %d\n", len(product.Materials))
		for _, material := range product.Materials {
			productMaterial, err := u.repository.Get(material.ProductCode)
			if err != nil {
				fmt.Printf(
					"- CouldNotRetrieveMaterialName [#%d]: invest $%.2f to buy %.2f%s and use %.2f%s to fabricate a lot ($%.6f per lot)\n",
					material.ProductCode,
					float64(material.InvestedCents)/100,
					material.InvestedAmount,
					material.Unit,
					material.AmountToFabricate,
					material.Unit,
					material.FabricationCostCents()/100,
				)
				continue
			}

			fmt.Printf(
				"- %s [#%d]: invest $%.2f to buy %.2f%s and use %.2f%s to fabricate a lot ($%.6f per lot)\n",
				productMaterial.Name,
				material.ProductCode,
				float64(material.InvestedCents)/100,
				material.InvestedAmount,
				material.Unit,
				material.AmountToFabricate,
				material.Unit,
				material.FabricationCostCents()/100,
			)
		}
	}

	fmt.Printf("Fabrication cost: $%.2f\n", float64(product.FabricationCostCents())/100)
	fmt.Printf("Selling price: $%.2f\n", float64(product.SellingPriceCents)/100)
	util.PrintLineSeparator()
}
