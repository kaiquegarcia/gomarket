package product

import (
	"fmt"
	"gomarket/internal/entity"
	"gomarket/internal/usecases/dto"
	"gomarket/pkg/util"
)

func (u *cliUsecases) Create() {
	// TODO
}

func (u *cliUsecases) createMaterial() int {
	productDTO := dto.ProductDTO{
		Materials: make([]entity.Material, 0),
	}
	productDTO.Name = util.Ask("what's the material name?")
	entity, err := u.repository.Insert(productDTO)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("material %s (#%d) created\n", entity.Name, entity.Code)
	return entity.Code
}
