package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: If/Else ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un clasificador de IMC:
	//   imc := 24.5
	//
	// Clasifica con if/else según los rangos:
	//   < 18.5       → "Bajo peso"
	//   18.5 - 24.9  → "Normal"
	//   25.0 - 29.9  → "Sobrepeso"
	//   ≥ 30.0       → "Obesidad"
	//
	// Imprime: "IMC: 24.5 → Categoría: Normal"

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Clasificador de año bisiesto usando la sentencia de init del if.
	//
	// Un año es bisiesto si:
	//   - Es divisible entre 4 Y
	//   - No es divisible entre 100, O SÍ es divisible entre 400
	//   Fórmula: (año%4==0 && año%100!=0) || año%400==0
	//
	// Declara: anios := []int{2000, 1900, 2024, 2023, 2100, 2400}
	// Para cada año, usando la sentencia de init del if:
	//   if esBisiesto := (condicion); esBisiesto {
	//       ...
	//   }
	// Imprime: "2024: bisiesto" o "2023: no bisiesto"

	// TODO
}
