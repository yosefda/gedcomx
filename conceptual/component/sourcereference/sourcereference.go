// Package sourcereference defines the data structure used to defines a reference to a source description.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#36-the-sourcereference-data-type
package sourcereference

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	enum "github.com/yosefda/gedcomx/enum/qualifier"
	"github.com/yosefda/gedcomx/typedef"
)

// Idenfitier is the identifier for the SourceReference data type.
const Idenfitier = "http://gedcomx.org/v1/SourceReference"

// Properties of the SourceReference data type.
type Properties struct {
	Description   typedef.URI
	DescriptionID string
	Attribution   *attribution.Properties
	Qualifiers    enum.SourceReference
}
