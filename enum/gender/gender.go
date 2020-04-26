// Package gender specifies a set of enumerated values that indicatie gender types.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#3121-known-gender-types
package gender

import (
	"github.com/yosefda/gedcomx/typedef"
)

// Gender enum type.
type Gender typedef.URI

// Gender defines gender type.
const (
	Male     Gender = "http://gedcomx.org/Male"
	Female   Gender = "http://gedcomx.org/Female"
	Unknown  Gender = "http://gedcomx.org/Unknown"
	Intersex Gender = "http://gedcomx.org/Intersex"
)
