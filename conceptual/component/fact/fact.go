// Package fact defines a data item that is presumed to be true about a specific subject, such as a person or relationship
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#314-the-fact-data-type
package fact

import (
	"github.com/yosefda/gedcomx/conceptual/component/placereference"
	"github.com/yosefda/gedcomx/conceptual/component/qualifier"
	"github.com/yosefda/gedcomx/enum/fact"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the Fact data type.
const URI = "http://gedcomx.org/v1/Fact"

// Fact data type.
type Fact struct {
	Type       *fact.Fact
	Date       *typedef.Date
	Place      *placereference.PlaceReference
	Value      string
	Qualifiers *[]qualifier.Qualifier
}
