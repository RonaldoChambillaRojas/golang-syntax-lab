package main

import "fmt"

// =========================================================
// FUNCIONES EN GO
// =========================================================
// Sintaxis: func nombre(params) tipoRetorno { ... }
// En JS:    function nombre(params) { ... }
//           const nombre = (params) => { ... }

// --- Función simple sin parámetros ni retorno ---
func saludar() {
	fmt.Println("¡Hola desde una función!")
}

// --- Función con parámetros y retorno ---
// Los tipos van DESPUÉS del nombre del parámetro (al revés que en TS)
// TS:  function sumar(a: number, b: number): number
// Go:  func sumar(a int, b int) int
func sumar(a int, b int) int {
	return a + b
}

// Cuando los parámetros tienen el mismo tipo, puedes agruparlos:
func multiplicar(a, b int) int { // equivale a (a int, b int)
	return a * b
}

// --- Función con múltiples parámetros de distintos tipos ---
func formatearMensaje(nombre string, edad int, activo bool) string {
	return fmt.Sprintf("Nombre: %s, Edad: %d, Activo: %t", nombre, edad, activo)
}

// --- Funciones son ciudadanos de primera clase (como en JS) ---
// Puedes asignarlas a variables, pasarlas como parámetros, etc.
func aplicarOperacion(a, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

func main() {
	// Llamar funciones
	saludar()
	fmt.Println("3 + 4 =", sumar(3, 4))
	fmt.Println("3 * 4 =", multiplicar(3, 4))
	fmt.Println(formatearMensaje("Ana", 28, true))

	// --- Funciones como valores (first-class functions) ---
	// Igual que en JS: const suma = (a, b) => a + b
	restar := func(a, b int) int {
		return a - b
	}
	fmt.Println("10 - 4 =", restar(10, 4))

	// Pasar función como argumento
	resultado := aplicarOperacion(10, 5, multiplicar)
	fmt.Println("Aplicar multiplicar(10,5):", resultado)

	// Función anónima inline (como arrow function de JS)
	resultado = aplicarOperacion(10, 5, func(a, b int) int {
		return a + b + 1 // operación personalizada inline
	})
	fmt.Println("Operación inline:", resultado)

	// --- Funciones en slices ---
	// En JS: const operaciones = [(a,b) => a+b, (a,b) => a-b]
	operaciones := []func(int, int) int{
		sumar,
		multiplicar,
		func(a, b int) int { return a - b },
	}

	a, b := 6, 3
	nombres := []string{"suma", "multiplicación", "resta"}
	for i, op := range operaciones {
		fmt.Printf("%s(%d, %d) = %d\n", nombres[i], a, b, op(a, b))
	}

	// --- Función que retorna una función (veremos más en closures, sec. 6) ---
	doble := multiplicadorPor(2)
	triple := multiplicadorPor(3)
	fmt.Println("\ndoble(5):", doble(5))
	fmt.Println("triple(5):", triple(5))
}

// multiplicadorPor retorna una función que multiplica por n.
// Es el patrón de "factory function" de JavaScript.
func multiplicadorPor(n int) func(int) int {
	return func(x int) int {
		return x * n
	}
}
