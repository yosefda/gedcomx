// Package idenfitier specifies a set of enumerated that identifies how the identifier
// is to be used and the nature of the resource to which the identifier resolves
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#known-identifier-types
package idenfitier

import "github.com/yosefda/gedcomx/typedef"

// Identifier enum type.
type Identifier typedef.URI

// Identifier defines how the identifier is to be used and the nature of the resource.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#known-identifier-types
const (
	Primary    Identifier = "http://gedcomx.org/Primary"
	Authority  Identifier = "http://gedcomx.org/Authority"
	Deprecated Identifier = "http://gedcomx.org/Deprecated"
)
