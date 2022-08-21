/* Documentación: */

package main

import "fmt"

type figura interface {
	area() float64
}

// Crear clase privada
type cuadrado struct {
	base float64
}

// Crear clase privada
type rectangulo struct {
	base   float64
	altura float64
}

// Crear función
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// Crear función
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// Crear función
func calcular(f figura) {
	fmt.Println("El area es:", f.area())
}

// Crear función
func main() {
	// Instanciar clases struct
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	// Ejecutar función
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Listar interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface)
}
