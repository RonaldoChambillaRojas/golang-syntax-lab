package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Mutabilidad y Punteros ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Para cada una de estas funciones, decide si debe usar puntero
	// o valor como receptor/parámetro, y justifica en un comentario:
	//
	//   a) calcularArea(r Rectangulo) float64
	//      Rectangulo tiene Base y Altura float64.
	//      ¿Por valor o por puntero? ¿Por qué?
	//
	//   b) aplicarDescuento(p *Producto, pct float64)
	//      Producto tiene Precio float64 y debe ser modificado.
	//      ¿Correcto así? ¿Por qué?
	//
	//   c) contarPalabras(texto string) map[string]int
	//      No modifica el texto.
	//      ¿El string debe ser puntero? ¿Por qué?
	//
	// Implementa las tres funciones y pruébalas.
	// El objetivo no es solo que compilen sino que entiendas el razonamiento.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "registro de configuración" con valores opcionales:
	//
	//   type Config struct {
	//       MaxConexiones *int      // nil = usar defecto (10)
	//       Timeout       *int      // nil = usar defecto (30 segundos)
	//       Host          *string   // nil = usar defecto ("localhost")
	//   }
	//
	//   func resolverConfig(c Config) (maxConn int, timeout int, host string)
	//     → si un campo es nil, usa el valor por defecto
	//     → si no es nil, usa el valor apuntado
	//
	// Prueba con:
	//   config1 := Config{}                           → todos defaults
	//   config2 := Config{MaxConexiones: intPtr(50)}  → max=50, resto defaults
	//   config3 := Config{Host: strPtr("prod.com"), Timeout: intPtr(60)}
	//
	// Crea funciones helper intPtr(v int) *int y strPtr(v string) *string
	// para construir punteros a literales fácilmente.

	// TODO
}
