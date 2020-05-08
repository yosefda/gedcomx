// Package identifier defines the data structure to supply an identifier of a genealogical resource.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#identifier-type
package identifier

import (
	enum "github.com/yosefda/gedcomx/enum/identifier"
	"github.com/yosefda/gedcomx/typedef"
)

// URI for the Identifier data type
const URI = "http://gedcomx.org/v1/Identifier"

// Identifier data type.
type Identifier struct {
	Value typedef.URI
	Type  enum.Identifier
}
