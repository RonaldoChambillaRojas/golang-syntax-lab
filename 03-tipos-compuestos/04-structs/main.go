package main

import "fmt"

// =========================================================
// STRUCTS EN GO
// =========================================================
// Un struct es una colección de campos con tipos fijos.
// Es el mecanismo de Go para crear tipos compuestos personalizados.
//
// En TypeScript:
//   interface Persona { nombre: string; edad: number; activo: boolean }
//   type Persona = { nombre: string; edad: number; activo: boolean }
//   class Persona { ... }
//
// En Go, el struct es más simple y directo, sin herencia:

type Persona struct {
	Nombre string
	Edad   int
	Activo bool
}

// Los structs pueden contener otros structs (composición)
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

type Empleado struct {
	Persona            // Embedding: Empleado "hereda" los campos de Persona
	Cargo     string
	Salario   float64
	Direccion Direccion // Composición: campo de tipo Direccion
}

func main() {
	// --- Instanciación con campos nombrados (recomendado) ---
	p1 := Persona{
		Nombre: "Ana García",
		Edad:   28,
		Activo: true,
	}
	fmt.Println("Persona:", p1)
	fmt.Printf("Nombre: %s, Edad: %d\n", p1.Nombre, p1.Edad)

	// --- Instanciación posicional (frágil, evitar en producción) ---
	// Si cambias el orden de los campos, este código se rompe silenciosamente
	p2 := Persona{"Carlos", 35, false}
	fmt.Println("Posicional:", p2)

	// --- Zero value de struct: todos los campos en su zero value ---
	var p3 Persona
	fmt.Printf("Zero: %+v\n", p3) // {Nombre: Edad:0 Activo:false}

	// --- Acceso y modificación de campos ---
	p1.Edad = 29
	fmt.Println("Cumpleaños:", p1.Edad)

	// --- Comparación de structs ---
	// Dos structs son iguales si todos sus campos son iguales
	// (solo funciona si todos los campos son comparables)
	a := Persona{"Ana", 28, true}
	b := Persona{"Ana", 28, true}
	c := Persona{"Ana", 29, true}
	fmt.Println("\na == b:", a == b) // true
	fmt.Println("a == c:", a == c)  // false

	// --- Struct con embedding (composición) ---
	emp := Empleado{
		Persona:  Persona{Nombre: "Luis", Edad: 32, Activo: true},
		Cargo:    "Desarrollador",
		Salario:  75000.00,
		Direccion: Direccion{Calle: "Av. Reforma 100", Ciudad: "CDMX", CP: "06600"},
	}

	// Los campos del embedded struct son accesibles directamente:
	fmt.Printf("\nEmpleado: %s, Cargo: %s\n", emp.Nombre, emp.Cargo) // emp.Persona.Nombre promovido
	fmt.Printf("Ciudad: %s\n", emp.Direccion.Ciudad)

	// =========================================================
	// STRUCTS ANÓNIMOS
	// =========================================================
	// Útiles para datos temporales de una sola vez.
	// En JS sería solo un objeto literal {}. En Go puedes hacer lo mismo:

	config := struct {
		Host string
		Port int
		TLS  bool
	}{
		Host: "api.ejemplo.com",
		Port: 443,
		TLS:  true,
	}
	fmt.Printf("\nConfig: %+v\n", config)

	// Slice de structs anónimos — común en tests y datos de ejemplo
	usuarios := []struct {
		ID     int
		Nombre string
	}{
		{1, "Ana"},
		{2, "Carlos"},
		{3, "María"},
	}
	for _, u := range usuarios {
		fmt.Printf("  Usuario %d: %s\n", u.ID, u.Nombre)
	}

	// =========================================================
	// PUNTERO A STRUCT
	// =========================================================
	// En Go, los structs son VALOR. Para modificar desde una función,
	// necesitas un puntero. (Profundizaremos en sección 7)
	p4 := &Persona{Nombre: "Elena", Edad: 25, Activo: true}
	p4.Edad = 26 // Go dereferencia automáticamente el puntero
	fmt.Printf("\nPuntero: %+v\n", *p4)
}
