package main

import (
	"fmt"
	"time"
)

// =========================================================
// CHANNELS — Comunicación entre goroutines
// =========================================================
// Un channel es un tubo tipado por donde las goroutines se envían valores.
// Es el mecanismo idiomático de comunicación en Go.
//
// En JS no hay equivalente directo — lo más cercano son las Promises/Queues.
//
// make(chan T)      → channel sin buffer (sincrónico)
// make(chan T, n)   → channel con buffer de n elementos (asincrónico)

func main() {
	// --- Channel básico (sin buffer) ---
	// Enviar bloquea hasta que alguien recibe, y viceversa.
	// Es como un apretón de manos sincrónico.
	fmt.Println("=== Channel sin buffer ===")

	ch := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "hola desde goroutine" // envía y bloquea hasta que alguien recibe
	}()

	msg := <-ch // recibe y bloquea hasta que alguien envíe
	fmt.Println("Recibido:", msg)

	// --- Channel con buffer ---
	// Los primeros n envíos no bloquean (n = tamaño del buffer).
	// Útil para desacoplar productor y consumidor.
	fmt.Println("\n=== Channel con buffer ===")

	buffered := make(chan int, 3) // buffer de 3
	buffered <- 1                 // no bloquea
	buffered <- 2                 // no bloquea
	buffered <- 3                 // no bloquea
	// buffered <- 4             // bloquearía (buffer lleno)

	fmt.Println("Len:", len(buffered), "Cap:", cap(buffered))
	fmt.Println(<-buffered, <-buffered, <-buffered)

	// --- Patrón Producer-Consumer ---
	fmt.Println("\n=== Producer-Consumer ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Consumer: procesa jobs y pone resultados
	go func() {
		for j := range jobs { // range sobre channel: itera hasta que se cierre
			result := j * j
			results <- result
		}
		close(results) // cierra el canal de resultados cuando terminamos
	}()

	// Producer: envía jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // cierra jobs para indicar "no hay más trabajo"

	// Recoger resultados
	for r := range results {
		fmt.Printf("  resultado: %d\n", r)
	}

	// --- Dirección de canal en tipos ---
	// chan<- T → solo escritura (send-only)
	// <-chan T → solo lectura (receive-only)
	// Útil para expresar intenciones en firmas de funciones
	fmt.Println("\n=== Direcciones de canal ===")
	generador := productor(5)
	consumidor(generador)

	// --- Cerrar un canal ---
	// close(ch) indica que no se enviarán más valores
	// Solo el PRODUCTOR debería cerrar, nunca el consumidor
	// Enviar a un canal cerrado → PANIC
	// Recibir de un canal cerrado retorna (zero, false)
	fmt.Println("\n=== Cerrar canal ===")
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	close(nums)

	v, ok := <-nums
	fmt.Println("Recibido:", v, ok) // 1 true
	v, ok = <-nums
	fmt.Println("Recibido:", v, ok) // 2 true
	v, ok = <-nums
	fmt.Println("Recibido:", v, ok) // 0 false (canal cerrado)
}

// productor genera n números en un channel de solo escritura
func productor(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

// consumidor recibe de un channel de solo lectura
func consumidor(ch <-chan int) {
	for v := range ch {
		fmt.Printf("  consumidor recibió: %d\n", v)
	}
}
