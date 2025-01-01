package gitbundle

import "testing"

func TestObjectID(t *testing.T) {
	valid := ObjectID("0123456789abcdef0123456789abcdef01234567")
	if !valid.Valid() {
		t.Errorf("ObjectID should be valid")
	}
	if valid.String() != "0123456789abcdef0123456789abcdef01234567" {
		t.Errorf("ObjectID should be the same as the input")
	}
	short := ObjectID("short")
	if short.Valid() {
		t.Errorf("ObjectID should be invalid due to length")
	}
	long := ObjectID("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	if long.Valid() {
		t.Errorf("ObjectID should be invalid due to characters")
	}
}
