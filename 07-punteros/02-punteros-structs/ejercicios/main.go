package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Punteros y Structs ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Crea un struct Nodo para una lista enlazada simple:
	//
	//   type Nodo struct {
	//       Valor     int
	//       Siguiente *Nodo  // puntero al siguiente nodo (nil si es el último)
	//   }
	//
	// Construye manualmente esta lista: 1 → 2 → 3 → 4 → 5 → nil
	// (crea cada nodo y enlaza con el campo Siguiente)
	//
	// Luego implementa: imprimirLista(cabeza *Nodo)
	//   → recorre la lista desde cabeza hasta nil
	//   → imprime: "1 → 2 → 3 → 4 → 5 → nil"
	//
	// E implementa: longitudLista(cabeza *Nodo) int
	//   → cuenta los nodos de la lista

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Expande la lista enlazada con estas funciones:
	//
	//   agregarAlFrente(cabeza **Nodo, valor int)
	//     → crea un nuevo nodo y lo pone al inicio
	//     → necesita **Nodo porque modifica el puntero cabeza
	//
	//   agregarAlFinal(cabeza *Nodo, valor int)
	//     → recorre hasta el último nodo y lo agrega
	//     → ¿necesita puntero o no? piénsalo antes de implementar
	//
	//   eliminarValor(cabeza **Nodo, valor int) bool
	//     → elimina el primer nodo con ese valor
	//     → retorna true si lo encontró y eliminó
	//
	// Prueba con:
	//   1. Lista inicial: 3 → 1 → 4 → 1 → 5
	//   2. Agrega 0 al frente
	//   3. Agrega 9 al final
	//   4. Elimina el valor 1
	//   5. Imprime la lista en cada paso

	// TODO
}
