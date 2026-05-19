package main

import (
	"errors"
	"fmt"
	"strconv"
)

// =========================================================
// MÚLTIPLES VALORES DE RETORNO — Feature exclusivo de Go
// =========================================================
// En JavaScript para retornar múltiples valores necesitas:
//   - Un objeto: return { valor, error }
//   - Una tupla/array: return [valor, error]
//   - Lanzar excepciones
//
// En Go, la convención es retornar (resultado, error) o (valor1, valor2).
// Esto hace el manejo de errores EXPLÍCITO y obligatorio.

// --- Patrón estándar: (valor, error) ---
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil // nil = sin error
}

// --- Múltiples valores sin error ---
func minMax(numeros []int) (int, int) {
	if len(numeros) == 0 {
		return 0, 0
	}
	min, max := numeros[0], numeros[0]
	for _, n := range numeros[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// =========================================================
// NAMED RETURN VALUES (retornos con nombre)
// =========================================================
// Puedes nombrar los valores de retorno en la firma.
// Útil para documentar qué retorna cada valor, o en funciones cortas.

func estadisticas(nums []float64) (promedio float64, suma float64, cantidad int) {
	cantidad = len(nums)
	if cantidad == 0 {
		return // "naked return": retorna los valores nombrados en su estado actual
	}
	for _, n := range nums {
		suma += n
	}
	promedio = suma / float64(cantidad)
	return // naked return: retorna promedio, suma, cantidad
}

// =========================================================
// BLANK IDENTIFIER _ — descartar valores que no necesitas
// =========================================================
// En JS no tienes esta necesidad porque puedes ignorar cualquier retorno.
// En Go si declaras una variable debes usarla, así que usas _ para ignorar.

func parsearEntero(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
	// --- Usar el patrón (valor, error) ---
	resultado, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", resultado)
	}

	// Error case
	_, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	// --- Múltiples valores ---
	numeros := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	min, max := minMax(numeros)
	fmt.Printf("\nMínimo: %d, Máximo: %d\n", min, max)

	// Solo quiero el máximo → ignoro el mínimo con _
	_, soloMax := minMax(numeros)
	fmt.Println("Solo máximo:", soloMax)

	// --- Named returns ---
	prom, suma, cant := estadisticas([]float64{10, 20, 30, 40, 50})
	fmt.Printf("\nEstadísticas: promedio=%.2f, suma=%.2f, cantidad=%d\n",
		prom, suma, cant)

	// Caso vacío
	prom, suma, cant = estadisticas([]float64{})
	fmt.Printf("Vacío: promedio=%.2f, suma=%.2f, cantidad=%d\n",
		prom, suma, cant)

	// --- Blank identifier en varios contextos ---
	// Ignorar el error (solo si estás SEGURO de que no puede fallar)
	num, _ := parsearEntero("42")
	fmt.Println("\nParsed:", num)

	// El _ en range cuando no necesitas el índice:
	for _, v := range []int{1, 2, 3} {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// =========================================================
	// ANTI-PATRÓN: nunca ignores errores de forma descuidada
	// =========================================================
	// Esto es MALO:
	//   valor, _ := operacionQuePodriaFallar()
	// Si ignoras el error, tu programa puede fallar silenciosamente más tarde.
	// Solo usa _ cuando estés 100% seguro de que no puede fallar,
	// o cuando el error es verdaderamente irrelevante.
}
