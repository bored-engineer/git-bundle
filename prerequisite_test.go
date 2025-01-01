package gitbundle

import (
	"reflect"
	"testing"
)

func TestPrerequisite(t *testing.T) {
	req, err := ParsePrerequisite([]byte("-0123456789abcdef0123456789abcdef01234567 optional comment"))
	if err != nil {
		t.Fatalf("ParsePrerequisite failed: %v", err)
	}
	if req.ObjectID != ObjectID("0123456789abcdef0123456789abcdef01234567") {
		t.Errorf("ObjectID should be '0123456789abcdef0123456789abcdef01234567', got %q", req.ObjectID)
	}
	if req.Comment != "optional comment" {
		t.Errorf("Comment should be 'optional comment', got %q", req.Comment)
	}
	if _, err := ParsePrerequisite([]byte("invalid")); err == nil {
		t.Errorf("ParsePrerequisite should fail for invalid input")
	}
}

func TestPrerequisites(t *testing.T) {
	reqs := Prerequisites{
		{ObjectID: ObjectID("0123456789abcdef0123456789abcdef01234567"), Comment: "optional comment"},
		{ObjectID: ObjectID("abcdef0123456789abcdef0123456789abcdef01"), Comment: "required"},
	}
	if !reflect.DeepEqual(reqs.Map(), map[ObjectID]string{
		"0123456789abcdef0123456789abcdef01234567": "optional comment",
		"abcdef0123456789abcdef0123456789abcdef01": "required",
	}) {
		t.Errorf("Map should be identical: %v", reqs.Map())
	}
}
