package main

import "fmt"

// Variable de nivel de paquete: visible en todo el archivo.
// Solo puede declararse con var, nunca con :=
// En JS sería un módulo-level const/let
var appNombre = "MiApp"

func main() {
	// =========================================================
	// VARIABLES EN GO
	// =========================================================

	// --- Forma 1: var con tipo explícito ---
	// Equivalente en TypeScript: let nombre: string = "Go"
	var nombre string = "Go"

	// --- Forma 2: var con inferencia de tipo ---
	// Go infiere el tipo a partir del valor. El tipo sigue siendo fijo.
	var version = 1.26 // Go infiere float64

	// --- Forma 3: := declaración corta (la más usada dentro de funciones) ---
	// Solo funciona DENTRO de funciones. Equivale a: var lenguaje string = "Go"
	lenguaje := "Golang"

	fmt.Println(nombre, version, lenguaje)

	// =========================================================
	// DIFERENCIA CRÍTICA con JavaScript
	// =========================================================
	// En JS: let x = 5; x = "texto" → válido (tipado dinámico)
	// En Go: el tipo se fija en la declaración y NO cambia
	//
	// contador := 0
	// contador = "diez" // ✗ ERROR: cannot use "diez" (type string) as type int

	contador := 0
	contador = 10 // ✓ reasignación del mismo tipo
	fmt.Println(contador)

	// =========================================================
	// MÚLTIPLES VARIABLES
	// =========================================================

	// Asignación múltiple en una línea
	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	// Bloque var: útil para agrupar variables de configuración relacionadas
	var (
		host string = "localhost"
		port int    = 8080
		ssl  bool   = false
	)
	fmt.Println(host, port, ssl)

	// =========================================================
	// SWAP (intercambio de valores)
	// =========================================================
	// En JS necesitas una temporal o destructuring: [a, b] = [b, a]
	// En Go la asignación múltiple hace lo mismo nativamente:
	a, b := "primero", "segundo"
	a, b = b, a
	fmt.Println(a, b) // segundo primero

	// =========================================================
	// REGLA: Variables declaradas DEBEN usarse
	// =========================================================
	// En JS esto es un warning o nada. En Go es un ERROR de compilación.
	// Esto obliga a mantener el código limpio sin variables zombie.
	//
	// Si necesitas ignorar un valor explícitamente, usa _ (blank identifier):
	resultado, _ := dividir(10, 3) // ignoras el residuo intencionalmente
	fmt.Println(resultado)

	fmt.Println(appNombre)
}

// dividir retorna cociente y residuo — veremos múltiples retornos en la sec. 5
func dividir(a, b int) (int, int) {
	return a / b, a % b
}
