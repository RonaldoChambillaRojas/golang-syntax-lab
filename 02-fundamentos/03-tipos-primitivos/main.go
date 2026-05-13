package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// =========================================================
	// TIPOS NUMÉRICOS ENTEROS
	// =========================================================
	// En JavaScript solo existe "number" (float64 internamente).
	// En Go los enteros tienen tamaño fijo y están separados de los decimales.

	var i int = 42       // int: tamaño depende de la plataforma (32 o 64 bits)
	var i8 int8 = 127    // -128 a 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	// Versiones sin signo (solo positivos, rango doble)
	var u uint = 42
	var u8 uint8 = 255   // byte es alias de uint8
	var u32 uint32 = 4294967295

	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u32)

	// El tipo int se usa en la mayoría de casos — elige el específico
	// solo cuando el tamaño importa (ej: protocolo de red, archivos binarios)

	// =========================================================
	// TIPOS NUMÉRICOS DECIMALES
	// =========================================================

	var f32 float32 = 3.14  // ~7 dígitos de precisión
	var f64 float64 = 3.14159265358979 // ~15 dígitos de precisión

	// float64 es el predeterminado (como en JavaScript)
	pi := 3.14159 // Go infiere float64

	fmt.Println(f32, f64, pi)
	fmt.Println("float64 min:", math.SmallestNonzeroFloat64)
	fmt.Println("float64 max:", math.MaxFloat64)

	// =========================================================
	// BOOL
	// =========================================================
	// Igual que en JS. Sin coerción: 0, "", nil NO son false en Go.
	// if 0 { } → ERROR en Go. Solo acepta expresiones bool.

	activo := true
	desactivado := false
	fmt.Println(activo, desactivado)
	fmt.Println(!activo) // false

	// =========================================================
	// STRING
	// =========================================================
	// Inmutable, codificado en UTF-8.
	// En JS los strings son UTF-16. En Go son secuencias de bytes UTF-8.

	saludo := "Hola, mundo"
	fmt.Println(saludo)
	fmt.Println("Longitud en bytes:", len(saludo)) // len cuenta bytes, no caracteres

	// Concatenación (igual que JS con +)
	nombre := "Ronaldo"
	mensaje := "Hola, " + nombre
	fmt.Println(mensaje)

	// String multilínea con backtick (raw string)
	// Equivalente a template literals de JS (pero sin interpolación)
	sql := `SELECT *
FROM usuarios
WHERE activo = true`
	fmt.Println(sql)

	// =========================================================
	// BYTE Y RUNE
	// =========================================================
	// byte  → alias de uint8  → representa un byte (ASCII)
	// rune  → alias de int32  → representa un punto de código Unicode (UTF-32)
	// En JS no existe esta distinción explícita

	var b byte = 'A'  // comillas simples para caracteres, no strings
	var r rune = '🚀' // los emojis necesitan rune (más de 1 byte)

	fmt.Printf("byte: %c (%d)\n", b, b)   // A (65)
	fmt.Printf("rune: %c (%d)\n", r, r)   // 🚀 (128640)

	// Iterar sobre un string con range da runes (Unicode correcto)
	// Iterar con índice da bytes (puede cortar caracteres multibyte)
	palabra := "café"
	fmt.Printf("Bytes: %d, Runes: %d\n", len(palabra), len([]rune(palabra)))

	// =========================================================
	// TAMAÑO EN MEMORIA
	// =========================================================
	fmt.Printf("\nTamaños:\n")
	fmt.Printf("int:     %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("int32:   %d bytes\n", unsafe.Sizeof(i32))
	fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(f64))
	fmt.Printf("bool:    %d bytes\n", unsafe.Sizeof(activo))
}
