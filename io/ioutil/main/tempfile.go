package main

import (
	"io/ioutil"
	"log"
	"os"
)

/***
TempFile creates a new temporary file in the directory dir with a name beginning with prefix,
opens the file for reading and writing, and returns the resulting *os.File.
If dir is the empty string, TempFile uses the default directory for temporary files (see os.TempDir).
Multiple programs calling TempFile simultaneously will not choose the same file.
The caller can use f.Name() to find the pathname of the file.
It is the caller's responsibility to remove the file when no longer needed.
 */

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
