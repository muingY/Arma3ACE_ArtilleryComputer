package artilleryCalculator

type RangeTableCell struct {
	Distance float64
	AimMill float64
	DeltaMillPer100m float64
	Eta float64
}

type M109A6RangeTable struct {
	RangeTableLow [5][]RangeTableCell
	RangeTableHigh [5][]RangeTableCell
}

type ShotResultData struct {
	Charge int
	AimMill float64
	Eta float64
}