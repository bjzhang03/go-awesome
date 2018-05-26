package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// Note that the mapstructure tags defined in the struct type
// can indicate which fields the values are mapped to.
type PersonFive struct {
	Name string `mapstructure:"person_name"`
	Age  int    `mapstructure:"person_age"`
}

func main() {
	input := map[string]interface{}{
		"person_name": "Mitchell",
		"person_age":  91,
	}

	var result PersonFive
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)

}
