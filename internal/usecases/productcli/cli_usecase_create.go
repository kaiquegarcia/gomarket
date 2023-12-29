package productcli

import (
	"fmt"
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"gomarket/internal/usecases/dto"
	"gomarket/pkg/util"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func (u *cliUsecases) Create() {
	productDTO := dto.ProductDTO{}
	productDTO.Name = util.AskCLI("what's the product name?")
	util.Try(3, func() (bool, string) {
		amountStr := util.AskCLI(
			fmt.Sprintf("how many %s do you fabricate by a lot? e.g.: I can make 8 cupcakes per lot", productDTO.Name),
		)

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return false, fmt.Sprintf("an error happened while trying to decode the number, can you try again? check the error:\n%s", err.Error())
		}

		if amount <= 0 {
			return false, "please inform a number higher than zero"
		}

		productDTO.QuantityPerLot = amount
		return true, ""
	})

	materialsLength := 0
	util.Try(3, func() (bool, string) {
		lenStr := util.AskCLI(
			fmt.Sprintf("how many materials do you need to spent to fabricate a %s?", productDTO.Name),
		)
		len, err := strconv.Atoi(lenStr)
		if err != nil {
			return false, fmt.Sprintf("an error happened while decoding the number, can you try again?\ncheck the error: %s", err.Error())
		}

		if len < 0 {
			return false, "please, inform a number higher or equal than zero"
		}

		materialsLength = len
		return true, ""
	})

	productDTO.Materials = make([]entity.Material, materialsLength)
	for index := range productDTO.Materials {
		material := entity.Material{}
		isCreated := strings.ToUpper(util.AskCLI(
			fmt.Sprintf("is the #%d material registered? (Y/N)", index+1),
		)) == "Y"

		var materialProduct *entity.Product
		if !isCreated {
			fmt.Println("let's create then")
			materialProduct = u.createMaterial(index)
		} else {
			materialProduct = u.askMaterialProductID(index)
		}

		material.ProductCode = materialProduct.Code
		util.Try(3, func() (bool, string) {
			unit := strings.ToLower(util.AskCLI(
				fmt.Sprintf("what's the unit of %s? (%s/%s/%s/%s)", materialProduct.Name, enum.MASS, enum.VOLUME, enum.LENGTH, enum.UNIT),
			))

			if !slices.Contains([]enum.UnitKind{enum.MASS, enum.VOLUME, enum.LENGTH, enum.UNIT}, enum.UnitKind(unit)) {
				return false, "invalid unit, please try again"
			}

			material.Unit = enum.UnitKind(unit)
			return true, ""
		})

		util.Try(3, func() (bool, string) {
			amountStr := util.AskCLI(
				fmt.Sprintf("how many '%s' %s are required to fabricate the product lot?", material.Unit, materialProduct.Name),
			)

			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				return false, fmt.Sprintf("an error happened while trying to decode the number, can you try again? check the error:\n%s", err.Error())
			}

			if amount <= 0 {
				return false, "please inform a number higher than zero"
			}

			material.AmountToFabricate = enum.Unit(amount)
			return true, ""
		})

		util.Try(3, func() (bool, string) {
			amountStr := util.AskCLI(
				fmt.Sprintf("how many '%s' %s do you buy at once? e.g.: I buy 1000ml of milk", material.Unit, materialProduct.Name),
			)

			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				return false, fmt.Sprintf("an error happened while trying to decode the number, can you try again? check the error:\n%s", err.Error())
			}

			if amount <= 0 {
				return false, "please inform a number higher than zero"
			}

			material.InvestedAmount = enum.Unit(amount)
			return true, ""
		})

		util.Try(3, func() (bool, string) {
			costCentsStr := util.AskCLI(
				fmt.Sprintf("how much cost to buy %.2f%s of %s (in cents)?", material.InvestedAmount, material.Unit, materialProduct.Name),
			)

			costCents, err := strconv.Atoi(costCentsStr)
			if err != nil {
				return false, fmt.Sprintf("an error happened while trying to decode the number, can you try again? check the error:\n%s", err.Error())
			}

			if costCents < 0 {
				return false, "please inform a number higher or equal than zero"
			}

			material.InvestedCents = costCents
			return true, ""
		})

		productDTO.Materials[index] = material
		fmt.Println("alright, let's keep going")
	}

	fmt.Println("last question...")
	util.Try(3, func() (bool, string) {
		sellingPriceCentsStr := util.AskCLI(
			fmt.Sprintf("what's the selling price of %s?", productDTO.Name),
		)

		sellingPriceCents, err := strconv.Atoi(sellingPriceCentsStr)
		if err != nil {
			return false, fmt.Sprintf("an error happened while trying to decode the number, can you try again? check the error:\n%s", err.Error())
		}

		if sellingPriceCents < 0 {
			return false, "please inform a number higher or equal than zero"
		}

		productDTO.SellingPriceCents = sellingPriceCents
		return true, ""
	})

	product, err := u.repository.Insert(productDTO)
	if err != nil {
		fmt.Printf("could not register the product. please try again later. check the error:\n%s\n", err.Error())
	} else {
		fmt.Printf("product '%s' registered on code #%d successfully", product.Name, product.Code)
	}

	util.FinishCLI()
}

func (u *cliUsecases) createMaterial(index int) *entity.Product {
	productDTO := dto.ProductDTO{
		Materials: make([]entity.Material, 0),
	}
	productDTO.Name = util.AskCLI(
		fmt.Sprintf("what's the name of the #%d material?", index+1),
	)
	entity, err := u.repository.Insert(productDTO)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("material %s (code: #%d) created\n", entity.Name, entity.Code)
	return entity
}

func (u *cliUsecases) askMaterialProductID(index int) *entity.Product {
	knowCode := strings.ToUpper(util.AskCLI(
		fmt.Sprintf("do you know the product code of the #%d material?", index+1),
	)) == "Y"

	var product *entity.Product = nil
	if knowCode {
		util.Try(3, func() (bool, string) {
			codeStr := util.AskCLI(
				fmt.Sprintf("what's the product code of the #%d material?", index+1),
			)
			code, err := strconv.Atoi(codeStr)
			if err != nil {
				return false, fmt.Sprintf("I could not decode the product code, can you try again?\ndecoding error: %s", err.Error())
			}

			if code <= 0 {
				return false, "please, inform a number higher than zero"
			}

			product, err = u.repository.Get(code)
			if err != nil {
				return false, fmt.Sprintf("could not get the product from storage, are you sure this code is right? check the error:\n%s", err.Error())
			}

			confirmation := strings.ToUpper(util.AskCLI(
				fmt.Sprintf("is '%s' the #%d material? (Y/N)", product.Name, index+1),
			)) == "Y"

			if !confirmation {
				return false, "ok... let's try again"
			}

			return true, ""
		})

		return product
	}

	nameToSearch := strings.ToLower(util.AskCLI(
		fmt.Sprintf("let's search then! what's the name of the #%d material?", index+1),
	))

	util.Try(3, func() (bool, string) {
		offset := 0
		limit := 20
		for {
			list, err := u.repository.List(offset, limit)
			if err != nil {
				return false, fmt.Sprintf("an error happened, can you try again? check the error:\n%s", err.Error())
			}

			if len(list) == 0 {
				break
			}

			for _, material := range list {
				if !strings.Contains(strings.ToLower(material.Name), nameToSearch) {
					continue
				}

				confirmation := strings.ToUpper(util.AskCLI(
					fmt.Sprintf("is it '%s' (code #%d)? (Y/N)", material.Name, material.Code),
				)) == "Y"
				if confirmation {
					product = material
					return true, ""
				}
			}

			offset += limit
		}

		return true, ""
	})

	if product != nil {
		return product
	}

	restart := strings.ToUpper(util.AskCLI(
		fmt.Sprintf("could not find any material with this name. do you want to restart the selection flow for the #%d material? (Y/N)\nPS.: if you abort, the entire product creation will be aborted.", index+1),
	)) == "Y"

	if restart {
		return u.askMaterialProductID(index)
	}

	util.FinishCLI()
	panic("")
}
