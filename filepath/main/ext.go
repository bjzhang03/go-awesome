package main

import (
	"fmt"
	"path/filepath"
)

/***
Ext returns the file name extension used by path.
The extension is the suffix beginning at the final dot in the final element of path;
it is empty if there is no dot.
 */
func main() {
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
}
