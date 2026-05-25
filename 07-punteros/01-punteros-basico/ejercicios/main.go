package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Punteros Básicos ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: intercambiar(a, b *int)
	//   → intercambia los valores de las dos variables apuntadas
	//   → sin retornar valores (modifica in-place vía punteros)
	//
	// Prueba:
	//   x, y := 10, 20
	//   intercambiar(&x, &y)
	//   // x debe ser 20, y debe ser 10
	//
	// Luego implementa: duplicarSlice(nums *[]int)
	//   → multiplica cada elemento del slice por 2 (in-place)
	//   → no retorna nada
	//
	// Prueba:
	//   nums := []int{1, 2, 3, 4, 5}
	//   duplicarSlice(&nums)
	//   // nums debe ser [2 4 6 8 10]
	//
	// Nota: para slices no necesitas el puntero para modificar elementos,
	// pero sí si quieres re-asignar el slice completo (append, reslice, etc.)

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "optional value" usando punteros:
	//
	//   type OptionalInt struct {
	//       valor *int  // nil = sin valor, non-nil = tiene valor
	//   }
	//
	//   (o *OptionalInt) Set(v int)
	//     → asigna el valor
	//
	//   (o *OptionalInt) Get() (int, bool)
	//     → retorna el valor y true si existe, o 0 y false si es nil
	//
	//   (o *OptionalInt) Clear()
	//     → elimina el valor (pone nil)
	//
	//   (o OptionalInt) String() string
	//     → "Some(valor)" si tiene valor, "None" si no
	//
	// Prueba todos los métodos y verifica los estados.
	// (Este patrón es similar a Optional<T> de Java o Maybe en Haskell)

	// TODO
}
