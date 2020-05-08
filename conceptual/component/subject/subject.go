// Package subject defines data structure to represent the abstract concept of a genealogical subject.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#311-the-subject-data-type
package subject

import (
	"github.com/yosefda/gedcomx/conceptual/component/evidencereference"
	"github.com/yosefda/gedcomx/conceptual/component/identifier"
	"github.com/yosefda/gedcomx/conceptual/component/sourcereference"
)

// URI is the identifier for the Subject data type.
const URI = "http://gedcomx.org/v1/Subject"

// Properties of the Subect data type.
type Properties struct {
	Extracted   bool
	Evidence    *[]evidencereference.EvidenceReference
	media       *[]sourcereference.SourceReference
	identifiers *[]identifier.Identifier
}
