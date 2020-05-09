// Package attribution defines the data structure to supply an identifier of a genealogical resource.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#32-the-attribution-data-type
package attribution

import "github.com/yosefda/gedcomx/typedef"

// URI for the Attribution data type
const URI = "http://gedcomx.org/v1/Attribution"

// Attribution data type.
type Attribution struct {
	Contributor   typedef.URI
	Modified      *typedef.Timestamp
	ChangeMessage string
	Creator       typedef.URI
	Created       *typedef.Timestamp
}
