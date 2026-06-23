package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Channels ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un generador de números de Fibonacci usando channels:
	//
	//   func fibonacci(n int) <-chan int
	//     → retorna un channel que emite los primeros n números de Fibonacci
	//     → usa una goroutine interna que genera y envía los valores
	//     → cierra el channel cuando termina
	//
	// Prueba:
	//   for v := range fibonacci(10) {
	//       fmt.Print(v, " ")
	//   }
	//   // 0 1 1 2 3 5 8 13 21 34

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un pipeline de procesamiento con channels:
	//
	//   Stage 1 - generar(nums ...int) <-chan int
	//     → pone los números en un channel y lo retorna
	//
	//   Stage 2 - cuadrado(in <-chan int) <-chan int
	//     → lee del canal in, eleva al cuadrado, escribe en nuevo canal
	//
	//   Stage 3 - filtrarMayores(in <-chan int, umbral int) <-chan int
	//     → pasa solo los valores > umbral
	//
	//   Stage 4 - sumar(in <-chan int) int
	//     → suma todos los valores del canal
	//
	// Encadénalos:
	//   nums := generar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//   cuadrados := cuadrado(nums)
	//   grandes := filtrarMayores(cuadrados, 25)
	//   total := sumar(grandes)
	//   fmt.Println(total)  // 36+49+64+81+100 = 330
	//
	// Cada stage vive en su propia goroutine.

	// TODO
}
