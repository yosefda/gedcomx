// Package enum specifies a set of enumerated that identify common relationship
// types that are relevant to genealogical research.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/relationship-types-specification.md
package enum

import "github.com/yosefda/gedcomx/typedef"

// Relationship enum type.
type Relationship typedef.URI

// Relationship defines common relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/relationship-types-specification.md#2-relationship-types
const (
	AncestorDescendant Relationship = "http://gedcomx.org/AncestorDescendant"
	Couple             Relationship = "http://gedcomx.org/Couple"
	EnslavedBy         Relationship = "http://gedcomx.org/EnslavedBy"
	Godparent          Relationship = "http://gedcomx.org/Godparent"
	ParentChild        Relationship = "http://gedcomx.org/ParentChild"
)
