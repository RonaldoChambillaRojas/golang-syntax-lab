package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Ejercicios: Conversiones de Tipo ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Tienes:  temperatura := 98  (grados Fahrenheit, tipo int)
	//
	// 1. Conviértela a float64 para hacer el cálculo con precisión.
	// 2. Aplica la fórmula: Celsius = (F - 32) * 5 / 9
	// 3. Imprime el resultado como float64.
	// 4. También imprime la temperatura en Celsius truncada a int.
	//
	// Nota: el paso 1 es obligatorio porque la división entera 5/9 = 0 en Go.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Tienes una lista de precios como strings (así llegan de una API):
	precios := []string{"19.99", "5.50", "102.75", "no-es-numero", "8.00"}
	//
	// 1. Recorre el slice con range.
	// 2. Intenta convertir cada string a float64 con strconv.ParseFloat.
	// 3. Si la conversión falla, imprime: "Precio inválido: <valor>"
	// 4. Si tiene éxito, acumula la suma en una variable total.
	// 5. Al final imprime: "Total válido: <suma>"
	//
	// Pista: for i, precio := range precios { ... }
	// No uses el índice i directamente si no lo necesitas (usa _).

	_ = precios // elimina este _ cuando uses el slice
	_ = strconv.ParseFloat

	// TODO
}
