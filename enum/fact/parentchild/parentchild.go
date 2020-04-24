// Package parentchild specifies a set of enumerated values that provides the means for descriptions
// of facts about parent-child relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package parentchild

import "github.com/yosefda/gedcomx/typedef"

// Fact enum type.
type Fact typedef.URI

// Fact is fact that applicable to  a parent-child relationship.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#23-parent-child-relationship-fact-types
const (
	AdoptiveParent     Fact = "http://gedcomx.org/AdoptiveParent"
	BiologicalParent   Fact = "http://gedcomx.org/BiologicalParent"
	ChildOrder         Fact = "http://gedcomx.org/ChildOrder"
	EnteringHeir       Fact = "http://gedcomx.org/EnteringHeir"
	ExitingHeir        Fact = "http://gedcomx.org/ExitingHeir"
	FosterParent       Fact = "http://gedcomx.org/FosterParent"
	GuardianParent     Fact = "http://gedcomx.org/GuardianParent"
	StepParent         Fact = "http://gedcomx.org/StepParent"
	SociologicalParent Fact = "http://gedcomx.org/SociologicalParent"
	SurrogateParent    Fact = "http://gedcomx.org/SurrogateParent"
)
