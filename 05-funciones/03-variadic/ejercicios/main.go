package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Funciones Variádicas ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: promedio(nums ...float64) (float64, bool)
	//   - Retorna el promedio de los números y true si hay elementos.
	//   - Si no hay elementos, retorna 0, false.
	//
	// Implementa: maximo(nums ...int) (int, bool)
	//   - Retorna el máximo y true si hay elementos.
	//   - Si no hay elementos, retorna 0, false.
	//
	// Pruébalas con:
	//   promedio(3.0, 5.5, 2.0, 8.5)  → 4.75, true
	//   promedio()                     → 0.0, false
	//   maximo(3, 1, 8, 2, 9, 4)      → 9, true

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un logger simple con niveles:
	//
	//   type NivelLog int
	//   const (
	//       INFO NivelLog = iota
	//       WARN
	//       ERROR
	//   )
	//
	//   log(nivel NivelLog, formato string, args ...interface{})
	//     Imprime con el formato: "[INFO] mensaje" / "[WARN] ..." / "[ERROR] ..."
	//     Usa fmt.Sprintf(formato, args...) para formatear el mensaje.
	//
	// Pruébalo:
	//   log(INFO, "Servidor iniciado en puerto %d", 8080)
	//   log(WARN, "Conexión lenta: %dms", 350)
	//   log(ERROR, "No se pudo conectar a %s", "db.host.com")
	//
	// Bonus: agrega un nivel mínimo global y no imprimas si el nivel es menor.
	//   nivelMinimo := WARN → no imprime INFO, sí WARN y ERROR

	// TODO
}

// Define aquí las funciones: promedio, maximo, log (y el tipo NivelLog)
