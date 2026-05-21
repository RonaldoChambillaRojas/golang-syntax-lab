package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Funciones como Valores ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa estas funciones de orden superior para strings:
	//
	//   filtrarStrings(lista []string, pred func(string) bool) []string
	//   transformarStrings(lista []string, fn func(string) string) []string
	//
	// Pruébalas con la lista:
	//   lenguajes := []string{"Go", "Rust", "Python", "JavaScript", "C", "TypeScript"}
	//
	// 1. Filtra los que tienen más de 3 caracteres (len(s) > 3)
	// 2. Filtra los que empiezan con "G" o "R" (strings.HasPrefix)
	// 3. Transforma todos a mayúsculas (strings.ToUpper)
	// 4. Transforma agregando " !" al final
	//
	// Imprime cada resultado.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "pipeline" de procesamiento de datos:
	//
	//   pipeline(datos []int, pasos ...func([]int) []int) []int
	//     → aplica cada paso en secuencia al resultado del anterior
	//     → retorna el resultado final
	//
	// Crea los siguientes pasos:
	//   soloPositivos    → filtra valores <= 0
	//   duplicar         → multiplica cada elemento por 2
	//   ordenarDescend   → ordena de mayor a menor (usa slices.Sort + slices.Reverse)
	//   limitarA(n int) func([]int) []int → retorna solo los primeros n elementos
	//
	// Pruébalo:
	//   datos := []int{-3, 5, -1, 8, 2, -7, 10, 3, 6}
	//   resultado := pipeline(datos, soloPositivos, duplicar, ordenarDescend, limitarA(3))
	//   Esperado: [20 16 12]
	//
	// Importa "slices" si lo necesitas.

	// TODO
}
