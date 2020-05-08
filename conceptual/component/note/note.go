// Package note defines the data structure used to represent a note that was contributed.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#33-the-note-data-type
package note

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the Node data type.
const URI = "http://gedcomx.org/v1/Note"

// Note data type.
type Note struct {
	Lang        typedef.LocaleTag
	Subject     string
	Text        string
	Attribution *attribution.Attribution
}
