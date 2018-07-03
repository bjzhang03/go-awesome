package main

import (
	"net/url"
	"log"
	"fmt"
	"encoding/json"
)

func main() {
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	str, _ := json.Marshal(m)
	fmt.Println(string(str))
}
