// Package gender defines data structure to represent a gender of a person.package gender
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#312-the-gender-data-type
package gender

import enum "github.com/yosefda/gedcomx/enum/gender"

// URI is the identifier for the Gender data type.
const URI = "http://gedcomx.org/v1/Gender"

// Gender data type.
type Gender struct {
	Type enum.Gender
}
