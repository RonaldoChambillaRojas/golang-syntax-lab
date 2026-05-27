package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Variables ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Declara las siguientes variables usando la forma más idiomática
	// para código dentro de una función:
	//   - tu nombre (string)
	//   - tu edad (int)
	//   - si ya programabas antes de este curso (bool)
	//
	// Imprime los tres valores con un solo fmt.Println.

	nombre := "Ronaldo"
	edad := 20
	altura := 1.68

	fmt.Println(nombre, edad, altura)

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Tienes: precio := 250.0 y descuento := 15.0 (es un porcentaje)
	//
	// 1. Calcula precioFinal = precio - (precio * descuento / 100)
	// 2. Imprime: "Precio final: <valor>"
	// 3. Usando asignación múltiple (sin variable temporal),
	//    intercambia precio y precioFinal.
	// 4. Imprime ambos valores después del intercambio.

	precio := 250.0
	descuento := 15.0

	precio_final := precio - (precio * descuento / 100)

	fmt.Println("Precio final: ", precio_final)

	precio_final, precio = precio, precio_final

	fmt.Println(precio, precio_final)
}
