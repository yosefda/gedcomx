// Package evidencereference defines the data structure used as a reference
// to data being used to derive the given instance of Subject.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#37-the-evidencereference-data-type
package evidencereference

import (
	"github.com/yosefda/gedcomx/conceptual/component/attribution"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the EvidenceReference data type.
const URI = "http://gedcomx.org/v1/EvidenceReference"

// EvidenceReference data type.
type EvidenceReference struct {
	Resource    typedef.URI
	Attribution *attribution.Attribution
}
