package main

import (
	"fmt"

	"github.com/kybouw/luhn/luhn"
)

func main() {
	fmt.Println("Hello, welcome to luhn!")

	var testString string = "17893729974"
	fmt.Printf("Validating: %s ... ", testString)
	if luhn.Verify(testString) {
		fmt.Print("GOOD\n")
	} else {
		fmt.Print("BAD\n")
	}

}
