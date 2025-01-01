package gitbundle

import (
	"bytes"
	"fmt"
)

// prerequisite = "-" obj-id SP comment LF
type Prerequisite struct {
	ObjectID ObjectID
	Comment  string
}

// Append appends the prerequisite line to the given buffer
func (p Prerequisite) Append(b []byte) []byte {
	return append(append(append(append(b, '-'), p.ObjectID...), ' '), []byte(p.Comment)...)
}

// Bytes returns the prerequisite line as a byte slice
func (p Prerequisite) Bytes() []byte {
	return p.Append(nil)
}

// String implements the fmt.Stringer interface
func (p Prerequisite) String() string {
	return string(p.Bytes())
}

// ParsePrerequisite parses a prerequisite from a line of a bundle file
func ParsePrerequisite(line []byte) (Prerequisite, error) {
	if len(line) < 2 || line[0] != '-' {
		return Prerequisite{}, fmt.Errorf("invalid prequisite line: %q", line)
	}
	objID, comment, _ := bytes.Cut(line[1:], []byte(" "))
	return Prerequisite{
		ObjectID: ObjectID(objID),
		Comment:  string(comment),
	}, nil
}

// "Prerequisites" lists the objects that are NOT included in the bundle and the reader of the bundle MUST already have, in order to use the data in the bundle.
type Prerequisites []Prerequisite

// Append appends the prerequisites to the given buffer
func (pp Prerequisites) Append(b []byte) []byte {
	for _, p := range pp {
		b = append(p.Append(b), '\n')
	}
	return b
}

// Bytes returns the prerequisites as a byte slice
func (pp Prerequisites) Bytes() []byte {
	return pp.Append(nil)
}

// String implements the fmt.Stringer interface
func (pp Prerequisites) String() string {
	return string(pp.Bytes())
}

// Map converts the slice into a map by object id
func (pp Prerequisites) Map() map[ObjectID]string {
	m := make(map[ObjectID]string, len(pp))
	for _, p := range pp {
		m[p.ObjectID] = p.Comment
	}
	return m
}
