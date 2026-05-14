package main

import "fmt"

// =========================================================
// MINIPROYECTO: Catálogo de Biblioteca
// =========================================================
// Integra arrays, slices, maps y structs en un proyecto cohesivo.

type Libro struct {
	ISBN      string
	Titulo    string
	Autor     string
	Anio      int
	Prestado  bool
}

type Biblioteca struct {
	Nombre  string
	Libros  []Libro
	Indices map[string]int // ISBN → índice en slice Libros
}

// nuevaBiblioteca inicializa una Biblioteca con su map correctamente.
func nuevaBiblioteca(nombre string) Biblioteca {
	return Biblioteca{
		Nombre:  nombre,
		Libros:  []Libro{},
		Indices: make(map[string]int),
	}
}

// agregarLibro agrega un libro al catálogo.
func agregarLibro(b *Biblioteca, libro Libro) {
	b.Indices[libro.ISBN] = len(b.Libros)
	b.Libros = append(b.Libros, libro)
}

// buscarPorISBN retorna el libro y si fue encontrado.
func buscarPorISBN(b *Biblioteca, isbn string) (Libro, bool) {
	idx, existe := b.Indices[isbn]
	if !existe {
		return Libro{}, false
	}
	return b.Libros[idx], true
}

// prestarLibro marca un libro como prestado si está disponible.
func prestarLibro(b *Biblioteca, isbn string) bool {
	idx, existe := b.Indices[isbn]
	if !existe || b.Libros[idx].Prestado {
		return false
	}
	b.Libros[idx].Prestado = true
	return true
}

// devolverLibro marca un libro como devuelto.
func devolverLibro(b *Biblioteca, isbn string) bool {
	idx, existe := b.Indices[isbn]
	if !existe || !b.Libros[idx].Prestado {
		return false
	}
	b.Libros[idx].Prestado = false
	return true
}

// buscarPorAutor retorna todos los libros de un autor.
func buscarPorAutor(b *Biblioteca, autor string) []Libro {
	var resultado []Libro
	for _, libro := range b.Libros {
		if libro.Autor == autor {
			resultado = append(resultado, libro)
		}
	}
	return resultado
}

func imprimirCatalogo(b *Biblioteca) {
	fmt.Printf("\n📚 Catálogo: %s (%d libros)\n", b.Nombre, len(b.Libros))
	fmt.Println("─────────────────────────────────────────────────")
	for _, libro := range b.Libros {
		estado := "Disponible"
		if libro.Prestado {
			estado = "PRESTADO"
		}
		fmt.Printf("  [%s] %s - %s (%d) [%s]\n",
			libro.ISBN, libro.Titulo, libro.Autor, libro.Anio, estado)
	}
	fmt.Println("─────────────────────────────────────────────────")
}

func main() {
	bib := nuevaBiblioteca("Biblioteca Central")

	// Agregar libros al catálogo
	libros := []Libro{
		{"978-0-13-468599-1", "The Go Programming Language", "Alan Donovan", 2015, false},
		{"978-0-13-110362-7", "The C Programming Language", "Brian Kernighan", 1978, false},
		{"978-0-20-161622-4", "The Pragmatic Programmer", "David Thomas", 1999, false},
		{"978-0-13-235088-4", "Clean Code", "Robert Martin", 2008, false},
		{"978-0-59-651798-1", "Learning Go", "Jon Bodner", 2021, false},
	}
	for _, libro := range libros {
		agregarLibro(&bib, libro)
	}

	imprimirCatalogo(&bib)

	// Operaciones
	fmt.Println("\n--- Operaciones ---")

	isbn := "978-0-13-468599-1"
	if libro, ok := buscarPorISBN(&bib, isbn); ok {
		fmt.Printf("Encontrado: %s\n", libro.Titulo)
	}

	if prestarLibro(&bib, isbn) {
		fmt.Printf("Libro '%s' prestado exitosamente\n", isbn)
	}

	if !prestarLibro(&bib, isbn) {
		fmt.Println("No se puede prestar: ya está prestado")
	}

	if devolverLibro(&bib, isbn) {
		fmt.Printf("Libro '%s' devuelto exitosamente\n", isbn)
	}

	// Buscar por autor
	librosMartin := buscarPorAutor(&bib, "Robert Martin")
	fmt.Printf("\nLibros de Robert Martin: %d\n", len(librosMartin))
	for _, l := range librosMartin {
		fmt.Printf("  - %s\n", l.Titulo)
	}

	imprimirCatalogo(&bib)
}
