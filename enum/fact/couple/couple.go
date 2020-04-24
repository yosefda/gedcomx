// Package couple specifies a set of enumerated values that provides the means for descriptions
// of facts about couple relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package couple

import (
	"github.com/yosefda/gedcomx/typedef"
)

// Fact enum type.
type Fact typedef.URI

// Fact is fact that applicable to a couple.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#22-couple-relationship-fact-types
const (
	Annulment           Fact = "http://gedcomx.org/Annulment"
	CommonLawMarriage   Fact = "http://gedcomx.org/CommonLawMarriage"
	CivilUnion          Fact = "http://gedcomx.org/CivilUnion"
	Divorce             Fact = "http://gedcomx.org/Divorce"
	DivorceFiling       Fact = "http://gedcomx.org/DivorceFiling"
	DomesticPartnership Fact = "http://gedcomx.org/DomesticPartnership"
	Engagement          Fact = "http://gedcomx.org/Engagement"
	Marriage            Fact = "http://gedcomx.org/Marriage"
	MarriageBanns       Fact = "http://gedcomx.org/MarriageBanns"
	MarriageContract    Fact = "http://gedcomx.org/MarriageContract"
	MarriageLicense     Fact = "http://gedcomx.org/MarriageLicense"
	MarriageNotice      Fact = "http://gedcomx.org/MarriageNotice"
	NumberOfChildren    Fact = "http://gedcomx.org/NumberOfChildren"
	Separation          Fact = "http://gedcomx.org/Separation"
)
