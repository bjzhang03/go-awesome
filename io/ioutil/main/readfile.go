package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

/***
ReadFile reads the file named by filename and returns the contents.
A successful call returns err == nil, not err == EOF.
Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
 */
func main() {
	//content, err := ioutil.ReadFile("testdata/hello")
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
}
