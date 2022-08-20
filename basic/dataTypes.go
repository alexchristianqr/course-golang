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

	// Variables
	x := 10
	y := 50

	// Sumar
	result := x + y
	fmt.Println("suma:", result)

	// Restar
	result = y - x // Como la variable result ya esta definida, solo reasignar con "="
	fmt.Println("resta:", result)

	// Multiplicar
	result = x * y
	fmt.Println("multiplicación:", result)

	// Dividir
	result = y / x // 50/10
	fmt.Println("división:", result)

	// Modulo
	result = (x % y) // El modulo es una división exacta
	fmt.Println("modulo:", result)

	// Incrementar valor
	x++
	fmt.Println("incremento de X:", x)

	// Disminuir valor
	x--
	fmt.Println("decremento de X:", x)

}
