package gitbundle

import (
	"bytes"
	"fmt"
)

// capability   = "@" key ["=" value] LF
// key          = 1*(ALPHA / DIGIT / "-")
// value        = *(%01-09 / %0b-FF)
type Capability struct {
	Key   string
	Value []byte
}

// Append appends the capability line to the given buffer
func (c Capability) Append(b []byte) []byte {
	if c.Value == nil {
		return append(append(b, '@'), c.Key...)
	}
	return append(append(append(append(b, '@'), c.Key...), '='), c.Value...)
}

// Bytes returns the capability line as a byte slice
func (c Capability) Bytes() (line []byte) {
	return c.Append(nil)
}

// String implements the fmt.Stringer interface
func (c Capability) String() string {
	return string(c.Bytes())
}

// ParseCapability parses a capability from a line of a bundle file
func ParseCapability(line []byte) (Capability, error) {
	if len(line) < 2 || line[0] != '@' || line[1] == '=' {
		return Capability{}, fmt.Errorf("invalid capability line: %q", line)
	}
	key, value, _ := bytes.Cut(line[1:], []byte("="))
	return Capability{
		Key:   string(key),
		Value: value,
	}, nil
}

// "Capabilities", which are only in the v3 format, indicate functionality that the bundle requires to be read properly.
type Capabilities []Capability

// Append appends the capabilities to the given buffer
func (cc Capabilities) Append(b []byte) []byte {
	for _, c := range cc {
		b = append(c.Append(b), '\n')
	}
	return b
}

// Bytes returns the capabilities as a byte slice
func (cc Capabilities) Bytes() []byte {
	return cc.Append(nil)
}

// String implements the fmt.Stringer interface
func (cc Capabilities) String() string {
	return string(cc.Bytes())
}

// Get returns the value of the given key in the capabilities
func (cc Capabilities) Get(key string) (value []byte, ok bool) {
	for _, c := range cc {
		if c.Key == key {
			return c.Value, true
		}
	}
	return nil, false
}

// Has returns true if the given key is present in the capabilities
func (cc Capabilities) Has(key string) bool {
	for _, c := range cc {
		if c.Key == key {
			return true
		}
	}
	return false
}
