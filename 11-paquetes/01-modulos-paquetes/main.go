package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	// Import con alias: útil cuando hay colisiones de nombres o nombres muy largos
	// mathrand "math/rand"  ← así nombrarías un alias
)

// =========================================================
// MÓDULOS Y PAQUETES
// =========================================================
// Un MÓDULO es una colección de paquetes relacionados (definido por go.mod).
// Un PAQUETE es una colección de archivos .go en el mismo directorio.
//
// Cada archivo .go pertenece exactamente a UN paquete.
// El nombre del paquete está en la primera línea: package nombre
//
// Convenciones:
// - El nombre del paquete == el nombre del directorio (casi siempre)
// - Nombres en minúscula, cortos, sin guiones ni underscore
// - package main → produce un ejecutable
// - package xxx  → produce una librería importable

// =========================================================
// EXPORTED vs UNEXPORTED
// =========================================================
// Mayúscula inicial = accesible desde otros paquetes (exported)
// minúscula inicial = solo dentro del mismo paquete (unexported)

// Exported: visible desde otros paquetes
type Servicio struct {
	Nombre  string // exported
	version string // unexported: solo accesible dentro del paquete main
}

// NewServicio es la forma idiomática de crear un tipo (constructor)
func NewServicio(nombre, version string) *Servicio {
	return &Servicio{
		Nombre:  nombre,
		version: version,
	}
}

func (s *Servicio) Version() string { // Exported method
	return s.version // puede acceder al campo unexported
}

func (s *Servicio) String() string {
	return fmt.Sprintf("%s@%s", s.Nombre, s.version)
}

// unexported function: helper interno
func generarID() string {
	return fmt.Sprintf("id_%d", rand.Intn(10000))
}

// =========================================================
// IMPORT PATHS
// =========================================================

func main() {
	// La stdlib de Go tiene ~150 paquetes. Los más usados:
	// fmt     → I/O
	// strings → manipulación de strings
	// strconv → conversiones
	// math    → funciones matemáticas
	// time    → tiempo y duración
	// os      → sistema operativo
	// io      → interfaces I/O
	// net/http → HTTP cliente y servidor
	// encoding/json → JSON
	// sort    → ordenamiento
	// sync    → primitivas de concurrencia
	// errors  → manejo de errores
	// context → cancelación y timeouts

	s := NewServicio("MiAPI", "1.2.3")
	fmt.Println("Servicio:", s)
	fmt.Println("Versión:", s.Version())
	// s.version = "hack" // ✗ ERROR: version is unexported

	// --- Usando paquetes de la stdlib ---
	ahora := time.Now()
	fmt.Printf("\nFecha: %s\n", ahora.Format("2006-01-02 15:04:05"))
	// Nota: Go usa la fecha de referencia "Mon Jan 2 15:04:05 MST 2006"
	// para formatear — es un diseño peculiar de Go

	palabras := strings.Fields("  hola   mundo   go  ")
	fmt.Println("Fields:", palabras)
	fmt.Println("Join:", strings.Join(palabras, "-"))
	fmt.Println("Contains:", strings.Contains("golang", "go"))
	fmt.Println("Replace:", strings.ReplaceAll("a.b.c.d", ".", "/"))

	id := generarID()
	fmt.Println("\nID generado:", id)

	// --- Import con alias ---
	// import rand "math/rand"  → rand.Intn(100)
	// import . "math"          → Sqrt(4) directamente (raro, evitar)
	// import _ "driver/mysql"  → solo efectos secundarios (init())

	// El alias más útil: cuando dos paquetes tienen el mismo nombre
	// import (
	//     nethttp "net/http"
	//     myhttp "mi-proyecto/internal/http"
	// )

	fmt.Println("\n=== Regla de exportación ===")
	fmt.Println("Mayúscula → exported (público)")
	fmt.Println("minúscula → unexported (privado al paquete)")
}
