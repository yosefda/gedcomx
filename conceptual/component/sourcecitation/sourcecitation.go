// Package sourcecitation defines the data structure for a container for the metadata necessary
// for an agent to identify a source(s).
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#35-the-sourcecitation-data-type
package sourcecitation

import "github.com/yosefda/gedcomx/typedef"

// Idenfitier is the identifier for the SourceCitation data type.
const Idenfitier = "http://gedcomx.org/v1/SourceCitation"

// Properties of the SourceCitation data type.
type Properties struct {
	Lang  typedef.LocaleTag
	Value string
}
