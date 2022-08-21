/* Documentación: */

package main

import "fmt"

// Pc : Crear clase
type Pc struct {
	ram   int
	disk  int
	brand string
}

// Crear puntero a memoria
func (pc Pc) ping() {
	fmt.Println(pc.ram, "Pong")
}

// Crear puntero a memoria
func (pc *Pc) duplicateRAM() {
	pc.ram = pc.ram * 2
}

func main() {
	// Ejemplo 01
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100       // Actualizar "b" y afecta "a"
	fmt.Println(a) // "a" tambien se actualiza

	// Ejemplo 02
	pc := Pc{ram: 16, disk: 500, brand: "Asus"}
	pc.ping()

	fmt.Println(pc)
	pc.duplicateRAM()
	fmt.Println(pc)
}
