package main

import "fmt"

// =========================================================
// CLOSURES EN GO
// =========================================================
// Un closure es una función que captura variables de su entorno.
// En Go funcionan exactamente igual que en JavaScript.
// Las variables capturadas son compartidas por referencia.

// --- Closure clásico: generador de contador ---
// En JS: function contador() { let n = 0; return () => ++n }
func nuevoContador() func() int {
	n := 0 // n vive en el heap porque el closure lo referencia
	return func() int {
		n++
		return n
	}
}

// --- Múltiples closures comparten el mismo estado ---
func nuevoContadorConReset() (incrementar func(), decrementar func(), resetear func(), leer func() int) {
	n := 0
	incrementar = func() { n++ }
	decrementar = func() { n-- }
	resetear = func() { n = 0 }
	leer = func() int { return n }
	return
}

// --- Closure para encapsular configuración ---
// Patrón: factory que retorna una función configurada
func multiplicadorPor(factor int) func(int) int {
	// factor queda capturado en el closure
	return func(x int) int {
		return x * factor
	}
}

// --- Memoización con closure ---
// Misma idea que en JS: cachear resultados costosos
func memoFibonacci() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if resultado, existe := cache[n]; existe {
			return resultado
		}
		resultado := fib(n-1) + fib(n-2)
		cache[n] = resultado
		return resultado
	}

	return fib
}

// --- Trampa clásica: closures en un loop ---
// Es EL bug más famoso de closures, igual en Go que en JS
func trampaLoop() {
	fmt.Println("\n=== Trampa: closure en loop ===")

	// BUG: todos los closures capturan la MISMA variable i
	funcsBug := make([]func(), 3)
	for i := 0; i < 3; i++ {
		funcsBug[i] = func() {
			fmt.Println(i) // captura la variable i, no su valor
		}
	}
	fmt.Println("Con bug (todos imprimen 3):")
	for _, f := range funcsBug {
		f() // imprime 3, 3, 3 — porque i llegó a 3 al final del loop
	}

	// SOLUCIÓN: capturar el valor en una variable local
	funcsFix := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // crea una nueva variable i por iteración (shadow)
		funcsFix[i] = func() {
			fmt.Println(i) // captura la i local de ESTA iteración
		}
	}
	fmt.Println("Corregido (imprime 0, 1, 2):")
	for _, f := range funcsFix {
		f()
	}

	// En Go 1.22+ el comportamiento del loop cambió:
	// for range crea una nueva variable por iteración automáticamente.
	// Pero entender el bug sigue siendo importante.
}

func main() {
	// --- Contador básico ---
	contar := nuevoContador()
	fmt.Println(contar()) // 1
	fmt.Println(contar()) // 2
	fmt.Println(contar()) // 3

	// Cada llamada a nuevoContador() crea un estado INDEPENDIENTE
	contarB := nuevoContador()
	fmt.Println("Contador B:", contarB()) // 1 (independiente de contar)

	// --- Contador con múltiples operaciones ---
	inc, dec, reset, leer := nuevoContadorConReset()
	inc()
	inc()
	inc()
	fmt.Println("\nDespués de 3 inc:", leer()) // 3
	dec()
	fmt.Println("Después de dec:", leer()) // 2
	reset()
	fmt.Println("Después de reset:", leer()) // 0

	// --- Factory functions ---
	doble := multiplicadorPor(2)
	triple := multiplicadorPor(3)
	fmt.Println("\ndoble(7):", doble(7))   // 14
	fmt.Println("triple(7):", triple(7))  // 21

	// --- Fibonacci con memoización ---
	fib := memoFibonacci()
	fmt.Println("\nFibonacci:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fib(i))
	}
	fmt.Println("fib(40):", fib(40)) // rápido gracias al cache

	trampaLoop()
}
