// Package date defines the defines a genealogical date.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#316-the-date-data-type
package date

import (
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the Date data type.
const URI = "http://gedcomx.org/v1/Date"

// Date data type.
type Date struct {
	Original string
	Formal   typedef.Date
}
