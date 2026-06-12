package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Pipelines con Generics ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa la función genérica:
	//   Chunk[T any](slice []T, tamaño int) [][]T
	//   → divide el slice en sub-slices de tamaño `tamaño`
	//   → el último chunk puede ser más pequeño
	//
	// Ejemplo:
	//   Chunk([]int{1,2,3,4,5,6,7}, 3) → [[1 2 3] [4 5 6] [7]]
	//   Chunk([]string{"a","b","c","d"}, 2) → [[a b] [c d]]
	//
	// También implementa:
	//   Unique[T comparable](slice []T) []T
	//   → retorna el slice sin duplicados, preservando el orden
	//
	// Prueba Unique con: []int{3,1,4,1,5,9,2,6,5,3}
	// Esperado: [3 1 4 5 9 2 6]

	// TODO

	// EJERCICIO 2 — Intermedio (desafiante) ─────────────────────
	//
	// Implementa un pipeline funcional que procese una lista de productos:
	//
	//   type Producto struct {
	//       Nombre    string
	//       Categoria string
	//       Precio    float64
	//       Stock     int
	//   }
	//
	// Con estas funciones genéricas (ya definidas en el tema anterior o reimplementa):
	//   Filter, Map, Reduce, GroupBy
	//
	// Aplica el siguiente pipeline sobre estos datos:
	//   productos := []Producto{
	//       {"Laptop", "Electrónica", 1200.0, 5},
	//       {"Mouse", "Electrónica", 25.0, 50},
	//       {"Camisa", "Ropa", 35.0, 100},
	//       {"Teclado", "Electrónica", 89.0, 20},
	//       {"Pantalón", "Ropa", 60.0, 30},
	//       {"Monitor", "Electrónica", 350.0, 8},
	//   }
	//
	// 1. Filtra solo los con Stock > 10
	// 2. Agrupa por Categoria
	// 3. Para cada categoría: calcula el precio promedio y el total de items
	// 4. Imprime:
	//    "Electrónica: 3 productos, precio promedio $XX.XX"
	//    "Ropa: 2 productos, precio promedio $XX.XX"
	//
	// Bonus: implementa todo encadenado, no en variables separadas.

	// TODO
}
