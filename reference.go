package gitbundle

import (
	"bytes"
	"fmt"
)

// reference    = obj-id SP refname LF
type Reference struct {
	ObjectID ObjectID
	Name     string
}

// Append appends the reference line to the given buffer
func (r Reference) Append(b []byte) []byte {
	return append(append(append(b, r.ObjectID...), ' '), []byte(r.Name)...)
}

// Bytes returns the reference line as a byte slice
func (r Reference) Bytes() []byte {
	return r.Append(nil)
}

// String implements the fmt.Stringer interface
func (r Reference) String() string {
	return string(r.Bytes())
}

// ParseReference parses a reference from a line of a bundle file
func ParseReference(line []byte) (Reference, error) {
	objID, name, ok := bytes.Cut(line, []byte(" "))
	if !ok {
		return Reference{}, fmt.Errorf("invalid reference line: %q", line)
	}
	return Reference{
		ObjectID: ObjectID(objID),
		Name:     string(name),
	}, nil
}

// References is a list of references
type References []Reference

// Append appends the references to the given buffer
func (rr References) Append(b []byte) []byte {
	for _, ref := range rr {
		b = append(ref.Append(b), '\n')
	}
	return b
}

// Bytes returns the references as a byte slice
func (rr References) Bytes() []byte {
	return rr.Append(nil)
}

// String implements the fmt.Stringer interface
func (rr References) String() string {
	return string(rr.Bytes())
}

// Map converts the slice into a map by reference name
func (rr References) Map() map[string]ObjectID {
	m := make(map[string]ObjectID, len(rr))
	for _, ref := range rr {
		m[ref.Name] = ref.ObjectID
	}
	return m
}
