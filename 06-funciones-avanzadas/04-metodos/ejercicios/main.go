package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Métodos ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Crea un struct CuentaBancaria con:
	//   - Titular string
	//   - Saldo   float64
	//
	// Agrega estos métodos:
	//   (c *CuentaBancaria) Depositar(monto float64) error
	//     → si monto <= 0, retorna error "monto debe ser positivo"
	//     → suma al saldo, retorna nil
	//
	//   (c *CuentaBancaria) Retirar(monto float64) error
	//     → si monto <= 0, retorna error "monto debe ser positivo"
	//     → si monto > saldo, retorna error "fondos insuficientes"
	//     → resta del saldo, retorna nil
	//
	//   (c CuentaBancaria) String() string
	//     → "Cuenta de <Titular>: $<Saldo con 2 decimales>"
	//
	// Prueba con una cuenta: deposita 1000, retira 250, intenta retirar 900.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa una pila (stack) genérica usando métodos:
	//
	//   type Pila struct {
	//       elementos []int
	//   }
	//
	//   (p *Pila) Push(v int)
	//     → agrega al tope de la pila
	//
	//   (p *Pila) Pop() (int, error)
	//     → extrae y retorna el elemento del tope
	//     → si está vacía, retorna error "pila vacía"
	//
	//   (p *Pila) Peek() (int, error)
	//     → retorna el tope sin extraerlo
	//     → si está vacía, retorna error "pila vacía"
	//
	//   (p Pila) Len() int
	//     → retorna el número de elementos
	//
	//   (p Pila) String() string
	//     → "Pila[<elementos de base a tope>]"
	//
	// Prueba: push 5 elementos, peek, pop 3 veces, imprime estado final.

	// TODO
}
