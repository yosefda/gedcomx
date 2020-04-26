// Package note defines the data structure used to represent a note that was contributed.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#33-the-note-data-type
package note

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	"github.com/yosefda/gedcomx/typedef"
)

// Identifier is the identifier for the Node data type.
const Identifier = "http://gedcomx.org/v1/Note"

// Properties of the Identifier data type.
type Properties struct {
	Lang        typedef.LocaleTag
	Subject     string
	Text        string
	Attribution *attribution.Properties
}
