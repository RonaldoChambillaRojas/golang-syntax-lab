package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Constantes ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Declara una constante Pi con el valor 3.14159.
	// Declara una variable radio con valor 7.0.
	// Calcula el área del círculo: Pi * radio * radio
	// Calcula la circunferencia: 2 * Pi * radio
	// Imprime ambos resultados con un mensaje descriptivo.

	const Pi = 3.14159
	var radio = 7.0

	area := Pi * radio * radio
	circunferencia := 2 * Pi * radio

	fmt.Printf("Área del círculo con radio %.2f: %.5f\n", radio, area)
	fmt.Printf("Circunferencia del círculo con radio %.2f: %.5f\n", radio, circunferencia)

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	// Definimos un tipo Talla basado en int y usamos iota para generar
	// constantes consecutivas. Como queremos que XS valga 1, usamos iota+1.
	type Talla int

	const (
		XS Talla = iota + 1
		S
		M
		L
		XL
	)

	talla := M

	fmt.Println("\nValores de Talla usando iota:") // mostramos los valores calculados
	fmt.Println("XS =", XS)
	fmt.Println("S  =", S)
	fmt.Println("M  =", M)
	fmt.Println("L  =", L)
	fmt.Println("XL =", XL)

	switch talla {
	case XS:
		fmt.Println("Talla seleccionada: Extra small")
	case S:
		fmt.Println("Talla seleccionada: Small")
	case M:
		fmt.Println("Talla seleccionada: Medium")
	case L:
		fmt.Println("Talla seleccionada: Large")
	case XL:
		fmt.Println("Talla seleccionada: Extra large")
	}
}
