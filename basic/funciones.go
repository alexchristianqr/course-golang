/* Documentación: https://pkg.go.dev/fmt#pkg-functions */

package main

import "fmt"

// Función principal
func main() {
	// Ejecutar funciones
	normalFunction("Hola Alex")
	tripleArguments("One", "Two", "Tree")
	returnValue("One")

	// Segunda prueba
	_, param2 := doubleReturn("One")
	fmt.Println(param2)
}

func normalFunction(message string) {
	fmt.Println(message)
}

// Función con 3 argumentos
func tripleArguments(param1, param2, param3 string) {
	fmt.Println(param1, param2, param3)
}

// Devolver un valor string
func returnValue(param1 string) string {
	return param1
}

// Devolver multiples valores
func doubleReturn(param1 string) (param2, param3 string) {
	return param1, param2
}
