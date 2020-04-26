// Package name specifies a set of enumerated values for name types.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#3131-known-name-types
package name

import "github.com/yosefda/gedcomx/typedef"

// Type enum type.
type Type typedef.URI

// NameType defines name type.
const (
	BirthName     Type = "http://gedcomx.org/BirthName"
	MarriedName   Type = "http://gedcomx.org/MarriedName"
	AlsoKnownAs   Type = "http://gedcomx.org/AlsoKnownAs"
	Nickname      Type = "http://gedcomx.org/Nickname"
	AdoptiveName  Type = "http://gedcomx.org/AdoptiveName"
	FormalName    Type = "http://gedcomx.org/FormalName"
	ReligiousName Type = "http://gedcomx.org/ReligiousName"
)
