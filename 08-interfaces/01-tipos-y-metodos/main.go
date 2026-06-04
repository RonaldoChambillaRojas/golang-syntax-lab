package main

import (
	"fmt"
	"math"
)

// =========================================================
// TIPOS PERSONALIZADOS Y MÉTODOS
// =========================================================
// En Go puedes crear tipos basados en tipos existentes y agregarles métodos.
// Esto es diferente a las clases de JS/TS: no hay herencia, no hay this.

// --- Tipo basado en string con métodos ---
type Direccion string

func (d Direccion) EsLocal() bool {
	return d == "localhost" || d == "127.0.0.1"
}

func (d Direccion) ConPuerto(puerto int) string {
	return fmt.Sprintf("%s:%d", d, puerto)
}

// --- Enum con iota y método String() ---
// La interface fmt.Stringer es:
//   type Stringer interface { String() string }
// Si tu tipo la implementa, fmt.Println lo usa automáticamente.

type Estado int

const (
	Pendiente Estado = iota
	EnProceso
	Completado
	Cancelado
)

func (e Estado) String() string {
	switch e {
	case Pendiente:
		return "Pendiente"
	case EnProceso:
		return "En Proceso"
	case Completado:
		return "Completado"
	case Cancelado:
		return "Cancelado"
	default:
		return fmt.Sprintf("Estado(%d)", int(e))
	}
}

func (e Estado) EsFinal() bool {
	return e == Completado || e == Cancelado
}

// --- Tipo para dinero que evita errores de float ---
type Centavos int64

func (c Centavos) String() string {
	return fmt.Sprintf("$%d.%02d", c/100, c%100)
}

func (c Centavos) Add(other Centavos) Centavos {
	return c + other
}

func (c Centavos) Multiply(factor float64) Centavos {
	return Centavos(float64(c) * factor)
}

// --- Struct con método context (tiene acceso a todos los campos) ---
type Pedido struct {
	ID       int
	Estado   Estado
	Subtotal Centavos
	Descuento float64 // porcentaje (0.0 - 1.0)
}

func (p *Pedido) Total() Centavos {
	return p.Subtotal.Multiply(1 - p.Descuento)
}

func (p *Pedido) Procesar() error {
	if p.Estado != Pendiente {
		return fmt.Errorf("pedido %d no puede procesarse desde estado %s", p.ID, p.Estado)
	}
	p.Estado = EnProceso
	return nil
}

func (p *Pedido) Completar() error {
	if p.Estado != EnProceso {
		return fmt.Errorf("pedido %d debe estar En Proceso para completar", p.ID)
	}
	p.Estado = Completado
	return nil
}

func (p Pedido) String() string {
	return fmt.Sprintf("Pedido#%d [%s] Subtotal:%s Total:%s",
		p.ID, p.Estado, p.Subtotal, p.Total())
}

// --- Tipos para geometría ---
type Metro float64

func (m Metro) Cuadrados() float64 {
	return float64(m * m)
}

func main() {
	// --- Tipos con métodos ---
	addr := Direccion("localhost")
	fmt.Println("¿Es local?", addr.EsLocal())
	fmt.Println("Con puerto:", addr.ConPuerto(8080))

	fmt.Println()

	// --- Estado enum ---
	e := Pendiente
	fmt.Println("Estado:", e)           // usa String() automáticamente
	fmt.Println("¿Es final?", e.EsFinal())
	e = Completado
	fmt.Println("Estado:", e)
	fmt.Println("¿Es final?", e.EsFinal())

	fmt.Println()

	// --- Dinero seguro ---
	precio := Centavos(1999) // $19.99
	iva := precio.Multiply(0.16)
	total := precio.Add(iva)
	fmt.Printf("Precio: %s, IVA: %s, Total: %s\n", precio, iva, total)

	fmt.Println()

	// --- Pedido con máquina de estados ---
	pedido := &Pedido{
		ID:       42,
		Estado:   Pendiente,
		Subtotal: Centavos(50000), // $500.00
		Descuento: 0.10,
	}
	fmt.Println(pedido)

	if err := pedido.Procesar(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(pedido)

	if err := pedido.Completar(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(pedido)

	// Intentar procesar de nuevo falla:
	if err := pedido.Procesar(); err != nil {
		fmt.Println("Error esperado:", err)
	}

	// --- math.Sqrt con tipo propio ---
	radio := Metro(5)
	area := math.Pi * radio.Cuadrados()
	fmt.Printf("\nÁrea del círculo: %.4f m²\n", area)
}
