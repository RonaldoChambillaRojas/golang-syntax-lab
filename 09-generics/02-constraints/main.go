package main

import (
	"fmt"
	"strings"
)

// =========================================================
// CONSTRAINTS — Restricciones de tipo en Generics
// =========================================================
// Las constraints definen qué tipos puede aceptar un type parameter.
// Son interfaces, pero usadas en el contexto de generics.

// --- comparable: tipos que soportan == y != ---
// Incluye: int, float64, string, bool, arrays, structs (con campos comparables)
// No incluye: slices, maps, funciones (no son comparables con ==)

func indiceEn[T comparable](slice []T, objetivo T) int {
	for i, v := range slice {
		if v == objetivo {
			return i
		}
	}
	return -1
}

// Frecuencia de elementos usando map con clave comparable
func frecuencia[T comparable](slice []T) map[T]int {
	resultado := make(map[T]int)
	for _, v := range slice {
		resultado[v]++
	}
	return resultado
}

// --- Ordered: tipos que soportan <, >, <=, >= ---
// Definimos nuestro propio constraint ya que cmp.Ordered existe en Go 1.21+
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// El ~ significa "cualquier tipo cuyo tipo subyacente sea X"
// Así, si defines: type MiInt int, Ordered lo acepta también

func minimo[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func maximo[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min3[T Ordered](a, b, c T) T {
	return minimo(minimo(a, b), c)
}

func ordenar[T Ordered](slice []T) []T {
	copia := make([]T, len(slice))
	copy(copia, slice)
	// Bubble sort simple para demostrar el concepto
	for i := 0; i < len(copia); i++ {
		for j := 0; j < len(copia)-i-1; j++ {
			if copia[j] > copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia
}

// --- Constraint personalizada ---
type Stringer interface {
	String() string
}

func imprimirTodos[T Stringer](items []T) {
	parts := make([]string, len(items))
	for i, item := range items {
		parts[i] = item.String()
	}
	fmt.Println("[" + strings.Join(parts, ", ") + "]")
}

// --- Union de tipos en constraint ---
type Numero interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func sumar[T Numero](nums ...T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func promedio[T Numero](nums ...T) float64 {
	if len(nums) == 0 {
		return 0
	}
	suma := sumar(nums...)
	return float64(suma) / float64(len(nums))
}

// --- Tipo con constraint personalizada y ~  ---
type MiInt int

func (m MiInt) String() string {
	return fmt.Sprintf("MiInt(%d)", int(m))
}

func main() {
	// --- comparable ---
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	strs := []string{"go", "rust", "go", "python", "go", "rust"}

	fmt.Println("Índice de 5:", indiceEn(ints, 5))        // 4
	fmt.Println("Índice de 'go':", indiceEn(strs, "go"))  // 0
	fmt.Println("Índice de 99:", indiceEn(ints, 99))       // -1

	fmt.Println("\nFrecuencia ints:", frecuencia(ints))
	fmt.Println("Frecuencia strs:", frecuencia(strs))

	// --- Ordered ---
	fmt.Println("\nMínimo(3, 7):", minimo(3, 7))
	fmt.Println("Mínimo(\"abc\", \"xyz\"):", minimo("abc", "xyz"))
	fmt.Println("Máximo(3.14, 2.71):", maximo(3.14, 2.71))
	fmt.Println("Min de 3:", min3(10, 5, 8))

	nums := []float64{3.14, 1.41, 2.71, 1.73, 0.57}
	fmt.Println("\nOrdenado:", ordenar(nums))
	palabras := []string{"banana", "apple", "cherry", "date"}
	fmt.Println("Ordenado:", ordenar(palabras))

	// --- Union de tipos numéricos ---
	fmt.Println("\nSuma ints:", sumar(1, 2, 3, 4, 5))
	fmt.Println("Suma float:", sumar(1.1, 2.2, 3.3))
	fmt.Println("Promedio:", promedio(10, 20, 30, 40, 50))

	// --- ~ permite tipos derivados ---
	var a, b MiInt = 10, 20
	fmt.Println("\nMínimo MiInt:", minimo(a, b)) // MiInt es ~int → funciona
}
