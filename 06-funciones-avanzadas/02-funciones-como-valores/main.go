package main

import (
	"fmt"
	"strings"
)

// =========================================================
// FUNCIONES COMO VALORES — Higher-order functions
// =========================================================
// Go trata las funciones como cualquier otro valor:
// puedes asignarlas a variables, pasarlas como parámetros,
// retornarlas desde otras funciones y almacenarlas en estructuras.
// Es idéntico al comportamiento de JavaScript en este aspecto.

// --- Tipo función: se puede definir con type ---
// Equivalente a TypeScript: type Predicado = (n: number) => boolean
type PredicadoInt func(int) bool
type TransformadorInt func(int) int

// --- Funciones de orden superior ---
// Como filter(), map(), reduce() de JavaScript

// filtrar es como Array.prototype.filter de JS
func filtrar(nums []int, pred PredicadoInt) []int {
	resultado := make([]int, 0)
	for _, n := range nums {
		if pred(n) {
			resultado = append(resultado, n)
		}
	}
	return resultado
}

// transformar es como Array.prototype.map de JS
func transformar(nums []int, fn TransformadorInt) []int {
	resultado := make([]int, len(nums))
	for i, n := range nums {
		resultado[i] = fn(n)
	}
	return resultado
}

// reducir es como Array.prototype.reduce de JS
func reducir(nums []int, inicial int, fn func(int, int) int) int {
	acc := inicial
	for _, n := range nums {
		acc = fn(acc, n)
	}
	return acc
}

// componer combina funciones: componer(f, g)(x) = g(f(x))
// En JS: const componer = (...fns) => x => fns.reduce((v, f) => f(v), x)
func componer(fns ...TransformadorInt) TransformadorInt {
	return func(x int) int {
		for _, fn := range fns {
			x = fn(x)
		}
		return x
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// --- Predicados como valores ---
	esPar := func(n int) bool { return n%2 == 0 }
	esImpar := func(n int) bool { return n%2 != 0 }
	esMayorA5 := func(n int) bool { return n > 5 }

	fmt.Println("Original:", nums)
	fmt.Println("Pares:", filtrar(nums, esPar))
	fmt.Println("Impares:", filtrar(nums, esImpar))
	fmt.Println("Mayores a 5:", filtrar(nums, esMayorA5))

	// Función inline como argumento (como arrow function de JS)
	fmt.Println("Múltiplos de 3:", filtrar(nums, func(n int) bool {
		return n%3 == 0
	}))

	// --- Transformaciones ---
	doble := func(n int) int { return n * 2 }
	cuadrado := func(n int) int { return n * n }

	fmt.Println("\nDobles:", transformar(nums, doble))
	fmt.Println("Cuadrados:", transformar(nums, cuadrado))

	// --- Reduce ---
	suma := reducir(nums, 0, func(acc, n int) int { return acc + n })
	producto := reducir(nums[:5], 1, func(acc, n int) int { return acc * n })
	fmt.Printf("\nSuma: %d, Producto(1-5): %d\n", suma, producto)

	// --- Composición ---
	duplicarYSumarUno := componer(doble, func(n int) int { return n + 1 })
	fmt.Println("\nComposición (doble + sumarUno):", transformar([]int{1, 2, 3}, duplicarYSumarUno))

	// --- Funciones como campos de struct ---
	// En JS: const handlers = { onClick: () => ..., onHover: () => ... }
	procesadores := map[string]func(string) string{
		"mayusculas": strings.ToUpper,
		"minusculas": strings.ToLower,
		"titulo":     strings.Title, //nolint — simplificado para el ejemplo
		"reverso": func(s string) string {
			runes := []rune(s)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			return string(runes)
		},
	}

	texto := "hola mundo"
	for nombre, fn := range procesadores {
		fmt.Printf("%-12s → %s\n", nombre, fn(texto))
	}
}
