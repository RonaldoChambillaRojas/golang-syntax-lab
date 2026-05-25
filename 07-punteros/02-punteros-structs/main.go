package main

import "fmt"

// =========================================================
// PUNTEROS A STRUCTS
// =========================================================

type Usuario struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

func (u *Usuario) Activar() {
	u.Activo = true
}

func (u *Usuario) Desactivar() {
	u.Activo = false
}

func (u *Usuario) CambiarEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email no puede estar vacío")
	}
	u.Email = email
	return nil
}

func (u Usuario) String() string {
	estado := "inactivo"
	if u.Activo {
		estado = "activo"
	}
	return fmt.Sprintf("Usuario{%d, %s, %s, %s}", u.ID, u.Nombre, u.Email, estado)
}

// =========================================================
// SLICES Y MAPS YA SON REFERENCIAS INTERNAS
// =========================================================
// Un slice es una estructura {puntero, len, cap}
// Modificar elementos del slice dentro de una función SÍ afecta al original.
// PERO re-asignar el slice (append, reslice) NO afecta al original
// a menos que uses *[]T

func agregarElemento(slice *[]int, elemento int) {
	*slice = append(*slice, elemento) // necesita puntero para re-asignar
}

func modificarElemento(slice []int, indice, valor int) {
	slice[indice] = valor // NO necesita puntero para modificar elementos
}

// =========================================================
// INTERFACES Y NIL — trampa importante
// =========================================================
// Una variable de interface tiene dos componentes: (tipo, valor)
// Una interface es nil SOLO si ambos son nil.
// Un puntero nil encerrado en una interface NO es nil interface.

type Animal interface {
	Hablar() string
}

type Perro struct{ Nombre string }

func (p *Perro) Hablar() string {
	return "Guau!"
}

func obtenerAnimal(quieres bool) Animal {
	var p *Perro // puntero nil a Perro
	if quieres {
		p = &Perro{Nombre: "Rex"}
	}
	// BUG CLÁSICO: si quieres=false, retorna *Perro(nil) dentro de Animal
	// La interface NO es nil, aunque el puntero sí lo sea
	return p
}

func main() {
	// --- Puntero a struct ---
	u := &Usuario{ID: 1, Nombre: "Ana", Email: "ana@mail.com"}
	fmt.Println("Antes:", u)

	u.Activar() // Go dereferencia automáticamente: (*u).Activar()
	u.CambiarEmail("ana.garcia@mail.com")
	fmt.Println("Después:", u)

	// --- Sin puntero: el método recibe una copia ---
	u2 := Usuario{ID: 2, Nombre: "Carlos", Email: "carlos@mail.com"}
	// u2.Activar() funciona igual porque Go toma &u2 automáticamente
	// cuando el método tiene receptor de puntero
	u2.Activar()
	fmt.Println("\nu2:", u2)

	// --- Comparando valor vs puntero ---
	// Structs por valor: se copian
	original := Usuario{ID: 3, Nombre: "María"}
	copia := original
	copia.Nombre = "Modificado"
	fmt.Printf("\nOriginal: %s, Copia: %s\n", original.Nombre, copia.Nombre)

	// Structs por puntero: comparten
	ptr1 := &Usuario{ID: 4, Nombre: "Elena"}
	ptr2 := ptr1          // ambos apuntan al mismo Usuario
	ptr2.Nombre = "Elena García"
	fmt.Printf("ptr1: %s, ptr2: %s\n", ptr1.Nombre, ptr2.Nombre) // iguales

	// --- Slices: elementos se comparten, re-asignación no ---
	nums := []int{1, 2, 3, 4, 5}
	modificarElemento(nums, 0, 99) // modifica in-place
	fmt.Println("\nDespués de modificar elemento:", nums)

	agregarElemento(&nums, 6) // necesita puntero para append
	fmt.Println("Después de agregar:", nums)

	// --- Interface nil trap ---
	fmt.Println("\n=== Interface nil trap ===")
	a := obtenerAnimal(true)
	if a != nil {
		fmt.Println("Animal:", a.Hablar())
	}

	b := obtenerAnimal(false)
	fmt.Println("b == nil:", b == nil) // false! aunque el *Perro es nil
	// Esto es un bug sutil. La solución es retornar nil explícitamente:
	// return nil  (en lugar de return p cuando p es nil)
}
