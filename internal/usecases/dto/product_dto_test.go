package dto_test

import (
	"gomarket/internal/entity"
	"gomarket/internal/enum"
	"gomarket/internal/usecases/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ProductDTO_ToEntity(t *testing.T) {
	t.Run("should encode entity with updatedAt nil", func(t *testing.T) {
		// Arrange
		dto := dto.ProductDTO{
			Name: "Cupcake",
			Materials: []entity.Material{
				{
					ProductCode:       123,
					UnitKind:          enum.MASS,
					AmountToFabricate: 120 * enum.GRAM,
					InvestedCents:     100,
				},
			},
			SellingPriceCents: 120,
		}
		createdAt, _ := time.Parse(time.DateOnly, "2023-12-28")

		// Act
		entity := dto.ToEntity(
			321,
			createdAt,
			nil,
		)

		// Assert
		assert.Equal(t, 321, entity.Code, "code should be 321")
		assert.Equal(t, "Cupcake", entity.Name, "name should be 'Cupcake'")
		assert.Equal(t, dto.Materials, entity.Materials, "materials should be equal to the present on DTO")
		assert.Equal(t, 120, entity.SellingPriceCents, "selling_price_cents should be equal to 120")
		assert.Equal(t, createdAt, entity.CreatedAt, "created_at should be equal to the parsed for the test")
		assert.Nil(t, entity.UpdatedAt, "updated_at should be nil")
	})

	t.Run("should encode entity with updatedAt filled", func(t *testing.T) {
		// Arrange
		dto := dto.ProductDTO{
			Name: "Cupcake",
			Materials: []entity.Material{
				{
					ProductCode:       123,
					UnitKind:          enum.MASS,
					AmountToFabricate: 120 * enum.GRAM,
					InvestedCents:     100,
				},
			},
			SellingPriceCents: 120,
		}
		createdAt, _ := time.Parse(time.DateOnly, "2023-12-28")
		updatedAt, _ := time.Parse(time.DateOnly, "2023-12-28")

		// Act
		entity := dto.ToEntity(
			321,
			createdAt,
			&updatedAt,
		)

		// Assert
		assert.Equal(t, 321, entity.Code, "code should be 321")
		assert.Equal(t, "Cupcake", entity.Name, "name should be 'Cupcake'")
		assert.Equal(t, dto.Materials, entity.Materials, "materials should be equal to the present on DTO")
		assert.Equal(t, 120, entity.SellingPriceCents, "selling_price_cents should be equal to 120")
		assert.Equal(t, createdAt, entity.CreatedAt, "created_at should be equal to the parsed for the test")
		assert.Equal(t, updatedAt, *entity.UpdatedAt, "updated_at should be equal to the parsed for the test")
	})
}
