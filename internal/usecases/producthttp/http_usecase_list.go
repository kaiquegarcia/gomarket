package producthttp

import (
	"context"
	"gomarket/pkg/format"
)

type ListInput struct {
	Page int `query:"page" form:"page" validate:"min=1"`
}

type ListOutput struct {
	Code            int    `json:"code"`
	Name            string `json:"name"`
	FabricationCost string `json:"fabrication_cost"`
	SellingPrice    string `json:"selling_price"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

const listLimit = 25

func (u *httpUsecases) List(ctx context.Context, input ListInput) ([]ListOutput, error) {
	offset := (input.Page - 1) * listLimit

	list, err := u.repository.List(offset, listLimit)
	if err != nil {
		return make([]ListOutput, 0), err
	}

	output := make([]ListOutput, len(list))
	for index, p := range list {
		updatedAt := ""
		if p.UpdatedAt != nil && !p.UpdatedAt.IsZero() {
			updatedAt = p.UpdatedAt.Format(format.BrazilianDateTime)
		}

		output[index] = ListOutput{
			Code:            p.Code,
			Name:            p.Name,
			FabricationCost: format.MoneyFromCents(float64(p.FabricationCostCents())),
			SellingPrice:    format.MoneyFromCents(float64(p.SellingPriceCents)),
			CreatedAt:       p.CreatedAt.Format(format.BrazilianDateTime),
			UpdatedAt:       updatedAt,
		}
	}

	return output, nil
}
