// Package qualifier specifies a set of enumerated values that provides the means for parts
// of names of persons to be described and qualified
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/name-part-qualifiers-specification.md
package qualifier

import "github.com/yosefda/gedcomx/typedef"

// NamePart enum type.
type NamePart typedef.URI

// NamePart defines part of name of a person.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/name-part-qualifiers-specification.md#2-name-part-qualifiers
const (
	Title          NamePart = "http://gedcomx.org/Title"
	Primary        NamePart = "http://gedcomx.org/Primary"
	Secondary      NamePart = "http://gedcomx.org/Secondary"
	Middle         NamePart = "http://gedcomx.org/Middle"
	Familiar       NamePart = "http://gedcomx.org/Familiar"
	Religious      NamePart = "http://gedcomx.org/Religious"
	Family         NamePart = "http://gedcomx.org/Family"
	Maiden         NamePart = "http://gedcomx.org/Maiden"
	Patronymic     NamePart = "http://gedcomx.org/Patronymic"
	Matronymic     NamePart = "http://gedcomx.org/Matronymic"
	Geographic     NamePart = "http://gedcomx.org/Geographic"
	Occupational   NamePart = "http://gedcomx.org/Occupational"
	Characteristic NamePart = "http://gedcomx.org/Characteristic"
	Postnom        NamePart = "http://gedcomx.org/Postnom"
	Particle       NamePart = "http://gedcomx.org/Particle"
	RootName       NamePart = "http://gedcomx.org/RootName"
)
