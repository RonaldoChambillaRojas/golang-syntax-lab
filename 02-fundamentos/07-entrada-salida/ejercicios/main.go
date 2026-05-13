package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Ejercicios: Entrada y Salida ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Sin leer input del usuario, imprime la siguiente tabla
	// usando fmt.Printf y los verbos apropiados:
	//
	// ┌─────────────┬───────┬──────────┐
	// │ Producto    │ Cant. │ Precio   │
	// ├─────────────┼───────┼──────────┤
	// │ Laptop      │     1 │ $1200.00 │
	// │ Mouse       │     2 │   $25.50 │
	// │ Teclado     │     1 │   $89.99 │
	// └─────────────┴───────┴──────────┘
	//
	// Pista: usa %-12s para alinear texto a la izquierda con ancho 12,
	// y %8.2f para alinear números a la derecha con 2 decimales.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Calculadora de IMC interactiva:
	// 1. Pide al usuario su nombre con fmt.Print y lee con bufio.Scanner
	// 2. Pide su peso en kg (como string, luego convierte con strconv.ParseFloat)
	// 3. Pide su altura en metros (igual)
	// 4. Calcula IMC = peso / (altura * altura)
	// 5. Imprime: "Hola <nombre>, tu IMC es: <valor con 2 decimales>"
	// 6. Clasifica con fmt.Printf:
	//    - IMC < 18.5 → "Bajo peso"
	//    - 18.5 ≤ IMC < 25 → "Normal"
	//    - 25 ≤ IMC < 30 → "Sobrepeso"
	//    - IMC ≥ 30 → "Obesidad"
	//
	// Para clasificar puedes usar if/else (ya lo puedes hacer intuitivamente).

	_ = bufio.NewScanner(os.Stdin) // elimina este _ cuando lo uses
	_ = strconv.ParseFloat
	_ = strings.TrimSpace

	// TODO
}
