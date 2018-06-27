package main

import (
	"fmt"
)

type Html2 []interface{}

func main() {
	html := make(Html2, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		switch value := element.(type) {
		case string:
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		case []byte:
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		case int:
			fmt.Printf("invalid type\n")
		default:
			fmt.Printf("unknown type\n")
		}
	}
}