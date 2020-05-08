// Package fact specifies a set of enumerated values that provides the means for descriptions
// of facts about parent child relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package fact

import (
	"fmt"
)

// Facts that are applicable to a parent child.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#22-couple-relationship-fact-types
var parentChildFacts = map[Fact]struct{}{
	AdoptiveParent:     void,
	BiologicalParent:   void,
	ChildOrder:         void,
	EnteringHeir:       void,
	ExitingHeir:        void,
	FosterParent:       void,
	GuardianParent:     void,
	StepParent:         void,
	SociologicalParent: void,
	SurrogateParent:    void,
}

// ParentChild defines fact that is applicable to a parent child.
type ParentChild struct {
	fact Fact
}

// Get returns fact about a parent child.
func (pc *ParentChild) Get() Fact {
	return pc.fact
}

// NewParentChildFact create Couple fact type.
func NewParentChildFact(fact Fact) (*ParentChild, error) {
	if _, ok := parentChildFacts[fact]; !ok {
		return nil, fmt.Errorf("Fact %s is not applicable to a parent child", fact)
	}

	return &ParentChild{fact}, nil
}
