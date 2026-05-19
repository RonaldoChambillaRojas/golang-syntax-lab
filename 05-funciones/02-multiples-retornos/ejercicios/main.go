package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Múltiples Retornos ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: euclides(a, b int) (cociente int, residuo int)
	//   Retorna el cociente y el residuo de dividir a entre b.
	//   Usa retornos con nombre (named returns).
	//
	// Pruébala:
	//   euclides(17, 5) → cociente=3, residuo=2
	//   euclides(100, 7) → cociente=14, residuo=2
	//
	// Imprime: "17 ÷ 5 = 3 con residuo 2"

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa: parsearFecha(s string) (dia, mes, anio int, err error)
	//   El string tiene formato "DD/MM/AAAA" (ej: "15/03/1990")
	//   Debe:
	//   1. Separar por "/" usando strings.Split(s, "/")
	//   2. Verificar que tenga exactamente 3 partes
	//      Si no: return 0, 0, 0, errors.New("formato inválido")
	//   3. Convertir cada parte a int con strconv.Atoi
	//      Si falla alguna: retornar el error
	//   4. Validar que día sea 1-31, mes 1-12, año > 0
	//      Si no: return 0, 0, 0, errors.New("valores fuera de rango")
	//   5. Si todo está bien: retornar dia, mes, anio, nil
	//
	// Pruébala con: "15/03/1990", "32/01/2024", "abc/03/1990", "15/13/2024"
	//
	// Importa "strings", "strconv", "errors" si los necesitas.

	// TODO
}

// Define aquí las funciones: euclides, parsearFecha
