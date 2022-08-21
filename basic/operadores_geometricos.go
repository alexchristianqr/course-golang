package main

import "fmt"

// Función principal
func main() {
	// Calcular el área de un cuadrado. Formula: base * base
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado:", areaCuadrado)

	// Calcular el área de un triangulo. Formula: (base * altura) / 2
	const baseTriangulo = 10
	const alturaTriangulo = 25
	areaTriangulo := (baseTriangulo * alturaTriangulo) / 2
	fmt.Println("areaTriangulo:", areaTriangulo)

	// Calcular el área de un circulo. Formula: pi * (radio * radio)
	const pi = 3.14
	const radioCirculo = 4
	areaCirculo := pi * (radioCirculo * radioCirculo)
	fmt.Println("areaCirculo:", areaCirculo)

	// Calcular el área de un trapecio. Formula: (altura * (baseMayor + baseMenor))/ 2
	const alturaTrapecio = 50
	const baseTrapecioMayor = 18
	const baseTrapecioMenor = 7
	areaTrapecio := (alturaTrapecio * (baseTrapecioMayor + baseTrapecioMenor)) / 2
	fmt.Println("areaTrapecio:", areaTrapecio)

}
