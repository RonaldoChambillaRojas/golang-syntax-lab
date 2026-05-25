package main

import "fmt"

// =========================================================
// PUNTEROS BÁSICOS
// =========================================================

func main() {
	// --- Declarar un puntero ---
	x := 42
	p := &x // p es un *int: "puntero a int"

	fmt.Printf("x:  valor=%d, dirección=%p\n", x, &x)
	fmt.Printf("p:  valor=%p (dirección que guarda), *p=%d\n", p, *p)

	// --- Modificar a través del puntero ---
	*p = 100 // desreferencia y asigna
	fmt.Println("x después de *p = 100:", x) // 100 — x fue modificado

	// --- Por qué esto importa: paso por valor vs referencia ---
	a := 5
	incrementarValor(a)
	fmt.Println("\nDespués de incrementarValor:", a) // 5 — NO cambió

	incrementarPuntero(&a)
	fmt.Println("Después de incrementarPuntero:", a) // 6 — SÍ cambió

	// --- new(): crea un puntero con zero value ---
	// Como new en JS (pero diferente semántica)
	pNum := new(int) // *int apuntando a un 0
	fmt.Println("\nnew(int):", *pNum) // 0
	*pNum = 42
	fmt.Println("Después de asignar:", *pNum)

	// --- Puntero nil: cuidado ---
	var pNil *int
	fmt.Println("\nPuntero nil:", pNil)        // <nil>
	fmt.Println("Es nil:", pNil == nil)        // true
	// *pNil = 5 // ✗ PANIC: runtime error: invalid memory address
	// Siempre verifica nil antes de desreferenciar un puntero externo

	// --- Punteros a strings y otros tipos ---
	texto := "hola"
	pTexto := &texto
	*pTexto = "mundo"
	fmt.Println("\nTexto modificado:", texto) // mundo

	// --- Dos punteros al mismo lugar ---
	n := 10
	p1 := &n
	p2 := &n
	*p1 = 50
	fmt.Printf("\nn=%d, *p1=%d, *p2=%d\n", n, *p1, *p2) // todos 50

	// --- Retornar puntero desde función ---
	// SEGURO en Go: el valor escapa al heap automáticamente
	ptr := crearEntero(99)
	fmt.Println("\nPuntero retornado:", *ptr)
}

// Esta función recibe una COPIA de a — no modifica el original
func incrementarValor(a int) {
	a++
	// la a local se descarta al salir de la función
}

// Esta función recibe la DIRECCIÓN de la variable — modifica el original
func incrementarPuntero(a *int) {
	*a++
}

// crearEntero retorna un puntero a un int.
// En C esto sería un dangling pointer; en Go el escape analysis
// detecta que el valor debe vivir en el heap. Es seguro.
func crearEntero(v int) *int {
	resultado := v
	return &resultado // Go mueve esto al heap automáticamente
}
