// Package attribution defines the data structure used to attribute who, when, and why to genealogical data.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#32-the-attribution-data-type
package attribution

import "github.com/yosefda/gedcomx/conceptual"

// Idenfitier is the identifier for the Attribution data type.
const Idenfitier = "http://gedcomx.org/v1/Attribution"

// Properties of the Identifier data type.
type Properties struct {
	Contributor   conceptual.URI
	Modified      conceptual.Timestamp
	ChangeMessage string
	Creator       conceptual.URI
	Created       conceptual.Timestamp
}
