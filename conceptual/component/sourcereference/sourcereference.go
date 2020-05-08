// Package sourcereference defines the data structure used to defines a reference to a source description.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#36-the-sourcereference-data-type
package sourcereference

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	enum "github.com/yosefda/gedcomx/enum/qualifier"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the SourceReference data type.
const URI = "http://gedcomx.org/v1/SourceReference"

// SourceReference data type.
type SourceReference struct {
	Description   typedef.URI
	DescriptionID string
	Attribution   *attribution.Attribution
	Qualifiers    enum.SourceReference
}
