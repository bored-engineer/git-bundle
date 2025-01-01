package gitbundle

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

const v2 = `# v2 git bundle
-abcdef0123456789abcdef0123456789abcdef01 comment
0123456789abcdef0123456789abcdef01234567 refs/heads/master
abcdef0123456789abcdef0123456789abcdef01 refs/heads/develop

`

const v3 = `# v3 git bundle
@simple
@complex=value
-abcdef0123456789abcdef0123456789abcdef01 comment
0123456789abcdef0123456789abcdef01234567 refs/heads/master
abcdef0123456789abcdef0123456789abcdef01 refs/heads/develop

`

func TestBundle(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    Bundle
		wantErr string
	}{
		"empty": {
			input:   "",
			wantErr: "EOF",
		},
		"not a bundle": {
			input:   "not a bundle\n",
			wantErr: `invalid bundle version: "not a bundle\n"`,
		},
		"invalid version": {
			input:   "# v1337 git bundle\n",
			wantErr: `unsupported bundle version: "1337"`,
		},
		"v2": {
			input: v2,
			want: Bundle{
				Version: "2",
				Prerequisites: Prerequisites{
					{
						ObjectID: "abcdef0123456789abcdef0123456789abcdef01",
						Comment:  "comment",
					},
				},
				References: References{
					{
						ObjectID: "0123456789abcdef0123456789abcdef01234567",
						Name:     "refs/heads/master",
					},
					{
						ObjectID: "abcdef0123456789abcdef0123456789abcdef01",
						Name:     "refs/heads/develop",
					},
				},
			},
		},
		"v3": {
			input: v3,
			want: Bundle{
				Version: "3",
				Capabilities: Capabilities{
					{
						Key: "simple",
					},
					{
						Key:   "complex",
						Value: []byte("value"),
					},
				},
				Prerequisites: Prerequisites{
					{
						ObjectID: "abcdef0123456789abcdef0123456789abcdef01",
						Comment:  "comment",
					},
				},
				References: References{
					{
						ObjectID: "0123456789abcdef0123456789abcdef01234567",
						Name:     "refs/heads/master",
					},
					{
						ObjectID: "abcdef0123456789abcdef0123456789abcdef01",
						Name:     "refs/heads/develop",
					},
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			have, err := Parse(bufio.NewReader(bytes.NewReader([]byte(tc.input))))
			if tc.wantErr != "" {
				if err == nil || err.Error() != tc.wantErr {
					t.Fatalf("Parse should fail with %q, got %v", tc.wantErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Parse should not fail, got %v", err)
				}
				if !reflect.DeepEqual(have, &tc.want) {
					t.Fatalf("Parse should have returned %#v, got %#v", tc.want, have)
				}
				if have := have.String(); have != tc.input {
					t.Errorf("String should be %q, got %q", tc.input, have)
				}
			}
		})
	}
}
