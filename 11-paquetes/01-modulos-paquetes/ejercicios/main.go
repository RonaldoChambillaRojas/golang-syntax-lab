package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Módulos y Paquetes ===")

	// EJERCICIO 1 — Conceptual + Código ───────────────────────
	//
	// En este ejercicio trabajas dentro del mismo paquete main,
	// pero pensando en visibilidad.
	//
	// Crea los siguientes tipos y funciones en este mismo archivo:
	//
	//   type producto struct {       ← unexported
	//       nombre string
	//       precio float64
	//   }
	//
	//   type Catalogo struct {       ← exported
	//       productos []producto     ← slice de tipo unexported (ok en mismo pkg)
	//   }
	//
	//   func NuevoCatalogo() *Catalogo         ← exported constructor
	//   func (c *Catalogo) Agregar(nombre string, precio float64)  ← exported
	//   func (c *Catalogo) Total() float64     ← exported: suma de precios
	//   func (c *Catalogo) Listar()            ← exported: imprime todos
	//   func (c *Catalogo) aplicarDescuento(pct float64) ← unexported
	//     → multiplica todos los precios por (1 - pct)
	//
	//   func AplicarDescuentoPublico(c *Catalogo, pct float64)  ← exported
	//     → llama a c.aplicarDescuento (puede porque está en el mismo paquete)
	//
	// Prueba:
	//   cat := NuevoCatalogo()
	//   cat.Agregar("Laptop", 1200.0)
	//   cat.Agregar("Mouse", 25.0)
	//   cat.Listar()
	//   fmt.Println("Total:", cat.Total())
	//   AplicarDescuentoPublico(cat, 0.10)
	//   cat.Listar()
	//
	// Reflexiona: ¿Qué pasaría si Catalogo estuviera en otro paquete?
	// Escribe la respuesta como un comentario.

	// TODO
}
