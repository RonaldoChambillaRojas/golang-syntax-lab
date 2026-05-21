package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Closures ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: nuevoAcumulador() func(int) int
	//   - Cada vez que se llama con un número, lo suma al total acumulado.
	//   - Retorna el total actual después de cada suma.
	//
	// Ejemplo:
	//   acumular := nuevoAcumulador()
	//   acumular(10) → 10
	//   acumular(5)  → 15
	//   acumular(20) → 35
	//
	// Verifica que dos instancias de nuevoAcumulador() son independientes.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa: memoizar(f func(int) int) func(int) int
	//   - Retorna una versión cacheada de f.
	//   - La primera vez que se llama con un valor n, calcula f(n) y lo cachea.
	//   - Las siguientes veces con el mismo n, retorna el cache sin recalcular.
	//   - Para demostrar que funciona, haz que f imprima "calculando <n>"
	//     solo cuando realmente calcula (no cuando usa el cache).
	//
	// Pruébalo con una función factorialLento(n int) int que calcule
	// el factorial de n de forma iterativa.
	//
	//   factorialCacheado := memoizar(factorialLento)
	//   factorialCacheado(5)  → imprime "calculando 5", retorna 120
	//   factorialCacheado(5)  → no imprime nada, retorna 120 (desde cache)
	//   factorialCacheado(3)  → imprime "calculando 3", retorna 6
	//
	// Nota: la clave de cache es el argumento n (tipo int).

	// TODO
}
