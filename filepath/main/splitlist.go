package main

import (
	"fmt"
	"path/filepath"
)

/***
SplitList splits a list of paths joined by the OS-specific ListSeparator,
usually found in PATH or GOPATH environment variables.
Unlike strings.Split, SplitList returns an empty slice when passed an empty string.
 */
func main() {
	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
}
