/* Documentación: https://go.dev/doc/effective_go#maps */

package main

import (
	"fmt"
)

func main() {
	// Obtener el arreglo de nombre de usuarios
	users := []string{"Alex", "Deysi", "Diana", "Monica", "Jenny", "David", "Ana"}

	// Iterar usuarios
	for index, user := range users {
		fmt.Println(index, user)
		// isPalindromo(strings.ToLower(user))
	}

	fmt.Println("-- break --")

	// Crear diccionario
	languages := make(map[string]string)
	languages["java"] = "multiparadigma"
	languages["php"] = "proposito general"
	fmt.Println("El diccionario es:", languages)

	fmt.Println("-- break --")

	// Encontrar un valor en el diccionario
	v, k := languages["nodejs"]
	fmt.Println("El resultado es:", v, k)

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
