/* Documentación: https://go.dev/doc/effective_go#if */

package main

import "fmt"

// Función principal
func main() {
	param01 := 1
	param02 := 2

	// Condición clasica
	if param01 == param02 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}

	// Condición "Y"
	if param01 == 1 && param02 == 2 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}

	// Condición "O"
	if param01 == 10 || param02 == 2 {
		fmt.Println("Verdad")
	} else {
		fmt.Println("Falso")
	}
}
