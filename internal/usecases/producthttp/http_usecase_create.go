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
}

func (u *httpUsecases) Create(ctx context.Context, input CreateInput) (*entity.Product, error) {
	dto, err := newProductDTO(input)
	if err != nil {
		return nil, err
	}

	return u.repository.Insert(*dto)
}

func newProductDTO(input CreateInput) (*dto.ProductDTO, error) {
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

		dto.Materials[index] = entity.Material{
			ProductCode:       m.ProductCode,
			Unit:              enum.UnitKind(m.Unit),
			AmountToFabricate: enum.Unit(m.AmountToFabricate),
			InvestedAmount:    enum.Unit(m.InvestedAmount),
			InvestedCents:     m.InvestedCents,
		}
	}

	return &dto, nil
}
