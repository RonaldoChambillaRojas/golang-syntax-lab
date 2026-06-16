package main

import (
	"errors"
	"fmt"
)

// =========================================================
// PANIC Y RECOVER
// =========================================================
// panic es el equivalente Go de throw, pero solo debe usarse cuando
// el programa está en un estado irrecuperable.
//
// JavaScript:   throw new Error("crítico")
// Go:           panic("crítico")   ← o panic(error)
//
// La diferencia filosófica:
//   JS: throw es para flujo normal de errores (404, validaciones, etc.)
//   Go: panic es para errores de PROGRAMACIÓN (invariantes rotos, índices
//       fuera de rango, nil pointer, bugs) — no para lógica de negocio.
//
// Para lógica de negocio → siempre retorna (valor, error)

// --- Cuándo Go usa panic automáticamente ---
// - Índice fuera de rango: slice[100] cuando len=5
// - Desreferencia de nil: (*nilPointer).Campo
// - Division por cero en enteros: 5/0
// - Type assertion fallida: x.(string) cuando x no es string
// - Enviar a un channel cerrado

// --- recover: intercepta el panic ---
// recover() solo funciona dentro de una función diferida (defer).
// Si no hay panic activo, recover() retorna nil.

func recuperarPanic() {
	// Este defer captura cualquier panic en esta goroutine
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic interceptado: %v\n", r)
		}
	}()

	panic("algo salió muy mal")
}

// --- Patrón: convertir panic a error ---
// Útil cuando llamas código de terceros que puede hacer panic
func operacionSegura(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = fmt.Errorf("panic: %w", v)
			case string:
				err = fmt.Errorf("panic: %s", v)
			default:
				err = fmt.Errorf("panic: %v", v)
			}
		}
	}()
	fn()
	return nil
}

// --- Panic con error (la forma recomendada si usas panic) ---
var ErrEstadoInvalido = errors.New("estado inválido")

type Pila struct {
	elementos []int
}

func (p *Pila) Push(v int) {
	p.elementos = append(p.elementos, v)
}

// Pop hace panic si la pila está vacía — comportamiento tipo Go stdlib
// (sync.Mutex hace panic si unlockeas un mutex sin lockear)
func (p *Pila) Pop() int {
	if len(p.elementos) == 0 {
		panic(ErrEstadoInvalido) // panic con un valor de tipo error
	}
	ultimo := p.elementos[len(p.elementos)-1]
	p.elementos = p.elementos[:len(p.elementos)-1]
	return ultimo
}

// PopSeguro envuelve Pop para convertir el panic a error
func (p *Pila) PopSeguro() (v int, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	return p.Pop(), nil
}

func main() {
	// --- Recover básico ---
	fmt.Println("Antes del panic")
	recuperarPanic()
	fmt.Println("El programa continúa después del panic interceptado")

	// --- operacionSegura ---
	fmt.Println()

	err := operacionSegura(func() {
		// Simula índice fuera de rango
		nums := []int{1, 2, 3}
		_ = nums[10] // panic: runtime error
	})
	fmt.Println("Error de operacionSegura:", err)

	err = operacionSegura(func() {
		// Sin panic
		fmt.Println("Operación exitosa")
	})
	fmt.Println("Error (nil esperado):", err)

	// --- Pila con panic ---
	fmt.Println()
	pila := &Pila{}
	pila.Push(1)
	pila.Push(2)

	fmt.Println("Pop:", pila.Pop()) // 2
	fmt.Println("Pop:", pila.Pop()) // 1

	// PopSeguro convierte el panic a error
	if v, err := pila.PopSeguro(); err != nil {
		fmt.Println("PopSeguro error:", err)
		fmt.Println("Is ErrEstadoInvalido:", errors.Is(err, ErrEstadoInvalido))
	} else {
		fmt.Println("PopSeguro:", v)
	}

	// =========================================================
	// REGLA: ¿panic o error?
	// =========================================================
	// panic cuando:
	//   ✓ El programa no puede continuar razonablemente (init() fallida)
	//   ✓ Se viola un invariante del sistema (bug del programador)
	//   ✓ Error de programación que nunca debería pasar en prod
	//
	// error (retorno) cuando:
	//   ✓ El archivo no existe
	//   ✓ La conexión falló
	//   ✓ El input del usuario es inválido
	//   ✓ Cualquier condición que el llamador puede manejar
	fmt.Println("\n=== Regla final ===")
	fmt.Println("Si el llamador puede manejarlo → retorna error")
	fmt.Println("Si es un bug del programador → panic")
}
