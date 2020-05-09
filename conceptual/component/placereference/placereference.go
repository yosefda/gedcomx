// Package placereference  defines a reference to a description of a place
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#317-the-placereference-data-type
package placereference

import "github.com/yosefda/gedcomx/conceptual/toplevel/placedescription"

// URI is the identifier for the PlaceReference data type.
const URI = "http://gedcomx.org/v1/PlaceReference"

// PlaceReference data type.
type PlaceReference struct {
	Original       string
	DescriptionRef *placedescription.PlaceDescription
}
