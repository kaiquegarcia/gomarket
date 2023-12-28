package dto

import (
	"gomarket/internal/entity"
	"time"
)

type ProductDTO struct {
	Name              string            `json:"name"`
	Materials         []entity.Material `json:"materials"`
	SellingPriceCents int               `json:"selling_price_cents"`
}

func (dto ProductDTO) ToEntity(
	code int,
	createdAt time.Time,
	updatedAt *time.Time,
) entity.Product {
	return entity.Product{
		Code:              code,
		Name:              dto.Name,
		Materials:         dto.Materials,
		SellingPriceCents: dto.SellingPriceCents,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}
}
