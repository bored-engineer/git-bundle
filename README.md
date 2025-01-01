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

	// The *bufio.Reader will be at the start of the packfile when Parse returns (without error)
	bundle, err := gitbundle.Parse(br)
	if err != nil {
		log.Fatalf("gitbundle.Parse failed: %v", err)
	}
	for refname, objID := range bundle.References.Map() {
		log.Printf("%s: %s", refname, objID)
	}
}
```
