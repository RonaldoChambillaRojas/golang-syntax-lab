package main

import "fmt"

// =========================================================
// MUTABILIDAD EN GO — Cuándo usar punteros
// =========================================================

// Regla de oro: usa puntero cuando:
// 1. La función necesita modificar el valor
// 2. El struct es grande (> 64 bytes aproximadamente)
// 3. Necesitas representar "sin valor" (nil)

// --- Guía de decisión por tipo ---

// int, float64, bool, string → pasa por valor casi siempre
func duplicar(n int) int { return n * 2 } // ✓ por valor

// Struct pequeño y no se modifica → por valor
type Punto struct{ X, Y float64 }
func distancia(a, b Punto) float64 { // ✓ por valor
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}

// Struct que se modifica → por puntero
type Contador struct {
	n      int
	nombre string
}

func (c *Contador) Incrementar() { c.n++ } // ✓ por puntero
func (c *Contador) Valor() int  { return c.n }

// Struct grande → por puntero para evitar copias costosas
type ConfiguracionServidor struct {
	Host          string
	Puerto        int
	MaxConexiones int
	Timeout       int
	Certificados  [100]byte // 100 bytes solo en este campo
}

func configurarTLS(c *ConfiguracionServidor) { // ✓ por puntero (struct grande)
	c.Puerto = 443
}

// =========================================================
// SLICES Y MAPS — Tienen referencia interna
// =========================================================
// Estos tipos ya contienen un puntero internamente.
// Para MODIFICAR elementos: no necesitas *[]T
// Para RE-ASIGNAR (append, reslice): sí necesitas *[]T

func agregarAlSlice(s *[]int, v int) {
	*s = append(*s, v) // re-asigna el header del slice
}

func modificarPrimero(s []int, v int) {
	if len(s) > 0 {
		s[0] = v // modifica el array subyacente (compartido)
	}
}

// Los maps siempre se pasan por referencia (su puntero interno)
func agregarAlMap(m map[string]int, clave string, valor int) {
	m[clave] = valor // NO necesitas *map[string]int
}

func main() {
	// --- Tipos valor ---
	x := 5
	fmt.Println("duplicar(x):", duplicar(x)) // retorna nuevo valor
	fmt.Println("x sin cambios:", x)

	// --- Struct por valor ---
	p1 := Punto{0, 0}
	p2 := Punto{3, 4}
	fmt.Printf("\nDistancia²: %.0f\n", distancia(p1, p2)) // 25

	// --- Struct con puntero ---
	c := &Contador{nombre: "visitas"}
	c.Incrementar()
	c.Incrementar()
	c.Incrementar()
	fmt.Printf("%s: %d\n", c.nombre, c.Valor())

	// --- Slice: modificar vs re-asignar ---
	nums := []int{1, 2, 3}
	modificarPrimero(nums, 99) // modifica el elemento
	fmt.Println("\nDespués de modificarPrimero:", nums) // [99 2 3]

	agregarAlSlice(&nums, 4) // re-asigna el slice
	fmt.Println("Después de agregarAlSlice:", nums) // [99 2 3 4]

	// Sin puntero, el append no se ve fuera:
	noFunciona(nums)
	fmt.Println("Después de noFunciona:", nums) // sin cambio

	// --- Map: siempre por referencia ---
	mapa := map[string]int{"a": 1}
	agregarAlMap(mapa, "b", 2)
	fmt.Println("\nMapa:", mapa) // map[a:1 b:2]

	// =========================================================
	// ZERO VALUE vs NO VALUE — punteros como opcionales
	// =========================================================
	// En TypeScript: number | null o number | undefined
	// En Go: *int (nil = sin valor, non-nil = tiene valor)

	var edad *int
	fmt.Println("\nEdad sin asignar:", edad) // <nil>

	edadValor := 25
	edad = &edadValor
	fmt.Println("Edad asignada:", *edad)

	// Función que acepta valor opcional
	imprimirEdad(nil)
	imprimirEdad(edad)
}

func noFunciona(s []int) {
	s = append(s, 999) // el slice local se re-asigna pero el original no
}

func imprimirEdad(edad *int) {
	if edad == nil {
		fmt.Println("Edad: desconocida")
		return
	}
	fmt.Printf("Edad: %d\n", *edad)
}
