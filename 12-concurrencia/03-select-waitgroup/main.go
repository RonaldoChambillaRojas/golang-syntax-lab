package main

import (
	"fmt"
	"time"
)

// =========================================================
// SELECT — Multiplexación de channels
// =========================================================
// select espera en múltiples operaciones de channel simultáneamente
// y ejecuta el primero que esté listo.
// Es como un switch, pero para operaciones de I/O concurrentes.
//
// En JS sería como: await Promise.race([p1, p2, p3])

func main() {
	// --- select básico ---
	fmt.Println("=== Select básico ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "resultado de ch1"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "resultado de ch2"
	}()

	// Espera el primero que llegue
	select {
	case msg := <-ch1:
		fmt.Println("ch1 ganó:", msg)
	case msg := <-ch2:
		fmt.Println("ch2 ganó:", msg) // este gana porque es más rápido
	}

	// --- Select con timeout ---
	fmt.Println("\n=== Select con timeout ===")

	ch := make(chan int)
	go func() {
		time.Sleep(200 * time.Millisecond) // tarda mucho
		ch <- 42
	}()

	select {
	case valor := <-ch:
		fmt.Println("Valor recibido:", valor)
	case <-time.After(100 * time.Millisecond): // timeout de 100ms
		fmt.Println("Timeout: la operación tardó demasiado")
	}

	// --- Select no bloqueante con default ---
	fmt.Println("\n=== Select no bloqueante ===")

	ready := make(chan bool, 1)

	select {
	case <-ready:
		fmt.Println("Estaba listo")
	default:
		fmt.Println("No listo, seguimos sin esperar") // este caso
	}

	// --- Fan-out: un productor → múltiples consumidores ---
	fmt.Println("\n=== Fan-out ===")

	trabajo := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		trabajo <- i
	}
	close(trabajo)

	resultados := make(chan int, 10)
	for w := 0; w < 3; w++ { // 3 workers
		go func(id int) {
			for j := range trabajo {
				resultados <- j * j
			}
		}(w)
	}

	// Recoger resultados con timeout
	total := 0
	esperados := 10
	recibidos := 0
	timeout := time.After(time.Second)

loop:
	for recibidos < esperados {
		select {
		case r := <-resultados:
			total += r
			recibidos++
		case <-timeout:
			fmt.Println("Timeout esperando resultados")
			break loop
		}
	}
	fmt.Println("Suma de cuadrados 1-10:", total) // 385

	// --- Canal de cancelación ---
	fmt.Println("\n=== Canal de cancelación ===")

	cancelar := make(chan struct{})
	datos := make(chan int)

	// Goroutine que genera datos hasta que se cancele
	go func() {
		defer close(datos)
		for i := 0; ; i++ {
			select {
			case datos <- i:
			case <-cancelar:
				fmt.Println("  Generador cancelado")
				return
			}
		}
	}()

	// Leer 5 valores y luego cancelar
	for i := 0; i < 5; i++ {
		fmt.Println(" dato:", <-datos)
	}
	close(cancelar) // señala la cancelación
	time.Sleep(10 * time.Millisecond) // pequeña espera para ver el mensaje
}
