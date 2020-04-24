// Package enum specifies a set of enumerated values that provides the means for descriptions
// of facts about parent-child relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package enum

import "github.com/yosefda/gedcomx/typedef"

// ParentChildFact enum type.
type ParentChildFact typedef.URI

// ParentChildFact is fact that applicable to  a parent-child relationship.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#23-parent-child-relationship-fact-types
const (
	AdoptiveParent     ParentChildFact = "http://gedcomx.org/AdoptiveParent"
	BiologicalParent   ParentChildFact = "http://gedcomx.org/BiologicalParent"
	ChildOrder         ParentChildFact = "http://gedcomx.org/ChildOrder"
	EnteringHeir       ParentChildFact = "http://gedcomx.org/EnteringHeir"
	ExitingHeir        ParentChildFact = "http://gedcomx.org/ExitingHeir"
	FosterParent       ParentChildFact = "http://gedcomx.org/FosterParent"
	GuardianParent     ParentChildFact = "http://gedcomx.org/GuardianParent"
	StepParent         ParentChildFact = "http://gedcomx.org/StepParent"
	SociologicalParent ParentChildFact = "http://gedcomx.org/SociologicalParent"
	SurrogateParent    ParentChildFact = "http://gedcomx.org/SurrogateParent"
)
