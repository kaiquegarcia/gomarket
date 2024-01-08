package entity

import (
	"gomarket/internal/enum"
	"gomarket/pkg/util"
)

// Material represents a product required to produce other product. With this struct is possible to calculate the total cost to fabricate a single unit of the product.
type Material struct {
	ProductCode       int           `json:"product_code"`
	Unit              enum.UnitKind `json:"unit"`
	FabricationUnitID enum.UnitID   `json:"fabrication_unit_id"`
	AmountToFabricate enum.Unit     `json:"amount_to_fabricate"`
	InvestUnitID      enum.UnitID   `json:"invest_unit_id"`
	InvestedAmount    enum.Unit     `json:"invested_amount"`
	InvestedCents     int           `json:"invested_cents"`
}

// FabricationCostCents will calculate the cost to use this Material to fabricate a single product, based on the numbers defined in {AmountToFabricate}, {InvestedAmount}, {InvestedCents} and {FabricatedProductCount}
func (m Material) FabricationCostCents() float64 {
	if m.InvestedCents == 0 || m.AmountToFabricate == 0 || m.InvestedAmount == 0 {
		return 0
	}

	baseUnitID := enum.DefaultUnitID(m.Unit)
	baseAmountToFabricate := util.Convert(m.AmountToFabricate, m.FabricationUnitID, baseUnitID)
	baseInvestedAmount := util.Convert(m.InvestedAmount, m.InvestUnitID, baseUnitID)
	// Math:
	// CostToFabricateALot <=====> AmountToFabricate
	// InvestedCents <===========> InvestedAmount
	// Which means:
	// CostToFabricateALot = (AmountToFabricate * InvestedCents) / InvestedAmount
	return (float64(baseAmountToFabricate) * float64(m.InvestedCents)) / float64(baseInvestedAmount)
}
