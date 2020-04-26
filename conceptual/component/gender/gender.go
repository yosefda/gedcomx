// Package gender defines data structure to represent a gender of a person.package gender
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#312-the-gender-data-type
package gender

import enum "github.com/yosefda/gedcomx/enum/gender"

// Idenfitier is the identifier for the Gender data type.
const Idenfitier = "http://gedcomx.org/v1/Gender"

// Male gender.
const Male = "http://gedcomx.org/Male"

// Female gender.
const Female = "http://gedcomx.org/Female"

// Unknown gender.
const Unknown = "http://gedcomx.org/Unknown"

// Intersex (assignment at birth).
const Intersex = "http://gedcomx.org/Intersex"

// Properties of the Gender data type.
type Properties struct {
	Type enum.Gender
}
