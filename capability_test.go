package gitbundle

import (
	"bytes"
	"testing"
)

func TestCapability(t *testing.T) {
	simple, err := ParseCapability([]byte("@key"))
	if err != nil {
		t.Fatalf("ParseCapability failed: %v", err)
	}
	if simple.Key != "key" {
		t.Errorf("Key should be 'key', got %q", simple.Key)
	}
	complex, err := ParseCapability([]byte("@key=value"))
	if err != nil {
		t.Fatalf("ParseCapability failed: %v", err)
	}
	if complex.Key != "key" {
		t.Errorf("Key should be 'key', got %q", complex.Key)
	}
	if !bytes.Equal(complex.Value, []byte("value")) {
		t.Errorf("Value should be 'value', got %q", complex.Value)
	}
	invalid, err := ParseCapability([]byte("invalid"))
	if err == nil {
		t.Errorf("ParseCapability should fail for invalid input: %v", invalid)
	}
}

func TestCapabilities(t *testing.T) {
	caps := Capabilities{
		{Key: "key", Value: []byte("value")},
		{Key: "other", Value: []byte("data")},
	}
	if !caps.Has("key") {
		t.Errorf("Capabilities should have key 'key'")
	}
	if !caps.Has("other") {
		t.Errorf("Capabilities should have key 'other'")
	}
	if caps.Has("missing") {
		t.Errorf("Capabilities should not have key 'missing'")
	}
	value, ok := caps.Get("key")
	if !ok {
		t.Errorf("Capabilities should have key 'key'")
	}
	if !bytes.Equal(value, []byte("value")) {
		t.Errorf("Value should be 'value', got %q", value)
	}
}
