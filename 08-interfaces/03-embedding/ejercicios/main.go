package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Embedding ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Crea una jerarquía de "animales" usando embedding (sin herencia):
	//
	//   type Animal struct {
	//       Nombre string
	//       Edad   int
	//   }
	//   func (a Animal) Dormir() string → "<Nombre> está durmiendo"
	//   func (a Animal) String() string → "Animal{<Nombre>, <Edad> años}"
	//
	//   type Perro struct {
	//       Animal
	//       Raza string
	//   }
	//   func (p Perro) Ladrar() string   → "¡Guau! Soy <Nombre>"
	//   func (p Perro) String() string   → "Perro{<Nombre>, raza=<Raza>}"
	//
	//   type Gato struct {
	//       Animal
	//       Indoor bool
	//   }
	//   func (g Gato) Maullar() string   → "¡Miau! Soy <Nombre>"
	//   func (g Gato) String() string    → "Gato{<Nombre>, indoor=<Indoor>}"
	//
	// Define una interface Mascota:
	//   type Mascota interface {
	//       fmt.Stringer
	//       Dormir() string
	//   }
	//
	// Crea un slice []Mascota con un Perro y un Gato,
	// e itera llamando a Dormir() y String() de cada uno.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un logger con niveles usando embedding de interfaces:
	//
	//   type LoggerBase struct {
	//       Prefix string
	//   }
	//   func (l LoggerBase) formatear(nivel, msg string) string
	//     → "[<Prefix>][<nivel>] <msg>"
	//
	//   type ConsoleLogger struct {
	//       LoggerBase
	//       // imprime a stdout
	//   }
	//   func (c ConsoleLogger) Info(msg string)
	//   func (c ConsoleLogger) Error(msg string)
	//
	//   type MemoryLogger struct {
	//       LoggerBase
	//       Logs []string  // almacena los mensajes en memoria
	//   }
	//   func (m *MemoryLogger) Info(msg string)
	//   func (m *MemoryLogger) Error(msg string)
	//   func (m MemoryLogger) Dump()  → imprime todos los logs almacenados
	//
	// Define una interface Logger:
	//   type Logger interface {
	//       Info(msg string)
	//       Error(msg string)
	//   }
	//
	// Implementa: ejecutarConLog(logger Logger, operacion func() error)
	//   → loggea "Iniciando operación", ejecuta operacion()
	//   → si retorna error: loggea el error
	//   → si no: loggea "Operación completada"
	//
	// Prueba con ConsoleLogger y MemoryLogger.

	// TODO
}
