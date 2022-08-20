/* Documentación: https://go.dev/doc/effective_go#maps */
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Obtener el arreglo de nombre de usuarios
	users := []string{"Alex", "Deysi", "Diana", "Monica", "Jenny", "David", "Ana"}

	// Iterar usuarios
	for index, user := range users {
		fmt.Println(index, user)
		isPalindromo(strings.ToLower(user))
	}

}

func isPalindromo(param01 string) {
	var textReverse string
	count := len(param01) // Longitud de string [A,l,e,x]
	for i := count - 1; i >= 0; i-- {
		textReverse += string(param01[i])
	}

	if param01 == textReverse {
		fmt.Println("Es un palindromo", textReverse)
	} else {
		fmt.Println("No es un palindromo", textReverse)

	}
}
