package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Switch ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Calculadora simple con switch:
	// Declara: a := 15.0, b := 4.0, operacion := "/"
	//
	// Usando switch sobre operacion, implementa:
	//   "+"  → suma
	//   "-"  → resta
	//   "*"  → multiplicación
	//   "/"  → división (si b == 0, imprime "División por cero")
	//   "%"  → módulo (convierte a int para usar %)
	//   default → "Operación desconocida"
	//
	// Imprime: "15 / 4 = 3.75"

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Validador de contraseña con switch en blanco:
	// Dada: password := "GoLang2024!"
	//
	// Evalúa estas condiciones en orden (con switch sin expresión):
	//   1. len(password) < 8             → "Demasiado corta"
	//   2. len(password) > 128           → "Demasiado larga"
	//   3. no contiene dígito            → "Debe tener al menos un número"
	//   4. no contiene mayúscula         → "Debe tener al menos una mayúscula"
	//   5. no contiene carácter especial → "Debe tener al menos un especial"
	//   default                          → "Contraseña válida"
	//
	// Para verificar caracteres, itera el string con range y usa:
	//   unicode.IsDigit(r), unicode.IsUpper(r)
	//   Para especiales: !unicode.IsLetter(r) && !unicode.IsDigit(r)
	//
	// Importa "unicode" si lo usas.
	// Prueba con varias contraseñas: "abc", "password", "Password1", "GoLang2024!"

	// TODO
}
