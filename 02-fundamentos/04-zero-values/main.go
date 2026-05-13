package main

import "fmt"

func main() {
	// =========================================================
	// ZERO VALUES — Valores por defecto en Go
	// =========================================================
	// En JavaScript: var x; → x es undefined
	// En Go: cada tipo tiene un valor útil por defecto, nunca undefined.
	// Esto elimina toda una categoría de bugs de "undefined is not a function".

	// Tipos numéricos → 0
	var entero int
	var decimal float64
	var complejo complex64

	// bool → false
	var booleano bool

	// string → "" (cadena vacía, no null)
	var texto string

	// Punteros, slices, maps, funciones, canales, interfaces → nil
	var puntero *int
	var slice []int
	var mapa map[string]int

	fmt.Println("=== Zero Values ===")
	fmt.Printf("int:        %d\n", entero)
	fmt.Printf("float64:    %f\n", decimal)
	fmt.Printf("complex64:  %v\n", complejo)
	fmt.Printf("bool:       %t\n", booleano)
	fmt.Printf("string:     %q\n", texto) // %q muestra las comillas
	fmt.Printf("*int:       %v\n", puntero)
	fmt.Printf("[]int:      %v\n", slice)
	fmt.Printf("map:        %v\n", mapa)

	// =========================================================
	// POR QUÉ ESTO IMPORTA
	// =========================================================

	// Puedes usar un int sin inicializar como contador:
	var contador int
	contador++
	contador++
	fmt.Println("\nContador:", contador) // 2, no undefined++

	// Puedes concatenar a un string vacío directamente:
	var resultado string
	resultado += "Hola"
	resultado += ", mundo"
	fmt.Println(resultado) // "Hola, mundo", no undefined+"Hola"

	// Un slice nil tiene len=0 y puedes hacer append directamente:
	var numeros []int
	numeros = append(numeros, 1, 2, 3)
	fmt.Println(numeros) // [1 2 3]

	// =========================================================
	// NIL no es lo mismo que zero value numérico
	// =========================================================
	// nil solo aplica a tipos de referencia: punteros, slices, maps,
	// funciones, channels, interfaces.
	// Intentar usar un mapa nil da panic en tiempo de ejecución:
	//   var m map[string]int
	//   m["clave"] = 1 // ✗ PANIC: assignment to entry in nil map

	// La solución: inicializar con make o literal
	mapaValido := make(map[string]int)
	mapaValido["clave"] = 1
	fmt.Println("\nMapa válido:", mapaValido)

	// =========================================================
	// ZERO VALUE DE STRUCTS
	// =========================================================
	type Punto struct {
		X, Y float64
	}

	var p Punto
	fmt.Printf("\nPunto zero: %+v\n", p) // {X:0 Y:0}

	// Cada campo del struct tiene su propio zero value
	type Persona struct {
		Nombre string
		Edad   int
		Activo bool
	}

	var persona Persona
	fmt.Printf("Persona zero: %+v\n", persona) // {Nombre: Edad:0 Activo:false}
}
