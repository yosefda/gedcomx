// Package enum specifies a set of enumerated values that provides the means for parts
// of names of persons to be described and qualified
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/name-part-qualifiers-specification.md
package enum

import "github.com/yosefda/gedcomx/typedef"

// NamePartQualifier enum type.
type NamePartQualifier typedef.URI

// NamePartQualifier defines part of name of a person.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/name-part-qualifiers-specification.md#2-name-part-qualifiers
const (
	Title          NamePartQualifier = "http://gedcomx.org/Title"
	Primary        NamePartQualifier = "http://gedcomx.org/Primary"
	Secondary      NamePartQualifier = "http://gedcomx.org/Secondary"
	Middle         NamePartQualifier = "http://gedcomx.org/Middle"
	Familiar       NamePartQualifier = "http://gedcomx.org/Familiar"
	Religious      NamePartQualifier = "http://gedcomx.org/Religious"
	Family         NamePartQualifier = "http://gedcomx.org/Family"
	Maiden         NamePartQualifier = "http://gedcomx.org/Maiden"
	Patronymic     NamePartQualifier = "http://gedcomx.org/Patronymic"
	Matronymic     NamePartQualifier = "http://gedcomx.org/Matronymic"
	Geographic     NamePartQualifier = "http://gedcomx.org/Geographic"
	Occupational   NamePartQualifier = "http://gedcomx.org/Occupational"
	Characteristic NamePartQualifier = "http://gedcomx.org/Characteristic"
	Postnom        NamePartQualifier = "http://gedcomx.org/Postnom"
	Particle       NamePartQualifier = "http://gedcomx.org/Particle"
	RootName       NamePartQualifier = "http://gedcomx.org/RootName"
)
