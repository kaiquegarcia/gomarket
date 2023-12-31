package productcli

import (
	"fmt"
	"gomarket/pkg/util"
)

func (u *cliUsecases) List() {
	total := u.repository.Count()
	if total == 0 {
		fmt.Println("there's no products to list. please use the 'create' command first")
		return
	}

	offset := 0
	limit := 5
	for {
		list, err := u.repository.List(offset, limit)
		if err != nil {
			fmt.Printf("an error happened: %s\n", err)
			util.PrintLineSeparator()
			return
		}

		fmt.Printf("listing %d-%d of %d elements\n", offset, offset+len(list), total)
		util.PrintLineSeparator()
		fmt.Println("| code | name | fabrication cost | selling cost | created at | updated at |")
		for _, product := range list {
			fmt.Printf(
				"| %d | %s | R$%.2f | R$%.2f | %s | %s |\n",
				product.Code,
				product.Name,
				float64(product.FabricationCostCents())/100.0,
				float64(product.SellingPriceCents)/100.0,
				product.CreatedAt,
				product.UpdatedAt,
			)
		}
		util.PrintLineSeparator()

		keepGoing := false
		if offset+len(list) < total {
			keepGoing = util.AskBoolCLI("do you want to see the next page?")
		}

		if !keepGoing {
			break
		}

		offset += limit
	}
}
