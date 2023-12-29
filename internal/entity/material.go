package entity

import (
	"gomarket/internal/enum"
)

// Material represents a product required to produce other product. With this struct is possible to calculate the total cost to fabricate a single unit of the product.
type Material struct {
	ProductCode       int           `json:"product_code"`
	Unit              enum.UnitKind `json:"unit"`
	AmountToFabricate enum.Unit     `json:"amount_to_fabricate"`
	InvestedAmount    enum.Unit     `json:"invested_amount"`
	InvestedCents     int           `json:"invested_cents"`
}

// FabricationCostCents will calculate the cost to use this Material to fabricate a single product, based on the numbers defined in {AmountToFabricate}, {InvestedAmount}, {InvestedCents} and {FabricatedProductCount}
func (m Material) FabricationCostCents() float64 {
	// Math:
	// CostToFabricateALot <=====> AmountToFabricate
	// InvestedCents <===========> InvestedAmount
	// Which means:
	// CostToFabricateALot = (AmountToFabricate * InvestedCents) / InvestedAmount
	return (float64(m.AmountToFabricate) * float64(m.InvestedCents)) / float64(m.InvestedAmount)
}
