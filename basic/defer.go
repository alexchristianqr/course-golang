/* Documentación: https://go.dev/doc/effective_go#defer */
package main

import "fmt"

func main() {
	// Ejemplo de defer
	defer printHelloWorld()

	result := true
	if result {
		fmt.Println("Es verdad")
	}
}

func printHelloWorld() {
	fmt.Println("Hola mundo")
}
