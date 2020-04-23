// Package identifier defines the data structure used to supply an identifier of a genealogical resource.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#identifier-type
package identifier

import (
	"github.com/yosefda/gedcomx/conceptual"
)

// Idenfitier is the identifier for the Identifier data type.
const Idenfitier = "http://gedcomx.org/v1/Identifier"

// Primary is the primary identifier for the resource.
const Primary = "http://gedcomx.org/Primary"

// Authority is used for the resource in an external authority or other expert system.
const Authority = "http://gedcomx.org/Authority"

// Deprecated is used when the resource that has been relegated, deprecated, or otherwise downgraded.
const Deprecated = "http://gedcomx.org/Deprecated"

// Properties of the Identifier data type.
type Properties struct {
	Value conceptual.URI
	Type  conceptual.URI
}
