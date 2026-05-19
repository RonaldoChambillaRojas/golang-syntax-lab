package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Manejo de Errores ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: raizCuadrada(n float64) (float64, error)
	//   - Si n < 0, retorna un error con message: "no se puede calcular raíz de número negativo"
	//   - Si n >= 0, retorna math.Sqrt(n), nil
	//
	// Implementa: logaritmo(n float64) (float64, error)
	//   - Si n <= 0, retorna error: "logaritmo indefinido para valores ≤ 0"
	//   - Si n > 0, retorna math.Log(n), nil
	//
	// Importa "math" si lo necesitas.
	//
	// Pruébalas con valores válidos e inválidos:
	//   raizCuadrada(16.0)  → 4, nil
	//   raizCuadrada(-4.0)  → error
	//   logaritmo(1.0)      → 0, nil
	//   logaritmo(0)        → error

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Crea un sistema de validación para un formulario de registro.
	//
	// Define un tipo de error personalizado:
	//   type ErrorFormulario struct {
	//       Errores map[string]string // campo → mensaje
	//   }
	//   func (e *ErrorFormulario) Error() string { ... }
	//     → debe retornar todos los errores juntos, ej:
	//       "errores de validación: email: inválido; password: muy corta"
	//
	// Implementa: validarRegistro(email, password, confirmacion string) error
	//   Valida:
	//   - email no vacío y contiene "@" (strings.Contains)
	//   - password con len >= 8
	//   - confirmacion == password
	//   Si hay errores, retorna *ErrorFormulario con todos los errores.
	//   Si no hay errores, retorna nil.
	//
	// Pruébala con casos que generen 0, 1 y múltiples errores.
	// Usa errors.As para extraer el *ErrorFormulario y mostrar cada campo.

	// TODO
}
