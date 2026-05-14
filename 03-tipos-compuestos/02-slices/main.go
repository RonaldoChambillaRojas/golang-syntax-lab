package main

import (
	"fmt"
	"slices"
)

func main() {
	// =========================================================
	// SLICES EN GO — El tipo de colección más usado
	// =========================================================
	// Un slice es una VISTA sobre un array subyacente.
	// Tiene 3 componentes internos: puntero, longitud (len), capacidad (cap).
	// Es el equivalente funcional al array de JavaScript, pero con matices.

	// --- Creación ---
	// Literal (más común)
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", numeros)
	fmt.Printf("len=%d, cap=%d\n", len(numeros), cap(numeros))

	// make([]T, longitud, capacidad) — preasigna memoria
	// Útil cuando sabes cuántos elementos tendrás
	scores := make([]int, 5, 10) // len=5, cap=10, todos en 0
	fmt.Println("make:", scores)

	// Slice nil (zero value de slice) — len=0, cap=0
	var nulo []int
	fmt.Println("nil:", nulo, len(nulo), nulo == nil) // [] 0 true

	// --- append: agregar elementos (como push en JS) ---
	frutas := []string{"manzana", "banana"}
	frutas = append(frutas, "cereza")        // un elemento
	frutas = append(frutas, "durazno", "uva") // varios
	fmt.Println("\nappend:", frutas)

	// --- Spread con ...: equivalente al spread de JS ---
	extra := []string{"kiwi", "mango"}
	frutas = append(frutas, extra...) // como: [...frutas, ...extra]
	fmt.Println("spread:", frutas)

	// --- Sub-slices: slice[low:high] ---
	// Retorna una vista del slice original (NO copia!)
	// En JS: array.slice(low, high) retorna una copia
	todos := []int{0, 1, 2, 3, 4, 5, 6}
	primerosTres := todos[0:3] // [0, 1, 2]  → índices 0, 1, 2
	del2al5 := todos[2:5]     // [2, 3, 4]
	desde3 := todos[3:]       // [3, 4, 5, 6]
	hasta4 := todos[:4]       // [0, 1, 2, 3]

	fmt.Println("\nSub-slices:")
	fmt.Println(primerosTres, del2al5, desde3, hasta4)

	// CUIDADO: el sub-slice comparte memoria con el original
	primerosTres[0] = 99
	fmt.Println("Original modificado:", todos) // [99 1 2 3 4 5 6]

	// Para una copia independiente usa copy():
	copia := make([]int, len(primerosTres))
	copy(copia, primerosTres)
	copia[0] = 0
	fmt.Println("Copia independiente:", copia, "original:", primerosTres)

	// --- Eliminar un elemento por índice ---
	// Go no tiene una función remove() nativa, la construyes:
	letras := []string{"a", "b", "c", "d", "e"}
	indiceBorrar := 2 // eliminar "c"
	letras = append(letras[:indiceBorrar], letras[indiceBorrar+1:]...)
	fmt.Println("\nEliminado índice 2:", letras) // [a b d e]

	// --- Funciones del paquete slices (Go 1.21+) ---
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	slices.Sort(nums)
	fmt.Println("\nOrdenado:", nums)
	fmt.Println("Contiene 5:", slices.Contains(nums, 5))
	idx, _ := slices.BinarySearch(nums, 5)
	fmt.Println("Índice de 5:", idx)

	// --- Iteración con range ---
	for i, v := range numeros {
		fmt.Printf("  [%d]=%d\n", i, v)
	}
}
