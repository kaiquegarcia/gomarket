package enum

// UnitKind represents the kind of an unit. It's followed by another data, a float32 number, describing it means.
type UnitKind string

const (
	MASS   UnitKind = "mg"
	VOLUME UnitKind = "ml"
	LENGTH UnitKind = "mm"
	UNIT   UnitKind = "u"
)

var UnitKinds = []string{string(MASS), string(VOLUME), string(LENGTH), string(UNIT)}

func DefaultUnitID(kind UnitKind) UnitID {
	switch kind {
	case MASS:
		return MILLIGRAM_ID
	case VOLUME:
		return MILLILITER_ID
	case LENGTH:
		return MILLIMETER_ID
	}

	return UNIT_ID
}
