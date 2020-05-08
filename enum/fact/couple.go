// Package fact specifies a set of enumerated values that provides the means for descriptions
// of facts about couple relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package fact

import (
	"fmt"
)

// Facts that are applicable to a couple.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#22-couple-relationship-fact-types
var coupleFacts = map[Fact]struct{}{
	Annulment:           void,
	CommonLawMarriage:   void,
	CivilUnion:          void,
	Divorce:             void,
	DivorceFiling:       void,
	DomesticPartnership: void,
	Engagement:          void,
	Marriage:            void,
	MarriageBanns:       void,
	MarriageContract:    void,
	MarriageLicense:     void,
	MarriageNotice:      void,
	Separation:          void,
	NumberOfChildren:    void,
}

// Couple defines fact that is applicable to a couple.
type Couple struct {
	fact Fact
}

// Get returns fact about a couple.
func (c *Couple) Get() Fact {
	return c.fact
}

// NewCoupleFact create Couple fact type.
func NewCoupleFact(fact Fact) (*Couple, error) {
	if _, ok := coupleFacts[fact]; !ok {
		return nil, fmt.Errorf("Fact %s is not applicable to a couple", fact)
	}

	return &Couple{fact}, nil
}
