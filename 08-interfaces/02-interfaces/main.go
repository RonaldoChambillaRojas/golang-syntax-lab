package main

import (
	"fmt"
	"math"
)

// =========================================================
// INTERFACES EN GO — Implementación implícita
// =========================================================
// Una interface define un CONJUNTO DE MÉTODOS.
// Cualquier tipo que tenga esos métodos la implementa automáticamente.
// No hay "implements", no hay herencia declarada.
//
// En TypeScript: interface Shape { area(): number }
// En Go:         type Forma interface { Area() float64 }

// --- Definición de interfaces ---
type Forma interface {
	Area() float64
	Perimetro() float64
}

// Stringer ya está en el paquete fmt, pero aquí lo definimos para ilustrar:
// type Stringer interface { String() string }

// --- Tipos que implementan Forma ---
type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo(r=%.2f)", c.Radio)
}

type Rectangulo struct {
	Base, Altura float64
}

func (r Rectangulo) Area() float64 { return r.Base * r.Altura }

func (r Rectangulo) Perimetro() float64 { return 2 * (r.Base + r.Altura) }

func (r Rectangulo) String() string {
	return fmt.Sprintf("Rect(%.2fx%.2f)", r.Base, r.Altura)
}

type Triangulo struct {
	A, B, C float64 // lados
}

func (t Triangulo) Perimetro() float64 { return t.A + t.B + t.C }

func (t Triangulo) Area() float64 { // Fórmula de Herón
	s := t.Perimetro() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// --- Usando interfaces como parámetro ---
// Esta función acepta CUALQUIER tipo que implemente Forma.
// En TS: function describir(f: Shape)
func describir(f Forma) {
	fmt.Printf("  Área: %.4f, Perímetro: %.4f\n", f.Area(), f.Perimetro())
}

func areaTotal(formas []Forma) float64 {
	total := 0.0
	for _, f := range formas {
		total += f.Area()
	}
	return total
}

// =========================================================
// INTERFACE VACÍA — any (antes interface{})
// =========================================================
// any = interface{} = acepta cualquier tipo
// Úsala cuando realmente no sabes el tipo en compilación.
// Evítala cuando puedas: pierdes type safety.

func imprimirCualquierCosa(v any) {
	fmt.Printf("Tipo: %T, Valor: %v\n", v, v)
}

// =========================================================
// TYPE ASSERTION — convertir de interface a tipo concreto
// =========================================================
func procesarForma(f Forma) {
	// Type assertion: "confío en que f es *Circulo"
	if c, ok := f.(Circulo); ok {
		fmt.Printf("  Es un círculo con radio %.2f\n", c.Radio)
		return
	}
	if r, ok := f.(Rectangulo); ok {
		fmt.Printf("  Es un rectángulo %v x %v\n", r.Base, r.Altura)
		return
	}
	fmt.Println("  Forma desconocida")
}

func main() {
	// --- Polimorfismo con interface ---
	formas := []Forma{
		Circulo{Radio: 5},
		Rectangulo{Base: 4, Altura: 6},
		Triangulo{A: 3, B: 4, C: 5},
	}

	fmt.Println("=== Formas ===")
	for _, f := range formas {
		fmt.Println(f) // usa String() si está disponible (fmt.Stringer)
		describir(f)
	}

	fmt.Printf("\nÁrea total: %.4f\n", areaTotal(formas))

	// --- Type assertions ---
	fmt.Println("\n=== Type assertions ===")
	for _, f := range formas {
		procesarForma(f)
	}

	// Assertion sin ok: si falla → PANIC
	// c := formas[0].(Circulo) ← seguro si sabes el tipo
	// r := formas[0].(Rectangulo) ← PANIC: Circulo is not Rectangulo

	// --- any / interface{} ---
	fmt.Println("\n=== any ===")
	imprimirCualquierCosa(42)
	imprimirCualquierCosa("hola")
	imprimirCualquierCosa(Circulo{3})
	imprimirCualquierCosa([]int{1, 2, 3})

	// --- Interfaces anidadas / composición ---
	// Puedes combinar interfaces con embedding
	type FormaSer interface {
		Forma
		fmt.Stringer // embed de la interface Stringer del paquete fmt
	}

	// Circulo y Rectangulo implementan FormaSer (tienen Area, Perimetro y String)
	var fs FormaSer = Circulo{Radio: 7}
	fmt.Printf("\nFormaSer: %s → área %.4f\n", fs.String(), fs.Area())
}
