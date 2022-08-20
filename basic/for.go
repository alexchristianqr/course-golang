/* Documentación: https://go.dev/doc/effective_go#for */

package main

import "fmt"

func main() {
	const counter = 10

	// Ciclo for tradicional
	for index := 1; index <= counter; index++ {
		fmt.Println("El indice es:", index)
	}

	fmt.Printf("-- break -- \n")

	// Ciclo for while
	index := 1
	for index <= counter {
		fmt.Println("El indice es:", index)
		index++
	}

	fmt.Printf("-- break -- \n")

	// Ciclo for forever
	index2 := 1
	for {
		fmt.Println("El indice es:", index2)
		index2++ // Incrementar

		if index2 == 3 {
			fmt.Println("Estás en el indice:", index2)
			continue
		}

		if index2 > counter {
			break
		}

	}

	fmt.Printf("-- break -- \n")

	// Ejemplo de ciclo for con indice decreciente
	for index := 10; index > 0; index-- {
		fmt.Println("El indice es:", index)
	}
}
