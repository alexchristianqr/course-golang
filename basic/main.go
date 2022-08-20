package main

import "fmt"

func main() {
	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	// Impresión en consola
	fmt.Println("pi:", pi)
	fmt.Println("pi2", pi2)

	// Variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	// Valores por defecto
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a:", a) // 0
	fmt.Println("b:", b) // 0
	fmt.Println("c:", c) // empty o vacio
	fmt.Println("d:", d) // false

	// Area y base de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado:", areaCuadrado)
}
