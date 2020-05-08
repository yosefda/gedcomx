// Package namepart defines the  data type is used to model a portion of a full name,
// including the terms that make up that portion.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#318-the-namepart-data-type
package namepart

import (
	enumName "github.com/yosefda/gedcomx/enum/name"
	enumQualifier "github.com/yosefda/gedcomx/enum/qualifier"
)

// URI is the identifier for the NamePart data type.
const URI = "http://gedcomx.org/v1/NamePart"

// NamePart data type.
type NamePart struct {
	Type      enumName.Part
	Value     string
	Qualifier *[]enumQualifier.NamePart
}
