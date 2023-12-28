package entity

import (
	"time"
)

// Product represents any item, for selling or not. If you use it to fabricate another product, it's a product as well.
type Product struct {
	Code              int        `json:"code"`
	Name              string     `json:"name"`
	Materials         []Material `json:"materials"`
	SellingPriceCents int        `json:"selling_price_cents"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}
