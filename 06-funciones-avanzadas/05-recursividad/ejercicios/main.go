package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Recursividad ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: esPalindromo(s string) bool
	//   - Un string es palíndromo si se lee igual al derecho que al revés.
	//   - Usa recursividad (NO uses strings.Reverse ni slices).
	//   - Caso base: string vacío o de 1 caracter → true
	//   - Caso recursivo: primer y último carácter iguales Y el centro es palíndromo
	//
	// Pista: convierte a []rune primero para manejar Unicode correctamente.
	//   runes := []rune(s)
	//   Compara runes[0] con runes[len-1], luego llama recursivo con runes[1:len-1]
	//
	// Prueba con: "racecar", "nivel", "oso", "golang", "A man a plan a canal Panama"
	// (para el último, quita espacios primero con strings.ReplaceAll y pasa a minúsculas)

	// TODO

	// EJERCICIO 2 — Intermedio (desafiante) ─────────────────────
	//
	// Implementa: combinaciones(elementos []string, k int) [][]string
	//   - Retorna todas las combinaciones de k elementos del slice.
	//   - El orden dentro de cada combinación no importa.
	//   - No repetir elementos.
	//
	// Ejemplo:
	//   combinaciones(["A","B","C","D"], 2)
	//   → [["A","B"], ["A","C"], ["A","D"], ["B","C"], ["B","D"], ["C","D"]]
	//
	// Algoritmo recursivo:
	//   - Si k == 0, retorna [[]] (una combinación vacía)
	//   - Si len(elementos) < k, retorna [] (imposible)
	//   - Para el primer elemento:
	//     a) Inclúyelo: combina el resto con k-1 y prepende el primero
	//     b) Exclúyelo: combina el resto con k
	//     Retorna la unión de a) y b)
	//
	// Prueba: combinaciones(["rojo","verde","azul","amarillo"], 3)

	// TODO
}
