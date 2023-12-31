package format

import (
	"fmt"
	"math"
)

func MoneyFromCents(cents float64) string {
	return fmt.Sprintf("$%.2f", math.Ceil(cents)/100)
}
