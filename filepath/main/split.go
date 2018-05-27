package main

import (
	"fmt"
	"path/filepath"
)

/***
Split splits path immediately following the final Separator, separating it into a directory and file name component.
If there is no Separator in path,
Split returns an empty dir and file set to path. The returned values have the property that path = dir+file.
 */
func main() {
	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}
	fmt.Println("On Unix:")
	for _, p := range paths {
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
}
