// Package name defines data structure to represent a name of a person.package name
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#313-the-name-data-type
package name

import (
	"github.com/yosefda/gedcomx/conceptual/component/date"
	"github.com/yosefda/gedcomx/conceptual/component/nameform"
	enumName "github.com/yosefda/gedcomx/enum/name"
)

// Idenfitier is the identifier for the Name data type.
const Idenfitier = "http://gedcomx.org/v1/Identifier"

// Properties of the Name data type.
type Properties struct {
	Type      enumName.Type
	NameForms *[]nameform.Properties
	Date      *date.Properties
}
