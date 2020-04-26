// Package qualifier specifies a set of enumerated values to qualify source reference.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/conceptual-model-specification.md#known-source-reference-qualifiers
package qualifier

import (
	"github.com/yosefda/gedcomx/typedef"
)

// SourceReference enum type.
type SourceReference typedef.URI

// SourceReference defines source reference.
const (
	CharacterRegion SourceReference = "http://gedcomx.org/CharacterRegion"
	RectangleRegion SourceReference = "http://gedcomx.org/RectangleRegion"
	TimeRegion      SourceReference = "http://gedcomx.org/TimeRegion"
)
