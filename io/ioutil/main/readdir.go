package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

/***
ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename.
 */
func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
