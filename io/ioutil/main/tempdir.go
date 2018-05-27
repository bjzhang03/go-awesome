package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/***
TempDir creates a new temporary directory in the directory dir with a name beginning with prefix
and returns the path of the new directory.
If dir is the empty string, TempDir uses the default directory for temporary files (see os.TempDir).
Multiple programs calling TempDir simultaneously will not choose the same directory.
It is the caller's responsibility to remove the directory when no longer needed.
 */
func main() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
