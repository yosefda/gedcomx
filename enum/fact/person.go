// Package fact specifies a set of enumerated values that provides the means for descriptions
// of facts about couple relationships.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md
package fact

import (
	"fmt"
)

// Facts that are applicable to a person.
// https://github.com/FamilySearch/gedcomx/blob/master/specifications/fact-types-specification.md#21-person-fact-types
var personFacts = map[Fact]struct{}{
	Adoption:                  void,
	AdultChristening:          void,
	Amnesty:                   void,
	AncestralHall:             void,
	AncestralPoem:             void,
	Apprenticeship:            void,
	Arrest:                    void,
	Award:                     void,
	Baptism:                   void,
	BarMitzvah:                void,
	BatMitzvah:                void,
	Birth:                     void,
	BirthNotice:               void,
	Blessing:                  void,
	Branch:                    void,
	Burial:                    void,
	Caste:                     void,
	Census:                    void,
	Christening:               void,
	Circumcision:              void,
	Clan:                      void,
	Confirmation:              void,
	Court:                     void,
	Cremation:                 void,
	Death:                     void,
	Education:                 void,
	EducationEnrollment:       void,
	Emigration:                void,
	Enslavement:               void,
	Ethnicity:                 void,
	Excommunication:           void,
	FirstCommunion:            void,
	Funeral:                   void,
	GenderChange:              void,
	GenerationNumber:          void,
	Graduation:                void,
	Heimat:                    void,
	Immigration:               void,
	Imprisonment:              void,
	Inquest:                   void,
	LandTransaction:           void,
	Language:                  void,
	Living:                    void,
	MaritalStatus:             void,
	Medical:                   void,
	MilitaryAward:             void,
	MilitaryDischarge:         void,
	MilitaryDraftRegistration: void,
	MilitaryInduction:         void,
	MilitaryService:           void,
	Mission:                   void,
	MoveFrom:                  void,
	MoveTo:                    void,
	MultipleBirth:             void,
	NationalID:                void,
	Nationality:               void,
	Naturalization:            void,
	NumberOfChildren:          void,
	NumberOfMarriages:         void,
	Obituary:                  void,
	OfficialPosition:          void,
	Occupation:                void,
	Ordination:                void,
	Pardon:                    void,
	PhysicalDescription:       void,
	Probate:                   void,
	Property:                  void,
	Race:                      void,
	Religion:                  void,
	Residence:                 void,
	Retirement:                void,
	Stillbirth:                void,
	TaxAssessment:             void,
	Tribe:                     void,
	Will:                      void,
	Visit:                     void,
	Yahrzeit:                  void,
}

// Person defines fact that is applicable to a person.
type Person struct {
	fact Fact
}

// Get returns fact about a person.
func (p *Person) Get() Fact {
	return p.fact
}

// NewPersonFact create Person fact type.
func NewPersonFact(fact Fact) (*Person, error) {
	if _, ok := personFacts[fact]; !ok {
		return nil, fmt.Errorf("Fact %s is not applicable to a person", fact)
	}

	return &Person{fact}, nil
}
