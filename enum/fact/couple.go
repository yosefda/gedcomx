// Package enum specifies a set of enumerated values that provides the means for descriptions
// of facts about couple relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package enum

import (
	"github.com/yosefda/gedcomx/typedef"
)

// CoupleFact enum type.
type CoupleFact typedef.URI

// CoupleFact is fact that applicable to a couple.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#22-couple-relationship-fact-types
const (
	Annulment           CoupleFact = "http://gedcomx.org/Annulment"
	CommonLawMarriage   CoupleFact = "http://gedcomx.org/CommonLawMarriage"
	CivilUnion          CoupleFact = "http://gedcomx.org/CivilUnion"
	Divorce             CoupleFact = "http://gedcomx.org/Divorce"
	DivorceFiling       CoupleFact = "http://gedcomx.org/DivorceFiling"
	DomesticPartnership CoupleFact = "http://gedcomx.org/DomesticPartnership"
	Engagement          CoupleFact = "http://gedcomx.org/Engagement"
	Marriage            CoupleFact = "http://gedcomx.org/Marriage"
	MarriageBanns       CoupleFact = "http://gedcomx.org/MarriageBanns"
	MarriageContract    CoupleFact = "http://gedcomx.org/MarriageContract"
	MarriageLicense     CoupleFact = "http://gedcomx.org/MarriageLicense"
	MarriageNotice      CoupleFact = "http://gedcomx.org/MarriageNotice"
	NumberOfChildren    CoupleFact = "http://gedcomx.org/NumberOfChildren"
	Separation          CoupleFact = "http://gedcomx.org/Separation"
)
