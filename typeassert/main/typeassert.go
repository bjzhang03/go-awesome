package main

import (
	"fmt"
)

type Html []interface{}

func main() {
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		if value, ok := element.(string); ok {
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}
	}
}
