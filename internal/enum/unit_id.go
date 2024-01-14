package enum

type UnitID string

const (
	MILLIGRAM_ID             UnitID = "milligram"
	GRAM_ID                  UnitID = "gram"
	OUNCE_ID                 UnitID = "ounce"
	POUND_ID                 UnitID = "pound"
	KILOGRAM_ID              UnitID = "kilogram"
	TONNE_ID                 UnitID = "tonne"
	MILLILITER_ID            UnitID = "milliliter"
	AMERICAN_TEASPOON_ID     UnitID = "american_teaspoon"
	AMERICAN_TABLESPOON_ID   UnitID = "american_tablespoon"
	AMERICAN_LIQUID_OUNCE_ID UnitID = "american_liquid_ounce"
	GLASS_ID                 UnitID = "glass"
	AMERICAN_PAINT_ID        UnitID = "american_paint"
	AMERICAN_LIQUID_QUART_ID UnitID = "american_liquid_quart"
	LITER_ID                 UnitID = "liter"
	AMERICAN_GALLON_ID       UnitID = "american_gallon"
	CUBIC_METER_ID           UnitID = "cubic_meter"
	MILLIMETER_ID            UnitID = "millimeter"
	CENTIMETER_ID            UnitID = "centimeter"
	INCH_ID                  UnitID = "inch"
	FOOT_ID                  UnitID = "foot"
	YARD_ID                  UnitID = "yard"
	METER_ID                 UnitID = "meter"
	KILOMETER_ID             UnitID = "kilometer"
	MILE_ID                  UnitID = "mile"
	UNIT_ID                  UnitID = "un"
)

var UnitIDs = []string{
	string(MILLIGRAM_ID),
	string(GRAM_ID),
	string(OUNCE_ID),
	string(POUND_ID),
	string(KILOGRAM_ID),
	string(TONNE_ID),
	string(MILLILITER_ID),
	string(AMERICAN_TEASPOON_ID),
	string(AMERICAN_TABLESPOON_ID),
	string(AMERICAN_LIQUID_OUNCE_ID),
	string(GLASS_ID),
	string(AMERICAN_PAINT_ID),
	string(AMERICAN_LIQUID_QUART_ID),
	string(LITER_ID),
	string(AMERICAN_GALLON_ID),
	string(CUBIC_METER_ID),
	string(MILLIMETER_ID),
	string(CENTIMETER_ID),
	string(INCH_ID),
	string(FOOT_ID),
	string(YARD_ID),
	string(METER_ID),
	string(KILOMETER_ID),
	string(MILE_ID),
	string(UNIT_ID),
}

// GetBaseUnit returns the amount of the specified id
func GetBaseUnit(id UnitID) Unit {
	switch id {
	case MILLIGRAM_ID:
		return MILLIGRAM
	case GRAM_ID:
		return GRAM
	case OUNCE_ID:
		return OUNCE
	case POUND_ID:
		return POUND
	case KILOGRAM_ID:
		return KILOGRAM
	case TONNE_ID:
		return TONNE
	case MILLILITER_ID:
		return MILLILITER
	case AMERICAN_TEASPOON_ID:
		return AMERICAN_TEASPOON
	case AMERICAN_TABLESPOON_ID:
		return AMERICAN_TABLESPOON
	case AMERICAN_LIQUID_OUNCE_ID:
		return AMERICAN_LIQUID_OUNCE
	case GLASS_ID:
		return GLASS
	case AMERICAN_PAINT_ID:
		return AMERICAN_PAINT
	case AMERICAN_LIQUID_QUART_ID:
		return AMERICAN_LIQUID_QUART
	case LITER_ID:
		return LITER
	case AMERICAN_GALLON_ID:
		return AMERICAN_GALLON
	case CUBIC_METER_ID:
		return CUBIC_METER
	case MILLIMETER_ID:
		return MILLIMETER
	case CENTIMETER_ID:
		return CENTIMETER
	case INCH_ID:
		return INCH
	case FOOT_ID:
		return FOOT
	case YARD_ID:
		return YARD
	case METER_ID:
		return METER
	case KILOMETER_ID:
		return KILOMETER
	case MILE_ID:
		return MILE
	case UNIT_ID:
		return 1
	}

	return 1
}
