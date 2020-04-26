// Package nameform defines a representation of a name (a "name form") within a given
// cultural context, such as a given language and script.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#319-the-nameform-data-type
package nameform

import (
	"github.com/yosefda/gedcomx/conceptual/component/namepart"
	"github.com/yosefda/gedcomx/typedef"
)

// Idenfitier is the identifier for the NameForm data type.
const Idenfitier = "http://gedcomx.org/v1/NameForm"

// Properties of the NameForm data type.
type Properties struct {
	Lang     typedef.LocaleTag
	FullText string
	Parts    *[]namepart.Properties
}
