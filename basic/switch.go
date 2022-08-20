/* Documentación: https://go.dev/doc/effective_go#switch */

package main

import "fmt"

func main() {
	param01 := 4
	param02 := 2

	// Ejemplo 01
	switch modulo := param01 % param02; modulo {
	case 0:
		fmt.Println("resultado: es impar")
	case 1:
		fmt.Println("resultado: es par")
	default:
		fmt.Println("resultado: valor no evaluado")
	}

	// Ejemplo 02
	result := param01 + param02
	switch {
	case result > 10:
		fmt.Println("Es mayor a 10")
	case result < 10:
		fmt.Println("Es menor a 10")
	default:
		fmt.Println("No condición")
	}

}
