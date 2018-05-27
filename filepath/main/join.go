package main

import (
"fmt"
"path/filepath"
)

/***
Join joins any number of path elements into a single path, adding a Separator if necessary.
Join calls Clean on the result; in particular, all empty strings are ignored.
On Windows, the result is a UNC path if and only if the first path element is a UNC path.
 */
func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
}
