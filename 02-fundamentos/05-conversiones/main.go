package main

import (
	"fmt"
	"strconv"
)

func main() {
	// =========================================================
	// CONVERSIONES EXPLÍCITAS EN GO
	// =========================================================
	// En JavaScript la coerción es implícita y silenciosa:
	//   5 + "5" → "55"   (número convertido a string automáticamente)
	//   "5" * 2 → 10     (string convertido a número automáticamente)
	//
	// En Go NO hay coerción implícita. SIEMPRE debes convertir explícitamente.
	// Esto evita una familia entera de bugs difíciles de detectar.

	// --- Conversión entre tipos numéricos ---
	var entero int = 42
	decimal := float64(entero) // int → float64
	pequeño := int32(entero)   // int → int32

	fmt.Printf("int: %v → float64: %v\n", entero, decimal)
	fmt.Printf("int: %v → int32: %v\n", entero, pequeño)

	// Cuidado con la pérdida de información al reducir tamaño:
	grande := 1234567890
	var chico int8 = int8(grande) // trunca sin advertencia
	fmt.Printf("int %d → int8 %d (truncado!)\n", grande, chico)

	// --- Conversión float → int (trunca, no redondea) ---
	precio := 19.99
	precioEntero := int(precio) // 19, no 20
	fmt.Printf("float64 %v → int %v (truncado)\n", precio, precioEntero)

	// =========================================================
	// STRING ↔ INT: usa el paquete strconv
	// =========================================================
	// TRAMPA COMÚN: string(65) NO produce "65", produce "A"
	// porque convierte un rune (punto Unicode) a su carácter.
	// Equivale a String.fromCharCode(65) en JS.

	runeTrap := string(65) // "A", no "65"!
	fmt.Println("string(65):", runeTrap) // A

	// Para convertir un número a su representación string:
	numero := 42
	textoNum := strconv.Itoa(numero) // Int to ASCII → "42"
	fmt.Printf("strconv.Itoa(%d) = %q\n", numero, textoNum)

	// O con fmt.Sprintf (más flexible, como template literals de JS):
	textoFmt := fmt.Sprintf("%d", numero)
	fmt.Printf("fmt.Sprintf: %q\n", textoFmt)

	// Para convertir string a int:
	texto := "123"
	valor, err := strconv.Atoi(texto) // puede fallar → devuelve error
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("strconv.Atoi(%q) = %d\n", texto, valor)
	}

	// Qué pasa con un string inválido:
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Error esperado:", err) // strconv.Atoi: parsing "abc": invalid syntax
	}

	// ParseFloat para decimales:
	textoDecimal := "3.14"
	valorDecimal, _ := strconv.ParseFloat(textoDecimal, 64) // 64 = float64
	fmt.Printf("ParseFloat(%q) = %f\n", textoDecimal, valorDecimal)

	// =========================================================
	// CONVERSIÓN DE BYTES Y RUNES
	// =========================================================
	// []byte(string) → slice de bytes del string
	// string([]byte) → string desde slice de bytes
	mensaje := "Hola"
	bytes := []byte(mensaje)
	fmt.Println("Bytes:", bytes) // [72 111 108 97]

	vuelto := string(bytes)
	fmt.Println("De vuelta:", vuelto) // "Hola"

	// =========================================================
	// RESUMEN: ¿Cuándo se necesita strconv?
	// =========================================================
	// Usar strconv cuando: necesitas la representación textual de un número
	// Usar T(valor) cuando: conviertes entre tipos numéricos compatibles
}
