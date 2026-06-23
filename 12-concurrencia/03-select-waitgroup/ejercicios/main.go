package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Select y WaitGroup ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa una "carrera" entre múltiples URLs (simuladas):
	//
	//   type Resultado struct {
	//       URL      string
	//       Respuesta string
	//       Duracion  time.Duration
	//   }
	//
	//   func simularPeticion(url string) Resultado
	//     → duerme una duración aleatoria (50-300ms con rand.Intn)
	//     → retorna Resultado{URL: url, Respuesta: "200 OK", Duracion: duracion}
	//
	//   func carrera(urls []string) Resultado
	//     → lanza una goroutine por URL
	//     → cada goroutine envía su resultado a un channel
	//     → retorna el PRIMERO que llegue (usa select)
	//
	// Prueba con 5 URLs: "api.go.com", "api.rust.com", etc.
	// Imprime: "Ganó: <URL> en <duración>"
	//
	// Importa "math/rand" y "time".

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "throttle": limita cuántas goroutines corren a la vez.
	//
	//   func throttle(tareas []func() string, limite int) []string
	//     → ejecuta todas las tareas pero con máximo `limite` goroutines simultáneas
	//     → el channel del semáforo controla el acceso: make(chan struct{}, limite)
	//     → retorna los resultados en el orden que lleguen (no necesariamente el original)
	//
	// Patrón del semáforo:
	//   sem <- struct{}{} // adquiere un slot (bloquea si sem está lleno)
	//   defer func() { <-sem }() // libera el slot al terminar
	//
	// Prueba:
	//   tareas := make([]func() string, 10)
	//   for i := range tareas {
	//       i := i
	//       tareas[i] = func() string {
	//           time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	//           return fmt.Sprintf("tarea-%d completada", i)
	//       }
	//   }
	//   resultados := throttle(tareas, 3)  // máximo 3 goroutines a la vez
	//
	// Verifica que nunca hay más de 3 ejecutando simultáneamente
	// (agrega un contador atómico con sync/atomic para verificarlo).

	// TODO
}
