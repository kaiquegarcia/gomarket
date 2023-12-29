package entity_test

import (
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func materialTestCase(
	t *testing.T,
	title string,
	amountToFabricate float64,
	investedAmount float64,
	investedCents int,
	expectedFabricationCostCents float64,
) {
	t.Run(title, func(t *testing.T) {
		// Arrange
		m := entity.Material{
			AmountToFabricate: enum.Unit(amountToFabricate),
			InvestedAmount:    enum.Unit(investedAmount),
			InvestedCents:     investedCents,
		}

		// Act
		cost := m.FabricationCostCents()

		// Assert
		assert.Equal(t, expectedFabricationCostCents, cost, "cost should be %d", expectedFabricationCostCents)
	})
}

func Test_Material_FabricationCostCents(t *testing.T) {
	// Formulas
	// CostToFabricateALot = (AmountToFabricate * InvestedCents) / InvestedAmount
	materialTestCase(t, "full consumption to fabrication costs the investedCost", 100, 100, 1000, 1000)
	materialTestCase(t, "half consumption to fabrication costs half the investedCost", 50, 100, 1000, 500)
	materialTestCase(t, "1/4 consumption to fabrication costs 1/4 the investedCost", 25, 100, 1000, 250)
	materialTestCase(t, "1/100 consumption to fabrication costs 1/100 the investedCost", 1, 100, 1000, 10)
	materialTestCase(t, "1/1000 consumption to fabrication costs 1/1000 the investedCost", 0.1, 100, 1000, 1)
	materialTestCase(t, "1/10000 consumption to fabrication should cost 1/10000 the investedCost", 0.01, 100, 1000, 0.1)
	materialTestCase(
		t,
		"complex calculation for fun",
		float64(2.5*enum.AMERICAN_TEASPOON),
		float64(1*enum.LITER),
		889,
		10.954524699999999,
	)
}
