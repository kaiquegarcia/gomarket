package entity_test

import (
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Product_FabricationCostCents(t *testing.T) {
	// Arrange
	// CostToFabricateALot = (AmountToFabricate * InvestedCents) / InvestedAmount
	recipt := []entity.Material{
		{ // softened butter
			// CostToFabricateALot = ((3 * 14.7868) * 1190) / 200000 = 0.26394438
			AmountToFabricate: 3 * enum.AMERICAN_TABLESPOON,
			InvestedAmount:    200 * enum.GRAM,
			InvestedCents:     1190,
		},
		{ // cocoa powder
			// CostToFabricateALot = ((2 * 14.7868) * 1500) / 200000 = 0.221802
			AmountToFabricate: 2 * enum.AMERICAN_TABLESPOON,
			InvestedAmount:    200 * enum.GRAM,
			InvestedCents:     1500,
		},
		{ // suggar
			// CostToFabricateALot = (120000 * 459) / 1000000 = 55.08
			AmountToFabricate: 120 * enum.GRAM,
			InvestedAmount:    1 * enum.KILOGRAM,
			InvestedCents:     459,
		},
		{ // wheat flour
			// CostToFabricateALot = (120000 * 549) / 1000000 = 65.88
			AmountToFabricate: 120 * enum.GRAM,
			InvestedAmount:    1 * enum.KILOGRAM,
			InvestedCents:     549,
		},
		{ // baking powder
			// CostToFabricateALot = (2000 * 579) / 100000 = 11.58
			AmountToFabricate: 2 * enum.GRAM,
			InvestedAmount:    100 * enum.GRAM,
			InvestedCents:     579,
		},
		{ // salt
			// CostToFabricateALot = (250 * 189) / 1000000 = 0.04725
			AmountToFabricate: 0.25 * enum.GRAM,
			InvestedAmount:    1 * enum.KILOGRAM,
			InvestedCents:     189,
		},
		{ // egg
			// CostToFabricateALot = (1 * 700) / 12 = 58.33333333333333
			AmountToFabricate: 1,
			InvestedAmount:    12,
			InvestedCents:     700,
		},
		{ // milk
			// CostToFabricateALot = (120 * 700) / 1000 = 84
			AmountToFabricate: 120 * enum.MILLILITER,
			InvestedAmount:    1 * enum.LITER,
			InvestedCents:     700,
		},
	}
	p := entity.Product{
		Materials:      recipt,
		QuantityPerLot: 8,
	}
	// TotalCostToFabricateALot = 0.26394438 + 0.221802 + 55.08 + 65.88 + 11.58 + 0.04725 + 58.333333333333336 + 84
	// TotalCostToFabricateALot = ~275.4063277133333
	// TotalCostToFabricateOneProduct = ~275.4063277133333 / 8 = 34.42579096416667 =~ 35

	// Act
	cost := p.FabricationCostCents()

	// Assert
	assert.Equal(t, 0.26394438, recipt[0].FabricationCostCents())
	assert.Equal(t, 0.221802, recipt[1].FabricationCostCents())
	assert.Equal(t, 55.08, recipt[2].FabricationCostCents())
	assert.Equal(t, 65.88, recipt[3].FabricationCostCents())
	assert.Equal(t, 11.58, recipt[4].FabricationCostCents())
	assert.Equal(t, 0.04725, recipt[5].FabricationCostCents())
	assert.Equal(t, 58.333333333333336, recipt[6].FabricationCostCents())
	assert.Equal(t, 84.0, recipt[7].FabricationCostCents())
	assert.Equal(t, 35, cost, "cost should be 35")
}
