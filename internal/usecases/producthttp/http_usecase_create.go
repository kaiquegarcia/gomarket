package producthttp

import (
	"context"
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"gomarket/internal/errs"
	"gomarket/internal/usecases/dto"

	"golang.org/x/exp/slices"
)

type CreateInput struct {
	Name              string                `json:"name" form:"name" validate:"required"`
	SellingPriceCents int                   `json:"selling_price_cents" form:"selling_price_cents" validate:"min=0"`
	QuantityPerLot    int                   `json:"quantity_per_lot" form:"quantity_per_lot" validate:"min=0"`
	Materials         []CreateInputMaterial `json:"materials" form:"materials" validate:"required"`
}

type CreateInputMaterial struct {
	ProductCode       int     `json:"product_code" form:"product_code" validate:"required,min=0"`
	Unit              string  `json:"unit" form:"unit"`
	AmountToFabricate float64 `json:"amount_to_fabricate" form:"amount_to_fabricate" validate:"required,min=0.01"`
	InvestedAmount    float64 `json:"invested_amount" form:"invested_amount" validate:"required,min=0.01"`
	InvestedCents     int     `json:"invested_cents" form:"invested_cents" validate:"required,min=1"`
	// TODO: add FabricationUnitID to allow create/update informing lower numbers (ex: 2L instead of 2000ml)
	// TODO: add InvestUnitID
}

func (u *httpUsecases) Create(ctx context.Context, input CreateInput) (*entity.Product, error) {
	dto, err := u.newProductDTO(input, 0)
	if err != nil {
		return nil, err
	}

	return u.repository.Insert(*dto)
}

func (u *httpUsecases) newProductDTO(input CreateInput, currentCode int) (*dto.ProductDTO, error) {
	dto := dto.ProductDTO{
		Name:              input.Name,
		Materials:         make([]entity.Material, len(input.Materials)),
		SellingPriceCents: input.SellingPriceCents,
		QuantityPerLot:    input.QuantityPerLot,
	}

	for index, m := range input.Materials {
		if !slices.Contains(enum.UnitKinds, m.Unit) {
			return nil, errs.InvalidUnitValidationErr
		}

		if m.ProductCode == currentCode {
			return nil, errs.MaterialCodeCantBeProductCodeErr
		}

		_, err := u.repository.Get(m.ProductCode)
		if err == errs.RegistryNotFoundErr {
			return nil, errs.MaterialNotFoundErr
		} else if err != nil {
			return nil, err
		}

		kind := enum.UnitKind(m.Unit)
		unitID := enum.DefaultUnitID(kind)
		// TODO: receive FabricationUnitID and InvestUnitID from input instead of using default unit ID
		dto.Materials[index] = entity.Material{
			ProductCode:       m.ProductCode,
			Unit:              kind,
			AmountToFabricate: enum.Unit(m.AmountToFabricate),
			FabricationUnitID: unitID,
			InvestedAmount:    enum.Unit(m.InvestedAmount),
			InvestUnitID:      unitID,
			InvestedCents:     m.InvestedCents,
		}
	}

	return &dto, nil
}
