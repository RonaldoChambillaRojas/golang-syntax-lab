package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Goroutines ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa una función sumaParalela(nums []int, trabajadores int) int
	//   → divide el slice en partes iguales (una por trabajador)
	//   → cada goroutine suma su parte
	//   → retorna la suma total
	//
	// Usa sync.WaitGroup y un slice de resultados parciales para recolectarlos.
	// (No uses channels todavía, los veremos después)
	//
	// Prueba:
	//   nums := make([]int, 1000)
	//   for i := range nums { nums[i] = i + 1 }  // 1 a 1000
	//   suma := sumaParalela(nums, 4)
	//   fmt.Println(suma)  // debe ser 500500
	//
	// Compara el tiempo con suma secuencial simple.
	// (Para números tan pequeños la paralela puede ser más lenta por overhead,
	// pero el patrón es lo importante)

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "pool de trabajadores" (worker pool):
	//
	//   type Trabajo struct {
	//       ID    int
	//       Valor int
	//   }
	//
	//   type Resultado struct {
	//       TrabajoID int
	//       Cuadrado  int
	//   }
	//
	//   func workerPool(trabajos []Trabajo, numWorkers int) []Resultado
	//     → crea numWorkers goroutines
	//     → cada goroutine toma trabajos de la lista (divide equitativamente)
	//     → calcula el cuadrado de Trabajo.Valor
	//     → retorna todos los resultados
	//
	// Prueba con 20 trabajos y 4 workers.
	// Imprime qué worker procesó qué trabajo (agrega un campo WorkerID al Resultado).
	//
	// Nota: para dividir los trabajos equitativamente entre workers,
	// puedes asignar: workerIdx = trabajoIdx % numWorkers

	// TODO
}
