package main

import (
	"fmt"
	"strings"
)

// =========================================================
// PIPELINES CON GENERICS
// =========================================================
// Un pipeline transforma datos a través de una cadena de funciones.
// En JavaScript usarías: array.filter(...).map(...).reduce(...)
// Con generics en Go puedes crear lo mismo de forma type-safe.

// --- Tipos base ---
type Predicate[T any] func(T) bool
type Mapper[T, U any] func(T) U
type Reducer[T, Acc any] func(Acc, T) Acc

// --- Funciones genéricas del pipeline ---

func Map[T, U any](slice []T, fn Mapper[T, U]) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, pred Predicate[T]) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, Acc any](slice []T, inicial Acc, fn Reducer[T, Acc]) Acc {
	acc := inicial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

// --- Tomar / saltar elementos ---
func Take[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

func Drop[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}

// --- Agrupar por clave ---
func GroupBy[T any, K comparable](slice []T, keyFn func(T) K) map[K][]T {
	grupos := make(map[K][]T)
	for _, v := range slice {
		k := keyFn(v)
		grupos[k] = append(grupos[k], v)
	}
	return grupos
}

// --- Flatten ---
func Flatten[T any](slices [][]T) []T {
	var resultado []T
	for _, s := range slices {
		resultado = append(resultado, s...)
	}
	return resultado
}

// --- Zip: combina dos slices en uno de pares ---
type Par[A, B any] struct {
	A A
	B B
}

func Zip[A, B any](as []A, bs []B) []Par[A, B] {
	n := len(as)
	if len(bs) < n {
		n = len(bs)
	}
	resultado := make([]Par[A, B], n)
	for i := 0; i < n; i++ {
		resultado[i] = Par[A, B]{as[i], bs[i]}
	}
	return resultado
}

// --- Constraints personalizadas ---
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](slice []T) T {
	return Reduce(slice, *new(T), func(acc, v T) T { return acc + v })
}

func main() {
	// --- Pipeline de enteros ---
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original:", numeros)

	// Equivalente en JS: numeros.filter(n => n%2==0).map(n => n*n)
	resultado := Map(
		Filter(numeros, func(n int) bool { return n%2 == 0 }),
		func(n int) int { return n * n },
	)
	fmt.Println("Pares al cuadrado:", resultado)

	suma := Sum(resultado)
	fmt.Println("Suma:", suma)

	// --- Pipeline de strings ---
	palabras := []string{"  Go  ", "rust", "PYTHON", "javascript", " TypeScript "}
	limpias := Map(palabras, strings.TrimSpace)
	minusculas := Map(limpias, strings.ToLower)
	largas := Filter(minusculas, func(s string) bool { return len(s) > 4 })
	fmt.Println("\nPipeline strings:", largas)

	// --- GroupBy ---
	type Persona struct {
		Nombre string
		Depto  string
	}

	personas := []Persona{
		{"Ana", "Ingeniería"},
		{"Carlos", "Ventas"},
		{"María", "Ingeniería"},
		{"Luis", "Ventas"},
		{"Elena", "Ingeniería"},
	}

	porDepto := GroupBy(personas, func(p Persona) string {
		return p.Depto
	})
	fmt.Println("\nPor departamento:")
	for depto, equipo := range porDepto {
		nombres := Map(equipo, func(p Persona) string {
			return p.Nombre
		})
		fmt.Printf("  %s: %s\n", depto, strings.Join(nombres, ", "))
	}

	// --- Zip ---
	claves := []string{"nombre", "edad", "ciudad"}
	valores := []any{"Ana", 28, "CDMX"}
	pares := Zip(claves, valores)
	fmt.Println("\nZip:")
	for _, p := range pares {
		fmt.Printf("  %s: %v\n", p.A, p.B)
	}

	// --- Flatten ---
	matriz := [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}
	flat := Flatten(matriz)
	fmt.Println("\nFlatten:", flat)
	fmt.Println("Sum of flat:", Sum(flat))
}
