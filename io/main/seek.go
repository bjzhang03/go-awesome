package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 16)

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 6)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

}