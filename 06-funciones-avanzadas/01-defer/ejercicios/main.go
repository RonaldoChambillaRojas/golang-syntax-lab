package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Defer ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa una función medirTiempo(nombre string) func()
	// que simule medir el tiempo de ejecución de una operación:
	//
	//   func medirTiempo(nombre string) func() {
	//       inicio := time.Now()  // registra cuándo empezó
	//       return func() {
	//           duracion := time.Since(inicio)
	//           fmt.Printf("%s tardó: %v\n", nombre, duracion)
	//       }
	//   }
	//
	// Úsala con defer en otra función:
	//
	//   func operacionLarga() {
	//       defer medirTiempo("operacionLarga")()  // nota el () al final
	//       // simula trabajo con time.Sleep(100 * time.Millisecond)
	//   }
	//
	// Importa "time" si lo necesitas.
	// Llama a operacionLarga() desde main y verifica el tiempo impreso.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "stack de llamadas" manual con defer:
	//
	// Crea una función trazar(nombre string) func() que:
	//   - Al ser llamada, imprime: "→ Entrando a: <nombre>"
	//   - Retorna una función que imprime: "← Saliendo de: <nombre>"
	//
	// Úsala con defer en tres funciones anidadas:
	//
	//   func nivel1() {
	//       defer trazar("nivel1")()
	//       nivel2()
	//   }
	//
	//   func nivel2() {
	//       defer trazar("nivel2")()
	//       nivel3()
	//   }
	//
	//   func nivel3() {
	//       defer trazar("nivel3")()
	//       fmt.Println("  [ejecutando nivel3]")
	//   }
	//
	// Llama a nivel1() y observa el orden de entrada/salida.
	// Debe verse como un call stack real.

	// TODO
}
