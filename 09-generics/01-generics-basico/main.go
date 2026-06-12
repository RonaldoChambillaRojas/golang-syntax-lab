package main

import "fmt"

// =========================================================
// GENERICS EN GO — Desde Go 1.18
// =========================================================
// Syntax: func Nombre[T constraint](params) retorno
// T es el "type parameter" — como <T> en TypeScript
// "any" es la constraint más permisiva: acepta cualquier tipo

// --- Función genérica simple ---
// Antes de generics necesitabas interface{} y type assertions.
// Ahora tienes type safety completo.

func primero[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T // zero value del tipo T
		return zero, false
	}
	return slice[0], true
}

func ultimo[T any](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[len(slice)-1], true
}

// --- Map genérico (como Array.map de JavaScript) ---
func mapSlice[T, U any](slice []T, fn func(T) U) []U {
	resultado := make([]U, len(slice))
	for i, v := range slice {
		resultado[i] = fn(v)
	}
	return resultado
}

// T y U pueden ser tipos diferentes: transforma []T en []U

// --- Filter genérico ---
func filter[T any](slice []T, pred func(T) bool) []T {
	resultado := make([]T, 0)
	for _, v := range slice {
		if pred(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

// --- Reduce genérico ---
func reduce[T, Acc any](slice []T, inicial Acc, fn func(Acc, T) Acc) Acc {
	acc := inicial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// --- Tipo genérico: Stack ---
// En TypeScript: class Stack<T> { private items: T[] = [] }
type Stack[T any] struct {
	elementos []T
}

func (s *Stack[T]) Push(v T) {
	s.elementos = append(s.elementos, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elementos) == 0 {
		var zero T
		return zero, false
	}
	ultimo := s.elementos[len(s.elementos)-1]
	s.elementos = s.elementos[:len(s.elementos)-1]
	return ultimo, true
}

func (s Stack[T]) Len() int {
	return len(s.elementos)
}

// --- Contains: requiere comparable ---
// "comparable" es una constraint especial: tipos que soportan == y !=
func contiene[T comparable](slice []T, objetivo T) bool {
	for _, v := range slice {
		if v == objetivo {
			return true
		}
	}
	return false
}

// --- Inferencia de tipos ---
// Go puede inferir T a partir de los argumentos — no necesitas especificarlo

func main() {
	// --- primero / ultimo ---
	ints := []int{10, 20, 30, 40}
	strs := []string{"Go", "Rust", "Python"}

	if v, ok := primero(ints); ok { // T inferido como int
		fmt.Println("Primero int:", v)
	}
	if v, ok := ultimo(strs); ok { // T inferido como string
		fmt.Println("Último string:", v)
	}
	_, ok := primero([]float64{})
	fmt.Println("Primero de slice vacío, ok:", ok)

	// --- mapSlice ---
	// int → string
	palabras := mapSlice(ints, func(n int) string {
		return fmt.Sprintf("num%d", n)
	})
	fmt.Println("\nMapSlice int→string:", palabras)

	// string → int (longitud)
	longitudes := mapSlice(strs, func(s string) int { return len(s) })
	fmt.Println("MapSlice string→int:", longitudes)

	// --- filter ---
	pares := filter(ints, func(n int) bool { return n%20 == 0 })
	fmt.Println("\nFiltrar múltiplos de 20:", pares)

	largos := filter(strs, func(s string) bool { return len(s) > 3 })
	fmt.Println("Filtrar strings largos:", largos)

	// --- reduce ---
	suma := reduce(ints, 0, func(acc, n int) int { return acc + n })
	fmt.Println("\nSuma:", suma)

	concatenar := reduce(strs, "", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + ", " + s
	})
	fmt.Println("Concatenar:", concatenar)

	// --- Stack genérico ---
	var stackInt Stack[int]
	stackInt.Push(1)
	stackInt.Push(2)
	stackInt.Push(3)
	fmt.Printf("\nStack tiene %d elementos\n", stackInt.Len())
	if v, ok := stackInt.Pop(); ok {
		fmt.Println("Pop:", v) // 3
	}

	var stackStr Stack[string]
	stackStr.Push("Go")
	stackStr.Push("es")
	stackStr.Push("genial")
	if v, ok := stackStr.Pop(); ok {
		fmt.Println("Pop string:", v) // genial
	}

	// --- contiene ---
	fmt.Println("\nContiene 20:", contiene(ints, 20))
	fmt.Println("Contiene 99:", contiene(ints, 99))
	fmt.Println("Contiene 'Go':", contiene(strs, "Go"))
}
