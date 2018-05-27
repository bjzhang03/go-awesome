package main

import (
	"path/filepath"
	"os"
	"fmt"
)

/***
Walk walks the file tree rooted at root, calling walkFn for each file or directory in the tree, including root.
All errors that arise visiting files and directories are filtered by walkFn.
The files are walked in lexical order,
which makes the output deterministic but means that for very large directories Walk can be inefficient.
Walk does not follow symbolic links.
 */
func main() {
	dir := "dir/to/walk"
	subDirToSkip := "skip" // dir/to/walk/skip

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file: %q\n", path)
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
