package enum

// Unit represents an evaluation of quantity of something. It's followed by another data, a string, which is responsible to describe its meaning.
type Unit float32

const (
	// mass
	MILLIGRAM Unit = 1
	GRAM      Unit = 1000
	OUNCE     Unit = 28349.5
	POUND     Unit = 453592
	KILOGRAM  Unit = 1000000
	TONNE     Unit = 1000000000

	// volume
	MILLILITER            Unit = 1
	AMERICAN_TEASPOON     Unit = 4.92892
	AMERICAN_TABLESPOON   Unit = 14.7868
	AMERICAN_LIQUID_OUNCE Unit = 29.5735
	GLASS                 Unit = 240
	AMERICAN_PAINT        Unit = 473.176
	AMERICAN_LIQUID_QUART Unit = 946.353
	LITER                 Unit = 1000
	AMERICAN_GALLON       Unit = 3785.41
	CUBIC_METER           Unit = 1000000

	// length
	MILLIMETER Unit = 1
	CENTIMETER Unit = 10
	INCH       Unit = 25.4
	FOOT       Unit = 304.8
	YARD       Unit = 914.4
	METER      Unit = 1000
	KILOMETER  Unit = 1000000
	MILE       Unit = 1609000
)
