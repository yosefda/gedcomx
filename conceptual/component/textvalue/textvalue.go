// Package textvalue defines data type for a literal text value.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#34-the-textvalue-data-type
package textvalue

import "github.com/yosefda/gedcomx/typedef"

// URI is the identifier for the TextValue data type.
const URI = "http://gedcomx.org/v1/TextValue"

// TextValue data type.
type TextValue struct {
	Lang  *typedef.LocaleTag
	Value string
}
