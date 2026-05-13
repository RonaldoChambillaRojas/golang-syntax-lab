package main

import (
	"fmt"
	"math"
)

func main() {
	// =========================================================
	// OPERADORES ARITMÉTICOS
	// =========================================================
	// Casi idénticos a JS, con dos diferencias importantes:
	// 1. La división entera entre ints da int (sin decimales)
	// 2. No existe ** para potencias (usa math.Pow)

	a, b := 10, 3

	fmt.Println("Suma:      ", a+b)  // 13
	fmt.Println("Resta:     ", a-b)  // 7
	fmt.Println("Mult:      ", a*b)  // 30
	fmt.Println("División:  ", a/b)  // 3 (entero! no 3.333)
	fmt.Println("Módulo:    ", a%b)  // 1

	// Para división con decimales, convierte a float64 primero:
	fmt.Println("Div float: ", float64(a)/float64(b)) // 3.3333...

	// Potencia: math.Pow (trabaja con float64)
	fmt.Println("Potencia:  ", math.Pow(2, 10)) // 1024

	// =========================================================
	// OPERADORES DE INCREMENTO/DECREMENTO
	// =========================================================
	// En JS: x++ puede ser expresión: y = x++
	// En Go: x++ es una SENTENCIA, no una expresión.
	// No puedes hacer: y := x++ ni pasarlo como argumento.

	contador := 5
	contador++ // incrementa en 1
	contador-- // decrementa en 1
	fmt.Println("Contador:", contador) // 5

	// =========================================================
	// OPERADORES DE COMPARACIÓN
	// =========================================================
	// Igual a JS, pero sin == vs === (en Go == siempre compara valor y tipo)
	// No existe el problema de == vs === porque no hay coerción implícita.

	x, y := 10, 20
	fmt.Println("\n== Comparaciones ==")
	fmt.Println(x == y)  // false
	fmt.Println(x != y)  // true
	fmt.Println(x < y)   // true
	fmt.Println(x > y)   // false
	fmt.Println(x <= 10) // true
	fmt.Println(x >= 10) // true

	// =========================================================
	// OPERADORES LÓGICOS
	// =========================================================
	// Igual que JS: &&, ||, !
	// Short-circuit evaluation también aplica (&&: si el primero es false,
	// el segundo no se evalúa; ||: si el primero es true, el segundo no)

	t, f := true, false
	fmt.Println("\n== Lógicos ==")
	fmt.Println(t && f) // false
	fmt.Println(t || f) // true
	fmt.Println(!t)     // false

	// =========================================================
	// OPERADORES BIT A BIT
	// =========================================================
	// Iguales a JS pero más usados en Go (para flags, permisos, etc.)

	p, q := 0b1100, 0b1010 // notación binaria
	fmt.Println("\n== Bit a bit ==")
	fmt.Printf("AND:    %04b\n", p&q)  // 1000
	fmt.Printf("OR:     %04b\n", p|q)  // 1110
	fmt.Printf("XOR:    %04b\n", p^q)  // 0110
	fmt.Printf("LSHIFT: %04b\n", p<<1) // 11000
	fmt.Printf("RSHIFT: %04b\n", p>>1) // 0110

	// =========================================================
	// OPERADORES DE ASIGNACIÓN
	// =========================================================

	n := 10
	n += 5  // n = n + 5 = 15
	n -= 3  // n = n - 3 = 12
	n *= 2  // n = n * 2 = 24
	n /= 4  // n = n / 4 = 6
	n %= 4  // n = n % 4 = 2
	fmt.Println("\nAsignación compuesta:", n) // 2

	// =========================================================
	// NO EXISTE EL OPERADOR TERNARIO
	// =========================================================
	// En JS: const max = a > b ? a : b
	// En Go debes usar if/else:
	max := a
	if b > a {
		max = b
	}
	fmt.Println("Máximo:", max)
}
