package main

import (
	"fmt"
	"maps"
)

func main() {
	// =========================================================
	// MAPS EN GO
	// =========================================================
	// Maps = diccionarios clave-valor.
	// En JS sería: const obj = {}  o  new Map()
	// Diferencias importantes:
	// - Las claves deben ser de un tipo comparable (==)
	// - Acceder a una clave inexistente retorna el ZERO VALUE (no undefined)
	// - El orden de iteración es ALEATORIO (intencional, por diseño)
	// - Un map nil produce PANIC en escritura (inicializa con make o literal)

	// --- Creación ---
	// Con make (útil cuando no conoces los valores iniciales)
	edades := make(map[string]int)
	edades["Ana"] = 28
	edades["Carlos"] = 35
	edades["María"] = 22

	// Literal (más conciso cuando conoces los valores)
	capitales := map[string]string{
		"México":  "Ciudad de México",
		"España":  "Madrid",
		"Francia": "París",
	}

	fmt.Println("Edades:", edades)
	fmt.Println("Capitales:", capitales)

	// --- Acceso a valores ---
	fmt.Println("\nEdad de Ana:", edades["Ana"]) // 28

	// Acceder a una clave que NO existe → zero value del tipo valor (no error!)
	fmt.Println("Edad de Pedro:", edades["Pedro"]) // 0 (zero value de int)

	// =========================================================
	// COMMA OK IDIOM — distinguir "clave no existe" de "valor es zero"
	// =========================================================
	// En JS: obj.clave === undefined (distingues con ===)
	// En Go: usas la forma de dos valores del acceso al map

	valor, existe := edades["Ana"]
	fmt.Printf("\n'Ana' existe: %t, valor: %d\n", existe, valor) // true, 28

	valor, existe = edades["Pedro"]
	fmt.Printf("'Pedro' existe: %t, valor: %d\n", existe, valor) // false, 0

	// =========================================================
	// ELIMINAR ELEMENTOS
	// =========================================================
	// En JS: delete obj.clave
	// En Go: delete(mapa, clave)

	delete(edades, "Carlos")
	fmt.Println("\nDespués de delete:", edades)

	// Eliminar una clave que no existe es seguro (no produce error)
	delete(edades, "clave_inexistente") // sin efecto

	// =========================================================
	// ITERACIÓN — el orden es ALEATORIO
	// =========================================================
	// Esto es por diseño: Go aleatoriza el orden para que los programas
	// no dependan accidentalmente de un orden específico.
	fmt.Println("\nIterando capitales:")
	for pais, capital := range capitales {
		fmt.Printf("  %s → %s\n", pais, capital)
	}

	// Solo las claves:
	fmt.Println("\nSolo claves:")
	for pais := range capitales {
		fmt.Println(" -", pais)
	}

	// =========================================================
	// MAPS COMO CONTADORES — patrón muy común
	// =========================================================
	// Equivalente a reducir un array a un objeto en JS:
	// words.reduce((acc, w) => ({ ...acc, [w]: (acc[w] || 0) + 1 }), {})

	palabras := []string{"go", "es", "bueno", "go", "es", "rápido", "go"}
	frecuencia := make(map[string]int)
	for _, p := range palabras {
		frecuencia[p]++ // si la clave no existe, zero value (0) + 1 = 1
	}
	fmt.Println("\nFrecuencia:", frecuencia)

	// =========================================================
	// COMPARAR MAPS (Go 1.21+)
	// =========================================================
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}

	fmt.Println("\nm1 == m2:", maps.Equal(m1, m2)) // true
	fmt.Println("m1 == m3:", maps.Equal(m1, m3)) // false

	// Nota: maps.Equal viene del paquete maps (Go 1.21+)
	// Antes, debías comparar manualmente con un loop
}
