package producthttp

import (
	"context"
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"gomarket/pkg/format"
)

type GetInput struct {
	Code int `validate:"required,min=1"`
}

type GetOutput struct {
	Code                   int                 `json:"code"`
	Name                   string              `json:"name"`
	Materials              []GetOutputMaterial `json:"materials"`
	SellingPriceCents      int                 `json:"selling_price_cents"`
	SellingPriceMoney      string              `json:"selling_price_money"`
	QuantityPerLot         int                 `json:"quantity_per_lot"`
	CostToProduceOneCents  int                 `json:"cost_to_produce_one_cents"`
	CostToProduceOneMoney  string              `json:"cost_to_produce_one_money"`
	CostToProduceALotCents int                 `json:"cost_to_produce_a_lot_cents"`
	CostToProduceALotMoney string              `json:"cost_to_produce_a_lot_money"`
	CreatedAt              string              `json:"created_at"`
	UpdatedAt              string              `json:"updated_at"`
}

type GetOutputMaterial struct {
	Name                   string        `json:"name"`
	Unit                   enum.UnitKind `json:"unit,omitempty"`
	AmountToFabricateALot  enum.Unit     `json:"amount_to_fabricate_a_lot"`
	InvestedAmount         enum.Unit     `json:"invested_amount"`
	InvestedCents          int           `json:"invested_cents"`
	InvestedMoney          string        `json:"invested_money"`
	CostToProduceALotCents float64       `json:"cost_to_produce_a_lot_cents"`
	CostToProduceALotMoney string        `json:"cost_to_produce_a_lot_money"`
}

func (u *httpUsecases) Get(ctx context.Context, input GetInput) (*GetOutput, error) {
	p, err := u.repository.Get(input.Code)
	if err != nil {
		return nil, err
	}

	return u.newGetOutput(p)
}

func (u *httpUsecases) newGetOutput(p *entity.Product) (*GetOutput, error) {
	materials := make([]GetOutputMaterial, len(p.Materials))
	for index, m := range p.Materials {
		mProduct, err := u.repository.Get(m.ProductCode)
		if err != nil {
			return nil, err
		}

		fabricationCostCents := m.FabricationCostCents()
		materials[index] = GetOutputMaterial{
			Name:                   mProduct.Name,
			Unit:                   m.Unit,
			AmountToFabricateALot:  m.AmountToFabricate,
			InvestedAmount:         m.InvestedAmount,
			InvestedCents:          m.InvestedCents,
			InvestedMoney:          format.MoneyFromCents(float64(m.InvestedCents)),
			CostToProduceALotCents: fabricationCostCents,
			CostToProduceALotMoney: format.MoneyFromCents(fabricationCostCents),
		}
	}

	fabricationCostCents := p.FabricationCostCents()
	fabricationLotCostCents := fabricationCostCents * p.QuantityPerLot
	updatedAt := ""
	if p.UpdatedAt != nil && !p.UpdatedAt.IsZero() {
		updatedAt = p.UpdatedAt.Format(format.BrazilianDateTime)
	}
	return &GetOutput{
		Code:                   p.Code,
		Name:                   p.Name,
		Materials:              materials,
		SellingPriceCents:      p.SellingPriceCents,
		SellingPriceMoney:      format.MoneyFromCents(float64(p.SellingPriceCents)),
		QuantityPerLot:         p.QuantityPerLot,
		CostToProduceOneCents:  fabricationCostCents,
		CostToProduceOneMoney:  format.MoneyFromCents(float64(fabricationCostCents)),
		CostToProduceALotCents: fabricationLotCostCents,
		CostToProduceALotMoney: format.MoneyFromCents(float64(fabricationLotCostCents)),
		CreatedAt:              p.CreatedAt.Format(format.BrazilianDateTime),
		UpdatedAt:              updatedAt,
	}, nil
}
