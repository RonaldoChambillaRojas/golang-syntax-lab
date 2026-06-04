package main

import "fmt"

// =========================================================
// EMBEDDING — Composición sobre herencia
// =========================================================
// Go no tiene herencia de clases. En su lugar, usa embedding:
// un tipo puede "incrustar" otro tipo y sus métodos son promovidos.
//
// En JavaScript/TypeScript:
//   class Animal { hablar() { ... } }
//   class Perro extends Animal { ... }  // herencia
//
// En Go:
//   type Animal struct { ... }
//   func (a Animal) Hablar() { ... }
//
//   type Perro struct {
//       Animal  // embedding: promueve los métodos de Animal
//       Nombre string
//   }
//   // Perro.Hablar() funciona automáticamente

// --- Embedding de struct ---
type Base struct {
	ID        int
	CreadoEn  string
}

func (b Base) Identificar() string {
	return fmt.Sprintf("ID=%d, creado=%s", b.ID, b.CreadoEn)
}

type Usuario struct {
	Base               // embed: ID, CreadoEn, Identificar() promovidos
	Nombre string
	Email  string
}

func (u Usuario) String() string {
	return fmt.Sprintf("Usuario{%s, nombre=%s}", u.Identificar(), u.Nombre)
}

type Articulo struct {
	Base
	Titulo   string
	Contenido string
}

func (a Articulo) String() string {
	return fmt.Sprintf("Articulo{%s, titulo=%s}", a.Identificar(), a.Titulo)
}

// --- Embedding de interface (composición de interfaces) ---
type Lector interface {
	Leer() string
}

type Escritor interface {
	Escribir(s string)
}

// ReadWriter combina Lector y Escritor (igual que io.ReadWriter en stdlib)
type LectorEscritor interface {
	Lector
	Escritor
}

// Buffer implementa LectorEscritor
type Buffer struct {
	datos string
}

func (b *Buffer) Leer() string         { return b.datos }
func (b *Buffer) Escribir(s string)    { b.datos += s }
func (b Buffer) String() string        { return fmt.Sprintf("Buffer(%q)", b.datos) }

// --- Sobrescribir métodos del embedded type ---
type LoggedBase struct {
	Base
	LogActivo bool
}

// Sobrescribe Identificar() de Base
func (l LoggedBase) Identificar() string {
	base := l.Base.Identificar() // llama al método del embedded explícitamente
	if l.LogActivo {
		return "[LOG] " + base
	}
	return base
}

// --- Embedding múltiple ---
type Motor struct{ CV int }
type Ruedas struct{ Cantidad int }
type Coche struct {
	Motor
	Ruedas
	Marca string
}

func (m Motor) Descripcion() string   { return fmt.Sprintf("Motor %dCV", m.CV) }
func (r Ruedas) Descripcion() string  { return fmt.Sprintf("%d ruedas", r.Cantidad) }

func main() {
	// --- Embedding básico ---
	u := Usuario{
		Base:   Base{ID: 1, CreadoEn: "2024-01-15"},
		Nombre: "Ana García",
		Email:  "ana@mail.com",
	}

	fmt.Println(u)
	fmt.Println("Identificar:", u.Identificar()) // método de Base, promovido
	fmt.Println("ID directo:", u.ID)              // campo de Base, promovido
	fmt.Println("ID explícito:", u.Base.ID)       // acceso explícito también funciona

	a := Articulo{
		Base:      Base{ID: 42, CreadoEn: "2024-03-20"},
		Titulo:    "Aprendiendo Go",
		Contenido: "...",
	}
	fmt.Println("\n" + a.String())
	fmt.Println("Identificar:", a.Identificar())

	// --- Interface embedding ---
	var rw LectorEscritor = &Buffer{}
	rw.Escribir("Hola")
	rw.Escribir(", mundo")
	fmt.Println("\nBuffer:", rw.Leer())

	// --- Sobrescribir método del embed ---
	lb := LoggedBase{
		Base:      Base{ID: 99, CreadoEn: "2024-12-01"},
		LogActivo: true,
	}
	fmt.Println("\n" + lb.Identificar())     // usa la versión de LoggedBase
	fmt.Println(lb.Base.Identificar()) // usa la versión de Base directamente

	// --- Embedding múltiple con ambigüedad ---
	coche := Coche{
		Motor:  Motor{CV: 150},
		Ruedas: Ruedas{Cantidad: 4},
		Marca:  "Go Motors",
	}
	fmt.Println("\nCoche:", coche.Marca)
	// coche.Descripcion() → AMBIGUO: tanto Motor como Ruedas tienen Descripcion()
	// Debes ser explícito:
	fmt.Println(coche.Motor.Descripcion())
	fmt.Println(coche.Ruedas.Descripcion())
}
