// Package qualifier defines the data structure used to supply additional details, annotations, tags,
// or other qualifying data to a specific data element.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#qualifier
package qualifier

import "github.com/yosefda/gedcomx/typedef"

// Idenfitier is the identifier for the Qualifier data type.
const Idenfitier = "http://gedcomx.org/v1/Qualifier"

// Properties of the Qualifier data type.
type Properties struct {
	Name  typedef.URI
	Value string
}
