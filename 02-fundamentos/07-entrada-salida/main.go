package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// =========================================================
	// SALIDA: fmt.Print, fmt.Println, fmt.Printf
	// =========================================================
	// fmt = "format" — el paquete de I/O estándar de Go
	// Equivalente a console.log en JavaScript

	fmt.Print("Sin salto de línea")   // no agrega \n
	fmt.Println()                     // solo salto de línea
	fmt.Println("Con salto de línea") // agrega \n automáticamente
	fmt.Println("varios", "valores", 42, true) // separa con espacio

	// fmt.Printf: formato similar a printf de C (y a template literals de JS)
	nombre := "Ronaldo"
	edad := 25
	altura := 1.75

	fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2f\n", nombre, edad, altura)

	// =========================================================
	// VERBOS DE FORMATO — los más usados
	// =========================================================
	fmt.Println("\n--- Verbos ---")
	fmt.Printf("%v\n", 42)        // valor por defecto (siempre funciona)
	fmt.Printf("%d\n", 42)        // entero decimal
	fmt.Printf("%f\n", 3.14)      // float (6 decimales por defecto)
	fmt.Printf("%.2f\n", 3.14159) // float con 2 decimales
	fmt.Printf("%s\n", "texto")   // string
	fmt.Printf("%q\n", "texto")   // string con comillas
	fmt.Printf("%t\n", true)      // bool
	fmt.Printf("%T\n", 42)        // TIPO de la variable → int
	fmt.Printf("%T\n", "hola")    // → string
	fmt.Printf("%b\n", 42)        // binario → 101010
	fmt.Printf("%x\n", 255)       // hexadecimal → ff
	fmt.Printf("%p\n", &nombre)   // dirección de memoria (puntero)

	// %+v en structs muestra también los nombres de los campos
	type Punto struct{ X, Y int }
	p := Punto{3, 4}
	fmt.Printf("%v\n", p)  // {3 4}
	fmt.Printf("%+v\n", p) // {X:3 Y:4}
	fmt.Printf("%#v\n", p) // main.Punto{X:3, Y:4} (Go syntax completa)

	// =========================================================
	// fmt.Sprintf — construir strings con formato (sin imprimir)
	// =========================================================
	// En JS usarías template literals: `Hola, ${nombre}!`
	saludo := fmt.Sprintf("Hola, %s! Tienes %d años.", nombre, edad)
	fmt.Println(saludo)

	// =========================================================
	// ENTRADA: fmt.Scan y fmt.Scanln
	// =========================================================
	// Comenta esta sección si no quieres que el programa espere input

	fmt.Print("\nEscribe tu nombre: ")

	// Opción 1: fmt.Scan (lee hasta espacio/newline)
	// var input string
	// fmt.Scan(&input)

	// Opción 2: bufio.Scanner (lee la línea completa, más robusto)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		fmt.Printf("Bienvenido, %s!\n", input)
	}

	// =========================================================
	// fmt.Sscanf — leer con formato desde un string
	// =========================================================
	// Útil para parsear strings estructurados
	var dia, mes, anio int
	_, err := fmt.Sscanf("15/06/1995", "%d/%d/%d", &dia, &mes, &anio)
	if err == nil {
		fmt.Printf("Fecha: día=%d, mes=%d, año=%d\n", dia, mes, anio)
	}
}
