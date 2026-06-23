package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Mutex y Context ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa una caché concurrente con TTL (time-to-live):
	//
	//   type EntradaCache struct {
	//       Valor     string
	//       Expira    time.Time
	//   }
	//
	//   type CacheSegura struct {
	//       mu    sync.RWMutex
	//       datos map[string]EntradaCache
	//   }
	//
	//   func NewCacheSegura() *CacheSegura
	//   func (c *CacheSegura) Set(clave, valor string, ttl time.Duration)
	//   func (c *CacheSegura) Get(clave string) (string, bool)
	//     → retorna false si la clave no existe O si ya expiró
	//   func (c *CacheSegura) Limpiar()
	//     → elimina todas las entradas expiradas (necesita Lock completo)
	//
	// Prueba con múltiples goroutines leyendo y escribiendo simultáneamente:
	//   - 3 goroutines escritoras: escriben con TTL de 200ms
	//   - 5 goroutines lectoras: leen durante 500ms
	//   - Después de 300ms, llama a Limpiar() y verifica el tamaño
	//
	// Importa "sync" y "time".

	// TODO

	// EJERCICIO 2 — Intermedio (el más complejo del curso) ──────
	//
	// Implementa un worker pool con context y cancelación:
	//
	//   type Job struct { ID int; Payload string }
	//   type Result struct { JobID int; Output string; Err error }
	//
	//   func workerPoolConContext(
	//       ctx context.Context,
	//       jobs <-chan Job,
	//       numWorkers int,
	//   ) <-chan Result
	//     → crea numWorkers goroutines
	//     → cada worker lee de jobs y procesa
	//     → si el context se cancela, los workers deben detenerse limpiamente
	//     → retorna un channel de resultados
	//     → cierra el channel de resultados cuando todos los workers terminan
	//
	//   Procesar un job: simula trabajo con time.Sleep(50ms) y retorna
	//   Result{JobID: job.ID, Output: "procesado: " + job.Payload}
	//
	// Prueba:
	//   ctx, cancel := context.WithTimeout(context.Background(), 300ms)
	//   defer cancel()
	//
	//   jobs := make(chan Job, 20)
	//   for i := 0; i < 20; i++ {
	//       jobs <- Job{ID: i, Payload: fmt.Sprintf("data-%d", i)}
	//   }
	//   close(jobs)
	//
	//   results := workerPoolConContext(ctx, jobs, 4)
	//   for r := range results {
	//       if r.Err != nil { fmt.Println("Error:", r.Err) }
	//       else { fmt.Println("OK:", r.Output) }
	//   }
	//
	// Con timeout de 300ms y 20 jobs a 50ms c/u con 4 workers:
	// ¿Cuántos jobs se procesarán antes del timeout?

	// TODO
}
