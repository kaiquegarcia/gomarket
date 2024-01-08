package util

import (
	"gomarket/internal/enum"
)

func Convert(fromAmount enum.Unit, fromID enum.UnitID, toID enum.UnitID) enum.Unit {
	if fromID == toID || fromID == enum.UnitID(0) {
		return fromAmount
	}

	fromAmount *= enum.GetBaseUnit(fromID)
	return fromAmount / enum.GetBaseUnit(toID)
}
