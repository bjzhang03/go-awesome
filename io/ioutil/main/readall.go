package main

import (
	"strings"
	"io/ioutil"
	"log"
	"fmt"
)

/***
ReadAll reads from r until an error or EOF and returns the data it read.
A successful call returns err == nil, not err == EOF.
Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
 */
func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
