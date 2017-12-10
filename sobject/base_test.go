package sobject

import (
	"reflect"
	"testing"
)

func TestLookup(t *testing.T) {
	if set, ok := Lookup(ContactObjectName).(ContactSet); !ok {
		t.Errorf("expected ContactSet, actual %s", reflect.TypeOf(set).String())
	}
	if set, ok := Lookup("Knowledge").(KnowledgeSet); !ok {
		t.Errorf("expected KnowledgeSet, actual %s", reflect.TypeOf(set).String())
	}
	if set, ok := Lookup(KnowledgeObjectName).(KnowledgeSet); !ok {
		t.Errorf("expected KnowledgeSet, actual %s", reflect.TypeOf(set).String())
	}
	if res := Lookup("Invalid"); res != nil {
		t.Errorf("expected nil, actual %s", reflect.TypeOf(res).String())
	}
}
