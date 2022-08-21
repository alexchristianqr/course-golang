/* Documentación: */

package main

import (
	"fmt"
	"sync"
	"time"
)

// Crear función
func singleLog(text string) {
	fmt.Println(text, time.Now())
}

// Crear función
func say(text string, wg *sync.WaitGroup) {
	if text == "#20" {
		defer wg.Done() // Finalizar WaitGroup  al terminar todas las rutinas asíncronas
	}
	fmt.Println(text, time.Now())
}

// Función principal
func main() {
	// Ejemplo 01
	// Ejecutar esta función sincrona
	/*singleLog("#1")
	singleLog("#2")
	singleLog("#3")
	singleLog("#4")
	singleLog("#5")
	singleLog("#6")
	singleLog("#7")
	singleLog("#8")
	singleLog("#9")
	singleLog("#10")
	singleLog("#11")
	singleLog("#12")
	singleLog("#13")
	singleLog("#14")
	singleLog("#15")
	singleLog("#16")
	singleLog("#17")
	singleLog("#18")
	singleLog("#19")
	singleLog("#20")*/

	// Ejemplo 02
	//El paquete sync actua de forma primitiva con las GoRoutines
	var wg sync.WaitGroup

	// Agregar GoRoutine al ciclo del paquete sync
	wg.Add(1)

	// Ejecutar esta función asíncrona
	go say("#1", &wg)
	go say("#2", &wg)
	go say("#3", &wg)
	go say("#4", &wg)
	go say("#5", &wg)
	go say("#6", &wg)
	go say("#7", &wg)
	go say("#8", &wg)
	go say("#9", &wg)
	go say("#10", &wg)
	go say("#11", &wg)
	go say("#12", &wg)
	go say("#13", &wg)
	go say("#14", &wg)
	go say("#15", &wg)
	go say("#16", &wg)
	go say("#17", &wg)
	go say("#18", &wg)
	go say("#19", &wg)
	go say("#20", &wg)

	// Esperar hasta finalizar todas las GoRoutines
	wg.Wait()

	// Ejecutar función al finalizar el WaitGroup
	go singleLog("Adios...")

	//Esperar 1 segundo
	time.Sleep(time.Second * 1)
}
