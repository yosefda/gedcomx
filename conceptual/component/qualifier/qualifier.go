// Package qualifier defines the data structure used to supply additional details, annotations, tags,
// or other qualifying data to a specific data element.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#qualifier
package qualifier

import "github.com/yosefda/gedcomx/typedef"

// URI is the identifier for the Qualifier data type.
const URI = "http://gedcomx.org/v1/Qualifier"

// Qualifier data type.
type Qualifier struct {
	Name  typedef.URI
	Value string
}
