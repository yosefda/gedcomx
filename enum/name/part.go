// Package name specifies a set of enumerated values for name part types.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#3181-known-name-part-types
package name

import "github.com/yosefda/gedcomx/typedef"

// Part enum type.
type Part typedef.URI

// Part defines name part type.
const (
	Prefix  Part = "http://gedcomx.org/Prefix"
	Suffix  Part = "http://gedcomx.org/Suffix"
	Given   Part = "http://gedcomx.org/Given"
	Surname Part = "http://gedcomx.org/Surname"
)
