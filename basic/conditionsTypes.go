package main

import (
	"fmt"
)

func main() {
	param01 := 1
	param02 := 2

	// Clasico
	if param01 == param02 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}

	// And
	if param01 == 1 && param02 == 2 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}

	// Or
	if param01 == 10 || param02 == 2 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}
}
