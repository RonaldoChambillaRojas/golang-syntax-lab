package main

import "fmt"

func main() {
	// =========================================================
	// ARRAYS EN GO
	// =========================================================
	// Los arrays en Go son MUY diferentes a los de JavaScript:
	// - Tamaño FIJO determinado en tiempo de compilación
	// - El tipo incluye el tamaño: [3]int ≠ [4]int (son tipos distintos!)
	// - Son tipos VALOR: se copian en asignación (como números, no referencias)
	//
	// Por eso, en la práctica Go usa SLICES casi siempre.
	// Los arrays se usan cuando el tamaño es fijo y conocido (buffers, etc.)

	// --- Declaración ---
	var numeros [5]int // zero value: [0 0 0 0 0]
	fmt.Println("Zero:", numeros)

	// Literal de array
	frutas := [3]string{"manzana", "banana", "cereza"}
	fmt.Println("Frutas:", frutas)

	// El compilador cuenta los elementos con ...
	colores := [...]string{"rojo", "verde", "azul"} // tamaño = 3
	fmt.Println("Colores:", colores)
	fmt.Println("Tamaño:", len(colores)) // 3

	// --- Acceso por índice (igual que JS) ---
	fmt.Println("Primera fruta:", frutas[0])
	frutas[1] = "mango"
	fmt.Println("Modificado:", frutas)

	// --- Los arrays son VALOR, no referencia ---
	// En JavaScript: const b = a crea una REFERENCIA al mismo array
	// En Go: b := a crea una COPIA completamente independiente
	original := [3]int{1, 2, 3}
	copia := original  // copia completa
	copia[0] = 99      // no afecta a original
	fmt.Println("Original:", original) // [1 2 3]
	fmt.Println("Copia:", copia)       // [99 2 3]

	// --- Iterar con for range ---
	// En JS: frutas.forEach((f, i) => ...)
	// En Go:
	fmt.Println("\nIterando:")
	for indice, valor := range frutas {
		fmt.Printf("  [%d] = %s\n", indice, valor)
	}

	// Si solo necesitas el valor (ignoras el índice con _):
	for _, fruta := range frutas {
		fmt.Println(" -", fruta)
	}

	// --- Arrays multidimensionales ---
	// Como matrices en álgebra lineal
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\nMatriz:", matriz)
	fmt.Println("Elemento [1][2]:", matriz[1][2]) // 6

	// --- Comparación ---
	// Los arrays del mismo tipo Y tamaño pueden compararse con ==
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	fmt.Println("\na == b:", a == b) // true
	fmt.Println("a == c:", a == c)  // false
	// En JS: [] == [] es false porque compara referencias
}
