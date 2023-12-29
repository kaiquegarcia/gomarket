package enum

// Unit represents an evaluation of quantity of something. It's followed by another data, a string, which is responsible to describe its meaning.
type Unit float64

const (
	// mass
	MILLIGRAM Unit = 1
	GRAM           = 1000 * MILLIGRAM
	OUNCE          = 28349.5 * MILLIGRAM
	POUND          = 453592 * MILLIGRAM
	KILOGRAM       = 1e+6 * MILLIGRAM
	TONNE          = 1e+9 * MILLIGRAM

	// volume
	MILLILITER            Unit = 1
	AMERICAN_TEASPOON          = 4.92892 * MILLILITER
	AMERICAN_TABLESPOON        = 14.7868 * MILLILITER
	AMERICAN_LIQUID_OUNCE      = 29.5735 * MILLILITER
	GLASS                      = 240 * MILLILITER
	AMERICAN_PAINT             = 473.176 * MILLILITER
	AMERICAN_LIQUID_QUART      = 946.353 * MILLILITER
	LITER                      = 1000 * MILLILITER
	AMERICAN_GALLON            = 3785.41 * MILLILITER
	CUBIC_METER                = 1e+6 * MILLILITER

	// length
	MILLIMETER Unit = 1
	CENTIMETER      = 10 * MILLIMETER
	INCH            = 25.4 * MILLIMETER
	FOOT            = 304.8 * MILLIMETER
	YARD            = 914.4 * MILLIMETER
	METER           = 1000 * MILLIMETER
	KILOMETER       = 1e+6 * MILLIMETER
	MILE            = 1.609e+6 * MILLIMETER
)
