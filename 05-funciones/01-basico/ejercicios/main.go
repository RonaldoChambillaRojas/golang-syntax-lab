package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Funciones Básicas ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa las siguientes funciones (defínelas fuera de main):
	//
	//   esPositivo(n int) bool
	//     → retorna true si n > 0
	//
	//   valorAbsoluto(n float64) float64
	//     → retorna el valor absoluto sin usar math.Abs
	//     (usa if: si n < 0, retorna -n, sino retorna n)
	//
	//   potencia(base, exponente int) int
	//     → calcula base^exponente con un bucle (sin math.Pow)
	//     → asume exponente >= 0
	//
	// Pruébalas en main con varios valores e imprime los resultados.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa una función aplicarATodos:
	//
	//   aplicarATodos(numeros []int, f func(int) int) []int
	//     → aplica la función f a cada elemento del slice
	//     → retorna un NUEVO slice con los resultados
	//     (es decir, implementa map() de JavaScript en Go)
	//
	// Pruébala con:
	//   numeros := []int{1, 2, 3, 4, 5}
	//
	// 1. Aplica una función que doble cada número.
	// 2. Aplica una función que eleve al cuadrado.
	// 3. Aplica valorAbsoluto (adaptada a int) a: []int{-3, 5, -2, 8, -1}
	//
	// Imprime cada resultado.

	// TODO
}

// Define aquí las funciones: esPositivo, valorAbsoluto, potencia, aplicarATodos
