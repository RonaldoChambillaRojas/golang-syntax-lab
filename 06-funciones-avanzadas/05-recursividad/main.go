package main

import "fmt"

// =========================================================
// RECURSIVIDAD EN GO
// =========================================================
// La recursividad funciona igual que en JavaScript.
// IMPORTANTE: Go NO optimiza tail-call (TCO).
// Para valores muy grandes, usa iteración o trampolín.

// --- Factorial: ejemplo clásico ---
func factorial(n int) int {
	if n <= 1 { // caso base
		return 1
	}
	return n * factorial(n-1) // llamada recursiva
}

// --- Fibonacci sin memoización: O(2^n) — muy ineficiente ---
func fibLento(n int) int {
	if n <= 1 {
		return n
	}
	return fibLento(n-1) + fibLento(n-2)
}

// --- Fibonacci con memoización: O(n) ---
func fibRapido(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if resultado, ok := memo[n]; ok {
		return resultado
	}
	memo[n] = fibRapido(n-1, memo) + fibRapido(n-2, memo)
	return memo[n]
}

// --- Suma de dígitos ---
// sumaDigitos(1234) → 1+2+3+4 = 10
func sumaDigitos(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumaDigitos(n/10)
}

// --- Potencia ---
func potencia(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 { // optimización: potencia par
		mitad := potencia(base, exp/2)
		return mitad * mitad
	}
	return base * potencia(base, exp-1)
}

// --- Aplanar slice anidado ---
// En JS: array.flat(Infinity)
type AnySlice []interface{}

func aplanar(lista AnySlice) []int {
	var resultado []int
	for _, elem := range lista {
		switch v := elem.(type) {
		case int:
			resultado = append(resultado, v)
		case AnySlice:
			resultado = append(resultado, aplanar(v)...)
		}
	}
	return resultado
}

// --- Búsqueda binaria recursiva ---
func busquedaBinaria(slice []int, objetivo, inicio, fin int) int {
	if inicio > fin {
		return -1 // no encontrado
	}
	medio := (inicio + fin) / 2
	if slice[medio] == objetivo {
		return medio
	}
	if slice[medio] < objetivo {
		return busquedaBinaria(slice, objetivo, medio+1, fin)
	}
	return busquedaBinaria(slice, objetivo, inicio, medio-1)
}

func main() {
	// --- Factorial ---
	fmt.Println("=== Factorial ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  %d! = %d\n", i, factorial(i))
	}

	// --- Fibonacci ---
	fmt.Println("\n=== Fibonacci ===")
	memo := make(map[int]int)
	for i := 0; i <= 15; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fibRapido(i, memo))
	}

	// --- Suma de dígitos ---
	fmt.Println("\n=== Suma de dígitos ===")
	fmt.Printf("  sumaDigitos(1234) = %d\n", sumaDigitos(1234))
	fmt.Printf("  sumaDigitos(9999) = %d\n", sumaDigitos(9999))

	// --- Potencia ---
	fmt.Println("\n=== Potencia ===")
	fmt.Printf("  2^10 = %d\n", potencia(2, 10))
	fmt.Printf("  3^5  = %d\n", potencia(3, 5))

	// --- Aplanar ---
	fmt.Println("\n=== Aplanar ===")
	anidado := AnySlice{1, AnySlice{2, 3, AnySlice{4, 5}}, 6, AnySlice{7}}
	fmt.Println("  Anidado:", anidado)
	fmt.Println("  Aplanado:", aplanar(anidado))

	// --- Búsqueda binaria ---
	fmt.Println("\n=== Búsqueda binaria ===")
	ordenado := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("  Slice:", ordenado)
	fmt.Printf("  Buscar 7: índice %d\n", busquedaBinaria(ordenado, 7, 0, len(ordenado)-1))
	fmt.Printf("  Buscar 6: índice %d\n", busquedaBinaria(ordenado, 6, 0, len(ordenado)-1))
}
