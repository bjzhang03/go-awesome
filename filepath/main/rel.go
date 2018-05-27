package main

import (
	"fmt"
	"path/filepath"
)

/***
Rel returns a relative path that is lexically equivalent to targpath when joined to basepath with an intervening separator.
That is, Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
On success, the returned path will always be relative to basepath, even if basepath and targpath share no elements.
An error is returned if targpath can't be made relative to basepath or if knowing the current working directory would be necessary to compute it.
Rel calls Clean on the result.
 */
func main() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

}
