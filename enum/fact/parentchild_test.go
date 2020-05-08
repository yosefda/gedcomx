package fact_test

import (
	"reflect"
	"testing"

	"github.com/yosefda/gedcomx/enum/fact"
)

func TestParentChildFact_CreateWithApplicableFact(t *testing.T) {
	f, err := fact.NewParentChildFact(fact.AdoptiveParent)
	if err != nil {
		t.Fatalf("Failed to create ParentChildFact with %s", fact.AdoptiveParent)
	}

	instanceOf := reflect.TypeOf(f)
	if instanceOf.String() != "*fact.ParentChild" {
		t.Fatalf("Instance is not of type *fact.ParentChild")
	}
}

func TestParentChildFact_CreateWithUnapplicableFact(t *testing.T) {
	_, err := fact.NewParentChildFact(fact.TaxAssessment)
	if err == nil {
		t.Fatalf("Must not create ParentChildFact with %s", fact.TaxAssessment)
	}
}

func TestParentChildFact_GetFact(t *testing.T) {
	pf, err := fact.NewParentChildFact(fact.AdoptiveParent)
	if err != nil {
		t.Fatalf("Failed to create ParentChildFact with %s", fact.AdoptiveParent)
	}

	if pf.Get() != fact.AdoptiveParent {
		t.Fatalf("Failed to return %s", fact.AdoptiveParent)
	}
}

func TestParentChildFact_ComparingWithFactEnum(t *testing.T) {
	pf, err := fact.NewParentChildFact(fact.AdoptiveParent)
	if err != nil {
		t.Fatalf("Failed to create ParentChildFact with %s", fact.AdoptiveParent)
	}

	if pf.Get() != fact.AdoptiveParent {
		t.Fatalf("Must be the same as %s", fact.AdoptiveParent)
	}

	if pf.Get() == fact.TaxAssessment {
		t.Fatalf("Must be different from %s", fact.TaxAssessment)
	}
}
