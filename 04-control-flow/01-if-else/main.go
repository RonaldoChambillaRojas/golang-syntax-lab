package main

import (
	"fmt"
	"strconv"
)

func main() {
	// =========================================================
	// IF / ELSE
	// =========================================================
	// Igual que JS, pero sin paréntesis obligatorios alrededor
	// de la condición (aunque tampoco son error si los pones).
	// Las llaves {} SÍ son obligatorias siempre.

	temperatura := 22.5

	if temperatura > 30 {
		fmt.Println("Hace calor")
	} else if temperatura > 20 {
		fmt.Println("Temperatura agradable")
	} else if temperatura > 10 {
		fmt.Println("Algo frío")
	} else {
		fmt.Println("Hace frío")
	}

	// =========================================================
	// IF CON SENTENCIA DE INICIALIZACIÓN — muy idiomático en Go
	// =========================================================
	// La variable declarada en la sentencia de init (antes del ;)
	// solo existe dentro del bloque if/else.
	// En JS harías: const result = operacion(); if (result.err) {...}
	// pero la variable quedaría en el scope externo.

	// Ejemplo clásico: verificar conversión de string a número
	if num, err := strconv.Atoi("42"); err != nil {
		fmt.Println("Error de conversión:", err)
	} else {
		fmt.Println("Número convertido:", num) // num solo existe aquí
	}
	// fmt.Println(num) // ✗ ERROR: num is undefined here

	// =========================================================
	// VARIABLE DE SOMBRA (Shadow Variable)
	// =========================================================
	// Una variable declarada dentro de un bloque "sombrea" a la exterior.
	// Es un bug silencioso muy común en Go.

	resultado := "exterior"
	if true {
		resultado := "interior" // NUEVA variable, sombrea a la exterior
		fmt.Println("Dentro:", resultado) // interior
		_ = resultado
	}
	fmt.Println("Fuera:", resultado) // exterior — NO fue modificada

	// Para modificar la variable exterior, no uses :=
	if true {
		resultado = "modificado" // asignación, no nueva declaración
	}
	fmt.Println("Modificado:", resultado) // modificado

	// =========================================================
	// EJEMPLO PRÁCTICO: Clasificador de notas
	// =========================================================
	nota := 87

	var calificacion string
	if nota >= 90 {
		calificacion = "A"
	} else if nota >= 80 {
		calificacion = "B"
	} else if nota >= 70 {
		calificacion = "C"
	} else if nota >= 60 {
		calificacion = "D"
	} else {
		calificacion = "F"
	}
	fmt.Printf("\nNota: %d → Calificación: %s\n", nota, calificacion)

	// =========================================================
	// NO HAY TERNARIO — alternativa idiomática
	// =========================================================
	// En JS: const etiqueta = activo ? "Activo" : "Inactivo"
	// En Go:
	activo := true
	etiqueta := "Inactivo"
	if activo {
		etiqueta = "Activo"
	}
	fmt.Println("Estado:", etiqueta)

	// Alternativa: función helper (lo veremos en sec. 5)
	// No hagas funciones de 1 línea tipo ternary — en Go si necesitas
	// una sola condición, el if/else de 3 líneas ES la forma correcta.
}
