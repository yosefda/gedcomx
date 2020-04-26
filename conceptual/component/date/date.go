// Package date defines the defines a genealogical date.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#316-the-date-data-type
package date

import (
	"github.com/yosefda/gedcomx/typedef"
)

// Idenfitier is the identifier for the Date data type.
const Idenfitier = "http://gedcomx.org/v1/Date"

// Properties of the Date data type.
type Properties struct {
	Original string
	Formal   typedef.Date
}
