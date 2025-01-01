package gitbundle

import (
	"reflect"
	"testing"
)

func TestReference(t *testing.T) {
	ref, err := ParseReference([]byte("0123456789abcdef0123456789abcdef01234567 refs/heads/master"))
	if err != nil {
		t.Fatalf("ParseReference failed: %v", err)
	}
	if ref.ObjectID != ObjectID("0123456789abcdef0123456789abcdef01234567") {
		t.Errorf("ObjectID should be '0123456789abcdef0123456789abcdef01234567', got %q", ref.ObjectID)
	}
	if ref.Name != "refs/heads/master" {
		t.Errorf("Name should be 'refs/heads/master', got %q", ref.Name)
	}
	if _, err := ParseReference([]byte("invalid")); err == nil {
		t.Errorf("ParseReference should fail for invalid input")
	}
}

func TestReferences(t *testing.T) {
	refs := References{
		{ObjectID: ObjectID("0123456789abcdef0123456789abcdef01234567"), Name: "refs/heads/master"},
		{ObjectID: ObjectID("abcdef0123456789abcdef0123456789abcdef01"), Name: "refs/heads/develop"},
	}
	if !reflect.DeepEqual(refs.Map(), map[string]ObjectID{
		"refs/heads/master":  ObjectID("0123456789abcdef0123456789abcdef01234567"),
		"refs/heads/develop": ObjectID("abcdef0123456789abcdef0123456789abcdef01"),
	}) {
		t.Errorf("Map should be identical: %v", refs.Map())
	}
}
