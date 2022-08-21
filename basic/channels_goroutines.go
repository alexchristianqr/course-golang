/* Documentación: */

package main

import (
	"fmt"
	"time"
)

// Función solo para recibir
// func sayText(text string, channel chan<- string)

// Función solo para salir
// func sayText(text string, channel <-chan string)

func sayText(text string, channel chan<- string) {
	channel <- text // Actualizar el texto en el canal
}

// Función principal
func main() {
	// Crear un canal de string de tamaño "1"
	channel := make(chan string, 1)

	// Log
	fmt.Println("Hola", time.Now())

	// Ejecutar
	go sayText("Bye", channel)

	// Imprimir variable actualizada
	fmt.Println(<-channel, time.Now())
}
