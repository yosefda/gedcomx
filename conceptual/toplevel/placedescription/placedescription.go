// Package placedescription describes the details of a place, functioning as a description of a place as a snapshot in time
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#27-the-placedescription-data-type
package placedescription

import (
	"github.com/yosefda/gedcomx/conceptual/component/date"
	"github.com/yosefda/gedcomx/conceptual/component/textvalue"
	enum "github.com/yosefda/gedcomx/enum/place"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the PlaceDescription data type.
const URI = "http://gedcomx.org/v1/PlaceDescription"

// PlaceDescription data type.
type PlaceDescription struct {
	Names               *[]textvalue.TextValue
	Type                enum.Type
	Place               typedef.URI
	Jurisdiction        *PlaceDescription
	Latitude            float32
	Longitude           float32
	TemporalDescription *date.Date
	SpatialDescription  typedef.URI
}
