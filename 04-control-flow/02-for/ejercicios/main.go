package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: For ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Tabla de multiplicar:
	// Imprime la tabla de multiplicar del 7 en este formato:
	//   7 x 1 = 7
	//   7 x 2 = 14
	//   ...
	//   7 x 12 = 84
	//
	// Luego, usando un for con range sobre un slice de ints
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], calcula la suma de todos
	// los cuadrados (1² + 2² + 3² + ... + 10²) e imprímela.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Criba de Eratóstenes: encuentra todos los primos hasta N=50
	//
	// Algoritmo:
	// 1. Crea un slice de bool de tamaño N+1, todos en true (esPrimo).
	//    Pista: make([]bool, N+1) da todos en false, necesitas inicializarlos.
	//    O usa: esPrimo := make([]bool, N+1); for i := range esPrimo { esPrimo[i] = true }
	// 2. esPrimo[0] y esPrimo[1] = false (0 y 1 no son primos).
	// 3. Para i de 2 hasta sqrt(N):
	//    Si esPrimo[i] es true, marca todos sus múltiplos como false:
	//    para j = i*i; j <= N; j += i → esPrimo[j] = false
	// 4. Imprime todos los números donde esPrimo[i] == true.
	//
	// Resultado esperado: 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47
	//
	// Pista: para sqrt puedes usar un for con i*i <= N como condición.

	// TODO
}
