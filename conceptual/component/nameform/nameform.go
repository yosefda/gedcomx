// Package nameform defines a representation of a name (a "name form") within a given
// cultural context, such as a given language and script.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#319-the-nameform-data-type
package nameform

import (
	"github.com/yosefda/gedcomx/conceptual/component/namepart"
	"github.com/yosefda/gedcomx/typedef"
)

// URI is the identifier for the NameForm data type.
const URI = "http://gedcomx.org/v1/NameForm"

// NameForm data type.
type NameForm struct {
	Lang     typedef.LocaleTag
	FullText string
	Parts    *[]namepart.NamePart
}
