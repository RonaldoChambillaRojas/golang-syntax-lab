package main

import "fmt"

// =========================================================
// CONSTANTES A NIVEL DE PAQUETE
// =========================================================
// En JavaScript: const PI = 3.14159 (pero los objetos/arrays siguen mutables)
// En Go: const es VERDADERAMENTE inmutable para cualquier tipo primitivo
const Pi = 3.14159265358979

// Bloque const — igual que bloque var pero para constantes
const (
	MaxReintentos = 3
	Timeout       = 30 // segundos
	AppVersion    = "1.0.0"
)

// =========================================================
// IOTA — Enumeraciones automáticas
// =========================================================
// iota es un contador que Go reinicia a 0 en cada bloque const
// y se incrementa en 1 por cada constante del bloque.
// Equivalente a TypeScript enums, pero sin keyword dedicada.
type DiaSemana int // tipo personalizado basado en int

const (
	Lunes    DiaSemana = iota + 1 // iota=0, +1 → 1
	Martes                        // iota=1, +1 → 2
	Miercoles                     // iota=2, +1 → 3
	Jueves                        // iota=3, +1 → 4
	Viernes                       // iota=4, +1 → 5
	Sabado                        // iota=5, +1 → 6
	Domingo                       // iota=6, +1 → 7
)

// iota con potencias de 2 — muy usado para flags de bits
type Permiso int

const (
	Lectura   Permiso = 1 << iota // 1 << 0 = 1
	Escritura                     // 1 << 1 = 2
	Ejecucion                     // 1 << 2 = 4
)

func main() {
	// =========================================================
	// CONSTANTES TIPADAS Y NO TIPADAS
	// =========================================================

	// Constante NO tipada: el tipo se determina en el punto de uso
	const factor = 2 // puede usarse como int, float64, etc. según contexto

	var entero int = 10 * factor    // factor actúa como int
	var decimal float64 = 1.5 * factor // factor actúa como float64

	fmt.Println(entero, decimal)

	// Constante TIPADA: el tipo está fijo y genera error si hay desajuste
	const limite int = 100
	// var x float64 = limite // ✗ ERROR: cannot use limite (type int) as float64
	_ = limite

	// =========================================================
	// USO PRÁCTICO DE CONSTANTES
	// =========================================================
	fmt.Println("Pi:", Pi)
	fmt.Println("Versión:", AppVersion)
	fmt.Println("Máx reintentos:", MaxReintentos)

	// =========================================================
	// USO DE IOTA
	// =========================================================
	hoy := Miercoles
	fmt.Println("Día número:", hoy) // 3

	// Verificar permisos con flags de bits
	misPermisos := Lectura | Escritura // combinar con OR bit a bit
	fmt.Println("Permiso de lectura:", misPermisos&Lectura != 0)   // true
	fmt.Println("Permiso de ejecución:", misPermisos&Ejecucion != 0) // false

	// =========================================================
	// LAS CONSTANTES NO PUEDEN CAMBIAR (como se esperaría)
	// =========================================================
	// Pi = 3 // ✗ ERROR: cannot assign to Pi (declared const)
}
