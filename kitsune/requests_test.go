package kitsune

import (
	"encoding/json"
	"fmt"
)

func ExampleJSONReader() {
	r, err := JSONReader("nyan")
	if err != nil {
		panic(err)
	}
	var neko string
	err = json.NewDecoder(r).Decode(&neko)
	if err != nil {
		panic(err)
	}
	fmt.Print(neko)
	// Output: nyan
}
