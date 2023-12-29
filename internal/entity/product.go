package entity

import (
	"math"
	"time"
)

// Product represents any item, for selling or not. If you use it to fabricate another product, it's a product as well.
type Product struct {
	Code              int        `json:"code"`
	Name              string     `json:"name"`
	Materials         []Material `json:"materials"`
	SellingPriceCents int        `json:"selling_price_cents"`
	QuantityPerLot    int        `json:"fabricated_products_count"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}

// FabricationCostCents will loop all to execute call Material.FabricationCostCents on the materials list, returning its sum
func (p Product) FabricationCostCents() int {
	if len(p.Materials) == 0 || p.QuantityPerLot == 0 {
		return 0
	}

	var total float64 = 0
	for _, m := range p.Materials {
		total += m.FabricationCostCents()
	}

	return int(math.Ceil(total / float64(p.QuantityPerLot)))
}
