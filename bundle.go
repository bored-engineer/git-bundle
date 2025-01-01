package gitbundle

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Bundle represents a parsed git bundle
type Bundle struct {
	Version       string
	Capabilities  Capabilities
	Prerequisites Prerequisites
	References    References
}

// Append appends the bundle to the given buffer
func (b *Bundle) Append(buf []byte) []byte {
	buf = append(buf, "# v"...)
	buf = append(buf, b.Version...)
	buf = append(buf, " git bundle\n"...)
	buf = b.Capabilities.Append(buf)
	buf = b.Prerequisites.Append(buf)
	buf = b.References.Append(buf)
	buf = append(buf, '\n')
	return buf
}

// Bytes returns the bundle as a byte slice
func (b *Bundle) Bytes() []byte {
	return b.Append(nil)
}

// String implements the fmt.Stringer interface
func (b *Bundle) String() string {
	return string(b.Bytes())
}

// WriteTo writes the bundle to the given io.Writer
func (b *Bundle) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.Bytes())
	return int64(n), err
}

// Parse parses a bundle from the given bufio.Reader
func Parse(r *bufio.Reader) (*Bundle, error) {
	// signature = "# v3 git bundle" LF
	versionLine, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	version, ok := strings.CutPrefix(versionLine, "# v")
	if !ok {
		return nil, fmt.Errorf("invalid bundle version: %q", versionLine)
	}
	version, ok = strings.CutSuffix(version, " git bundle\n")
	if !ok {
		return nil, fmt.Errorf("invalid bundle version: %q", versionLine)
	}
	switch version {
	default:
		return nil, fmt.Errorf("unsupported bundle version: %q", version)
	case "2", "3":
		// We support only version 2 and 3 bundles right now
	}
	bundle := &Bundle{
		Version: version,
	}
	for {
		// Read until the next LF
		line, err := r.ReadBytes('\n')
		if err != nil {
			return nil, err
		} else if len(line) == 1 && line[0] == '\n' {
			// An empty LF line means the rest of the data is the packfile
			return bundle, nil
		} else if len(line) <= 1 {
			return nil, fmt.Errorf("invalid line: %q", line)
		}
		// Trim the trailing LF
		line = line[:len(line)-1]
		// Based on the first character, we can determine the type of line
		switch line[0] {
		case '@':
			if bundle.Prerequisites != nil || bundle.References != nil {
				return nil, fmt.Errorf("capabilities must come first")
			}
			if version == "2" {
				return nil, fmt.Errorf("capabilities are not supported in version 2 bundles")
			}
			capability, err := ParseCapability(line)
			if err != nil {
				return nil, err
			}
			bundle.Capabilities = append(bundle.Capabilities, capability)
		case '-':
			if bundle.References != nil {
				return nil, fmt.Errorf("prerequisites must come first")
			}
			prerequisite, err := ParsePrerequisite(line)
			if err != nil {
				return nil, err
			}
			bundle.Prerequisites = append(bundle.Prerequisites, prerequisite)
		default:
			reference, err := ParseReference(line)
			if err != nil {
				return nil, err
			}
			bundle.References = append(bundle.References, reference)
		}
	}
}
