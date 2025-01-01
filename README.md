# git-bundle [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/git-bundle.svg)](https://pkg.go.dev/github.com/bored-engineer/git-bundle)
A self-contained Golang package for parsing (and serializing) the [Git Bundle Format](https://git-scm.com/docs/bundle-format) `v2` and `v3`, ex:
```go
package main

import (
	"bufio"
	"log"
	"os"

	gitbundle "github.com/bored-engineer/git-bundle"
)

func main() {
	f, err := os.Open("<path to bundle file>")
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	br := bufio.NewReader(f)

	bundle, err := gitbundle.Parse(br)
	if err != nil {
		log.Fatalf("gitbundle.Parse failed: %v", err)
	}
	for refname, objID := range bundle.References.Map() {
		log.Printf("%s: %s", refname, objID)
	}
}
```
When [Parse](https://pkg.go.dev/github.com/bored-engineer/git-bundle#Parse) returns the `*bufio.Reader` position will be at the start of the [git packfile](https://git-scm.com/book/en/v2/Git-Internals-Packfiles) section and can be directly read/used by another package such as [go-git's 
plumbing/format/packfile package](https://pkg.go.dev/github.com/go-git/go-git/v5@v5.13.0/plumbing/format/packfile). Similarly, when [(*Bundle).WriteTo](https://pkg.go.dev/github.com/bored-engineer/git-bundle#Bundle.WriteTo) is used only the bundle header will be written to the provided `io.Writer`, appending the actual packfile contents is left as an exercise for the user.
