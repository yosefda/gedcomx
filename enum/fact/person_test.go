package fact_test

import (
	"reflect"
	"testing"

	"github.com/yosefda/gedcomx/enum/fact"
)

func TestPersonFact_CreateWithApplicableFact(t *testing.T) {
	f, err := fact.NewPersonFact(fact.TaxAssessment)
	if err != nil {
		t.Fatalf("Failed to create PersonFact with %s", fact.TaxAssessment)
	}

	instanceOf := reflect.TypeOf(f)
	if instanceOf.String() != "*fact.Person" {
		t.Fatalf("Instance is not of type *fact.Person")
	}
}

func TestPersonFact_CreateWithUnapplicableFact(t *testing.T) {
	_, err := fact.NewPersonFact(fact.AdoptiveParent)
	if err == nil {
		t.Fatalf("Must not create PersonFact with %s", fact.AdoptiveParent)
	}
}

func TestPersonFact_GetFact(t *testing.T) {
	pf, err := fact.NewPersonFact(fact.TaxAssessment)
	if err != nil {
		t.Fatalf("Failed to create PersonFact with %s", fact.TaxAssessment)
	}

	if pf.Get() != fact.TaxAssessment {
		t.Fatalf("Failed to return %s", fact.TaxAssessment)
	}
}

func TestPersonFact_ComparingWithFactEnum(t *testing.T) {
	pf, err := fact.NewPersonFact(fact.TaxAssessment)
	if err != nil {
		t.Fatalf("Failed to create PersonFact with %s", fact.TaxAssessment)
	}

	if pf.Get() != fact.TaxAssessment {
		t.Fatalf("Must be the same as %s", fact.TaxAssessment)
	}

	if pf.Get() == fact.AdoptiveParent {
		t.Fatalf("Must be different from %s", fact.AdoptiveParent)
	}
}
