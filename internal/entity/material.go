package entity

import "gomarket/internal/enum"

// Material represents a product required to produce other product. With this struct is possible to calculate the total cost to fabricate a single unit of the product.
type Material struct {
	ProductCode   int           `json:"product_code"`
	Unit          enum.UnitKind `json:"unit"`
	Amount        enum.Unit     `json:"amount"`
	InvestedCents int           `json:"invested_cents"`
}
