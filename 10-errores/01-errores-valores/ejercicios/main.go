package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Errores como Valores ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un tipo ErrRangoInvalido:
	//   type ErrRangoInvalido struct {
	//       Valor int
	//       Min   int
	//       Max   int
	//   }
	//   func (e *ErrRangoInvalido) Error() string
	//     → "valor <Valor> fuera del rango [<Min>, <Max>]"
	//
	// Implementa: validarRango(valor, min, max int) error
	//   → retorna *ErrRangoInvalido si valor < min o valor > max
	//   → retorna nil si está en rango
	//
	// Implementa: convertirALetra(nota int) (string, error)
	//   → usa validarRango(nota, 0, 100) primero
	//   → si ok: 90-100=A, 80-89=B, 70-79=C, 60-69=D, 0-59=F
	//   → si error: propaga con fmt.Errorf("convertirALetra: %w", err)
	//
	// Prueba con notas: 85, 50, -10, 110, 100
	// Usa errors.As para extraer el *ErrRangoInvalido cuando falla.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un parser de CSV simple con manejo de errores:
	//
	//   type ErrParseo struct {
	//       Linea  int
	//       Col    int
	//       Motivo string
	//   }
	//   func (e *ErrParseo) Error() string → "línea <L>, col <C>: <Motivo>"
	//
	//   type Fila struct { Nombre string; Edad int; Email string }
	//
	//   func parsearCSV(csv string) ([]Fila, error)
	//     → el formato es: "nombre,edad,email\n..."
	//     → separa líneas con strings.Split(csv, "\n")
	//     → separa campos con strings.Split(linea, ",")
	//     → valida que cada línea tenga exactamente 3 campos
	//     → convierte edad a int con strconv.Atoi
	//     → si algo falla, retorna ErrParseo con la línea y columna del error
	//
	// Prueba con:
	//   csv := "Ana,28,ana@mail.com\nCarlos,abc,carlos@mail.com\nMaría,25,maria@mail.com"
	//   La segunda línea debería fallar en col=2 (edad inválida)

	// TODO
}
