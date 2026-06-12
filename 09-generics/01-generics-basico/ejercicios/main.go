package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Generics Básico ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un tipo genérico Par[A, B any]:
	//   type Par[A, B any] struct {
	//       Primero A
	//       Segundo B
	//   }
	//   func NuevoPar[A, B any](a A, b B) Par[A, B]
	//   func (p Par[A, B]) String() string → "(Primero, Segundo)"
	//   func (p Par[A, B]) Invertir() Par[B, A]  → invierte el orden
	//
	// Prueba:
	//   p1 := NuevoPar("Go", 42)           → Par{Primero:"Go", Segundo:42}
	//   p2 := NuevoPar(3.14, true)         → Par{Primero:3.14, Segundo:true}
	//   p3 := p1.Invertir()                → Par{Primero:42, Segundo:"Go"}
	//   fmt.Println(p1, p2, p3.Primero)

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un tipo genérico Resultado[T any]:
	//   type Resultado[T any] struct { ... }  // puede ser Ok o Err
	//
	//   func Ok[T any](valor T) Resultado[T]
	//   func Err[T any](err error) Resultado[T]
	//   func (r Resultado[T]) EsOk() bool
	//   func (r Resultado[T]) Valor() (T, error)
	//     → retorna (valor, nil) si Ok
	//     → retorna (zero, error) si Err
	//   func (r Resultado[T]) Mapear(fn func(T) T) Resultado[T]
	//     → si Ok, aplica fn al valor y retorna Ok(fn(valor))
	//     → si Err, retorna el mismo Err sin llamar fn
	//
	// Este patrón es similar a Result<T, E> de Rust o Either de Haskell.
	//
	// Prueba:
	//   r1 := Ok(42)
	//   r2 := r1.Mapear(func(n int) int { return n * 2 })
	//   valor, err := r2.Valor()  → 84, nil
	//
	//   r3 := Err[int](errors.New("algo salió mal"))
	//   r4 := r3.Mapear(func(n int) int { return n * 2 })
	//   _, err = r4.Valor()  → 0, error

	// TODO
}
