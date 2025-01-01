package gitbundle

// obj-id    =  40*(HEXDIGIT)
// HEXDIG    =  DIGIT / "a" / "b" / "c" / "d" / "e" / "f"
type ObjectID string

// String implements the fmt.Stringer interface
func (oid ObjectID) String() string {
	return string(oid)
}

// Valid returns true if the object ID is valid
func (oid ObjectID) Valid() bool {
	if len(oid) != 40 {
		return false
	}
	for _, c := range oid {
		if !('0' <= c && c <= '9') && !('a' <= c && c <= 'f') {
			return false
		}
	}
	return true
}
