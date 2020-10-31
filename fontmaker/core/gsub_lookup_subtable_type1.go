package core

//GSUBLookupSubTableType1Format1 Single Substitution Format 1
type GSUBLookupSubTableType1Format1 struct {
	coverageOffset int64
	deltaGlyphID   int
}

//LookupType get lookup type
func (g GSUBLookupSubTableType1Format1) LookupType() uint {
	return 1
}

//Format get format
func (g GSUBLookupSubTableType1Format1) Format() uint {
	return 1
}

//GSUBLookupSubTableType1Format2 Single Substitution Format 2
type GSUBLookupSubTableType1Format2 struct {
	coverageOffset     int64
	substituteGlyphIDs []uint
}

//LookupType get lookup type
func (g GSUBLookupSubTableType1Format2) LookupType() uint {
	return 1
}

//Format get format
func (g GSUBLookupSubTableType1Format2) Format() uint {
	return 2
}