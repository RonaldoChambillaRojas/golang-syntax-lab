package main

import (
	"fmt"
	"sync"
	"time"
)

// =========================================================
// GOROUTINES — Concurrencia ligera de Go
// =========================================================
// Una goroutine es como una función que se ejecuta concurrentemente.
// No es un thread del SO: el Go runtime gestiona miles de goroutines
// sobre un número pequeño de threads del SO.
//
// Costo inicial: ~2KB de stack (vs ~1MB de un thread)
// Puedes tener millones de goroutines sin problemas.
//
// En JavaScript: no hay goroutines. Usas async/await con el event loop.

// --- Ejemplo básico ---
func trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done() // notifica al WaitGroup cuando termina

	fmt.Printf("  Trabajador %d iniciando\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond) // simula trabajo
	fmt.Printf("  Trabajador %d terminado\n", id)
}

// --- Race condition: el bug clásico de concurrencia ---
// Ocurre cuando dos goroutines acceden a la misma variable sin sincronización

func demoRaceCondition() {
	fmt.Println("\n=== Race Condition (peligroso) ===")
	contador := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador++ // ✗ RACE CONDITION: múltiples goroutines modifican contador
		}()
	}
	wg.Wait()
	// El resultado esperado es 1000, pero puede ser cualquier número
	fmt.Println("Contador (con race):", contador, "(debería ser 1000, puede no serlo)")
}

// --- La solución: sync.Mutex ---
func demoMutex() {
	fmt.Println("\n=== Con Mutex (seguro) ===")
	contador := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()   // adquiere el lock
			contador++  // operación segura
			mu.Unlock() // libera el lock
		}()
	}
	wg.Wait()
	fmt.Println("Contador (con mutex):", contador, "(siempre 1000)")
}

// --- Procesamiento paralelo con WaitGroup ---
func procesarParalelo(items []string) []string {
	resultados := make([]string, len(items))
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		go func(idx int, s string) {
			defer wg.Done()
			// simulamos procesamiento costoso
			time.Sleep(50 * time.Millisecond)
			resultados[idx] = fmt.Sprintf("[procesado] %s", s)
		}(i, item) // capturamos i y item como argumentos (evita closure bug)
	}

	wg.Wait()
	return resultados
}

func main() {
	// --- WaitGroup básico ---
	fmt.Println("=== WaitGroup básico ===")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go trabajador(i, &wg)
	}

	fmt.Println("Esperando a todos los trabajadores...")
	wg.Wait()
	fmt.Println("Todos terminaron")

	// --- Procesamiento paralelo ---
	fmt.Println("\n=== Procesamiento paralelo ===")
	items := []string{"archivo1.txt", "archivo2.txt", "archivo3.txt", "archivo4.txt"}

	inicio := time.Now()
	resultados := procesarParalelo(items)
	fmt.Printf("Tiempo: %v (vs ~%v en serie)\n", time.Since(inicio), time.Duration(len(items)*50)*time.Millisecond)
	for _, r := range resultados {
		fmt.Println(" ", r)
	}

	// --- Race condition y solución ---
	demoRaceCondition()
	demoMutex()

	// --- Goroutine anónima inline ---
	fmt.Println("\n=== Goroutine anónima ===")
	done := make(chan struct{}) // channel de señal (sin datos)
	go func() {
		fmt.Println("Ejecutando en otra goroutine")
		close(done) // señala que terminó
	}()
	<-done // espera la señal
	fmt.Println("Goroutine terminó")
}
