// Package textvalue defines data type for a literal text value.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#34-the-textvalue-data-type
package textvalue

import "github.com/yosefda/gedcomx/typedef"

// Idenfitier is the identifier for the TextValue data type.
const Idenfitier = "http://gedcomx.org/v1/TextValue"

// Properties of the TextValue data type.
type Properties struct {
	Lang  typedef.LocaleTag
	Value string
}
