package fact_test

import (
	"reflect"
	"testing"

	"github.com/yosefda/gedcomx/enum/fact"
)

func TestCoupleFact_CreateWithApplicableFact(t *testing.T) {
	f, err := fact.NewCoupleFact(fact.NumberOfChildren)
	if err != nil {
		t.Fatalf("Failed to create CoupleFact with %s", fact.NumberOfChildren)
	}

	instanceOf := reflect.TypeOf(f)
	if instanceOf.String() != "*fact.Couple" {
		t.Fatalf("Instance is not of type *fact.Couple")
	}
}

func TestCoupleFact_CreateWithUnapplicableFact(t *testing.T) {
	_, err := fact.NewCoupleFact(fact.TaxAssessment)
	if err == nil {
		t.Fatalf("Must not create CoupleFact with %s", fact.TaxAssessment)
	}
}

func TestCoupleFact_GetFact(t *testing.T) {
	pf, err := fact.NewCoupleFact(fact.NumberOfChildren)
	if err != nil {
		t.Fatalf("Failed to create CoupleFact with %s", fact.NumberOfChildren)
	}

	if pf.Get() != fact.NumberOfChildren {
		t.Fatalf("Failed to return %s", fact.NumberOfChildren)
	}
}

func TestCoupleFact_ComparingWithFactEnum(t *testing.T) {
	pf, err := fact.NewCoupleFact(fact.NumberOfChildren)
	if err != nil {
		t.Fatalf("Failed to create CoupleFact with %s", fact.NumberOfChildren)
	}

	if pf.Get() != fact.NumberOfChildren {
		t.Fatalf("Must be the same as %s", fact.NumberOfChildren)
	}

	if pf.Get() == fact.TaxAssessment {
		t.Fatalf("Must be different from %s", fact.TaxAssessment)
	}
}
