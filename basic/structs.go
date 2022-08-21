/* Documentación: */

package main

import "fmt"

type student struct {
	name string
	age  int
}

// Función principal
func main() {
	// Ejemplo 01
	myStudent := student{name: "Alex Christian", age: 28}
	fmt.Println(myStudent)

	fmt.Println("-- break --")

	// Ejemplo 02
	var otherStudent student
	otherStudent.name = "Luis Miguel"
	otherStudent.age = 26
	fmt.Println(otherStudent)
}
