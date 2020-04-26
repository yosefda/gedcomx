// Package confidence specifies a set of enumerated values that indicatie confidence level for an assertion.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#3101-known-confidence-levels
package confidence

import (
	"github.com/yosefda/gedcomx/typedef"
)

// Confidence enum type.
type Confidence typedef.URI

// Confidence defines confidence level.
const (
	High   Confidence = "http://gedcomx.org/High"
	Medium Confidence = "http://gedcomx.org/Medium"
	Low    Confidence = "http://gedcomx.org/Low"
)
