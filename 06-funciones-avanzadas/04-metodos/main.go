package main

import (
	"fmt"
	"math"
)

// =========================================================
// MÉTODOS EN GO
// =========================================================
// Un método es una función con un RECEPTOR: un parámetro especial
// que define el tipo al que pertenece el método.
//
// En JavaScript: los métodos son funciones dentro de una clase
//   class Rectangulo {
//     constructor(b, h) { this.base = b; this.altura = h; }
//     area() { return this.base * this.altura; }
//   }
//
// En Go: los métodos se definen FUERA del struct, con un receptor
// No hay class, no hay this — solo structs y funciones con receptor.

type Rectangulo struct {
	Base   float64
	Altura float64
}

// --- Receptor por valor: el método recibe una COPIA del struct ---
// Úsalo cuando el método NO modifica el struct.
// Equivale a que el método trabaje con una snapshot del objeto.
func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Base + r.Altura)
}

func (r Rectangulo) String() string {
	return fmt.Sprintf("Rectángulo(%.2f x %.2f)", r.Base, r.Altura)
}

// --- Receptor por puntero: el método recibe una referencia al struct ---
// Úsalo cuando el método MODIFICA el struct.
// Es el equivalente a los métodos que modifican "this" en JS.
func (r *Rectangulo) Escalar(factor float64) {
	r.Base *= factor   // modifica el struct original
	r.Altura *= factor
}

// =========================================================
// MÉTODOS SOBRE TIPOS PRIMITIVOS
// =========================================================
// No solo los structs pueden tener métodos.
// Cualquier TIPO DEFINIDO en el mismo paquete puede tener métodos.

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", float64(f))
}

// =========================================================
// INTERFACE Stringer — implementación implícita
// =========================================================
// Si tu tipo tiene un método String() string, fmt lo usa automáticamente.
// Es el equivalente de toString() en JavaScript.

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

func main() {
	// --- Método por valor ---
	r := Rectangulo{Base: 5, Altura: 3}
	fmt.Println("Rectángulo:", r)        // usa String()
	fmt.Printf("Área: %.2f\n", r.Area())
	fmt.Printf("Perímetro: %.2f\n", r.Perimetro())

	// --- Método por puntero ---
	r.Escalar(2)
	fmt.Println("Después de Escalar(2):", r) // 10 x 6

	// Nota: aunque r no es un puntero, Go lo pasa automáticamente como &r
	// cuando el método tiene receptor de puntero. Puedes también:
	rPtr := &Rectangulo{Base: 4, Altura: 2}
	rPtr.Escalar(3)
	fmt.Println("Ptr escalado:", *rPtr)

	// --- Tipos primitivos con métodos ---
	temp := Celsius(100)
	fmt.Printf("\n%s = %s\n", temp, temp.ToFahrenheit())

	fahr := Fahrenheit(98.6)
	fmt.Printf("%s = %s\n", fahr, fahr.ToCelsius())

	// --- Círculo ---
	c := Circulo{Radio: 5}
	fmt.Printf("\n%s\n", c)
	fmt.Printf("Área: %.4f\n", c.Area())
	fmt.Printf("Perímetro: %.4f\n", c.Perimetro())

	// =========================================================
	// CUÁNDO USAR RECEPTOR POR VALOR vs PUNTERO
	// =========================================================
	// Por valor (r Tipo):
	//   ✓ El método no necesita modificar el struct
	//   ✓ El struct es pequeño y barato de copiar
	//   ✓ Semántica de valor (como int, float64)
	//
	// Por puntero (*Tipo):
	//   ✓ El método necesita modificar el struct
	//   ✓ El struct es grande (copiar sería costoso)
	//   ✓ Consistencia: si algún método necesita puntero, usa puntero en todos
	fmt.Println("\n=== Regla práctica ===")
	fmt.Println("Si algún método modifica el struct → usa *Tipo para todos")
}
