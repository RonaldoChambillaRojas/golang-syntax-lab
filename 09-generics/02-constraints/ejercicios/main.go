package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Constraints ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un tipo Set[T comparable] (conjunto):
	//   type Set[T comparable] struct { elementos map[T]struct{} }
	//
	//   func NuevoSet[T comparable]() Set[T]
	//   func (s *Set[T]) Agregar(v T)
	//   func (s Set[T]) Contiene(v T) bool
	//   func (s *Set[T]) Eliminar(v T)
	//   func (s Set[T]) Tamaño() int
	//   func (s Set[T]) ToSlice() []T       → convierte el set a un slice
	//   func (s Set[T]) Union(otro Set[T]) Set[T]  → unión de dos sets
	//   func (s Set[T]) Interseccion(otro Set[T]) Set[T] → intersección
	//
	// Nota: map[T]struct{} es la forma idiomática de representar un set en Go
	// (struct{} ocupa 0 bytes, solo importan las claves)
	//
	// Prueba con int y string.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa funciones genéricas para estadísticas:
	//
	//   type Estadisticas[T Numerico] struct {
	//       Min, Max, Suma T
	//       Promedio       float64
	//       Conteo         int
	//   }
	//
	//   type Numerico interface {
	//       ~int | ~int64 | ~float32 | ~float64
	//   }
	//
	//   func calcularStats[T Numerico](nums []T) (Estadisticas[T], error)
	//     → retorna error si el slice está vacío
	//     → calcula Min, Max, Suma, Promedio, Conteo
	//
	//   func (e Estadisticas[T]) String() string
	//     → imprime las estadísticas de forma legible
	//
	// Prueba con []int{3, 1, 4, 1, 5, 9, 2, 6} y []float64{1.5, 2.3, 0.8, 3.7}

	// TODO
}
