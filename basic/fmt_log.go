/* Documentación: https://pkg.go.dev/fmt */

package main

import "fmt"

// Función principal
func main() {
	// Declaración de variables
	title := "Platzi"
	quantity := 500

	// Utilizando Println
	fmt.Println(title, quantity)

	// Utilizando Printf
	fmt.Printf("%s tienes más de %d cursos\n", title, quantity)
	fmt.Printf("%v tienes más de %v cursos\n", title, quantity)

	// Utilizando Sprintf
	message := fmt.Sprintf("%v tienes más de %v cursos\n", title, quantity)
	fmt.Println(message)

	// Obtener el tipo de dato por consola
	fmt.Printf("title: %T\n", title)
	fmt.Printf("quantity: %T\n", quantity)
}
