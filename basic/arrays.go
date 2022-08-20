/* Documentación: */
package main

import "fmt"

func main() {
	// Ejemplo de Array
	var dataArray [10]int // Por defecto todos los valores se inicializan con CERO

	// Set valores
	dataArray[0] = 1
	dataArray[1] = 2

	// Imprimir en consola
	fmt.Println("El arreglo es:", dataArray)
	fmt.Println("La longitud o el tamaño es:", len(dataArray))
	fmt.Println("La capacidad es:", cap(dataArray))

	fmt.Println("-- break --")

	// Ejemplo de Slice
	dataSlice := []int{0, 1, 2, 3, 4, 5}

	// Imprimir en consola
	fmt.Println("Posición [0]:", dataSlice[0])
	fmt.Println("Posición [:3]:", dataSlice[:3])
	fmt.Println("Posición [2:4]:", dataSlice[2:4])

}
