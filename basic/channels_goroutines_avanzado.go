/* Documentación: */

package main

import (
	"fmt"
	"time"
)

func message(text string, channel chan string) {
	channel <- text
}

// Función principal
func main() {
	// Crear un canal de rutinas de string de tamaño "2"
	channel := make(chan string, 2)
	channel <- "Mensaje 01"
	channel <- "Mensaje 02"

	// Log
	fmt.Println(len(channel), cap(channel), time.Now())

	// Ejemplo "cerrar" channel
	close(channel)

	// Iterar canales
	for message := range channel {
		fmt.Println(message, time.Now())
	}

	// Ejemplo "select" channel
	channel2 := make(chan string)
	channel3 := make(chan string)
	go message("#2", channel2)
	go message("#3", channel3)

	// Iterar 2 veces
	for i := 0; i < 2; i++ {
		select {
		case ch2 := <-channel2:
			fmt.Println("Canal #2 ejecutado", ch2, time.Now())
		case ch3 := <-channel3:
			fmt.Println("Canal #3 ejecutado", ch3, time.Now())

		}
	}

}
