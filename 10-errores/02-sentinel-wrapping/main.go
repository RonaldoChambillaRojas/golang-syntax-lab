package main

import (
	"errors"
	"fmt"
)

// =========================================================
// SENTINEL ERRORS Y WRAPPING
// =========================================================
// Sentinel errors son errores declarados como variables de paquete.
// Permiten comparar errores con == o errors.Is sin depender del mensaje.
//
// En Node.js podrías hacer:
//   class NotFoundError extends Error { constructor() { super("not found") } }
//   if (err instanceof NotFoundError) { ... }
//
// En Go:
//   var ErrNotFound = errors.New("not found")
//   if errors.Is(err, ErrNotFound) { ... }

// --- Sentinel errors de dominio ---
var (
	ErrNotFound     = errors.New("no encontrado")
	ErrUnauthorized = errors.New("no autorizado")
	ErrConflict     = errors.New("conflicto: el recurso ya existe")
	ErrBadRequest   = errors.New("solicitud inválida")
)

// --- Simulación de repositorio ---
type Usuario struct {
	ID    int
	Nombre string
	Admin bool
}

var usuarios = map[int]Usuario{
	1: {1, "Ana", true},
	2: {2, "Carlos", false},
}

func obtenerUsuario(id int) (Usuario, error) {
	u, ok := usuarios[id]
	if !ok {
		return Usuario{}, fmt.Errorf("obtenerUsuario(id=%d): %w", id, ErrNotFound)
	}
	return u, nil
}

func crearUsuario(id int, nombre string) error {
	if nombre == "" {
		return fmt.Errorf("crearUsuario: %w", ErrBadRequest)
	}
	if _, existe := usuarios[id]; existe {
		return fmt.Errorf("crearUsuario(id=%d): %w", id, ErrConflict)
	}
	usuarios[id] = Usuario{id, nombre, false}
	return nil
}

func eliminarUsuario(solicitanteID, objetivoID int) error {
	solicitante, err := obtenerUsuario(solicitanteID)
	if err != nil {
		return fmt.Errorf("eliminarUsuario: verificar solicitante: %w", err)
	}
	if !solicitante.Admin {
		return fmt.Errorf("eliminarUsuario: %w", ErrUnauthorized)
	}
	if _, err := obtenerUsuario(objetivoID); err != nil {
		return fmt.Errorf("eliminarUsuario: objetivo: %w", err)
	}
	delete(usuarios, objetivoID)
	return nil
}

// =========================================================
// errors.Is — compara errores en la cadena
// =========================================================
// errors.Is(err, target) busca en toda la cadena de wrapping.
// Es como instanceof pero recorre la cadena de errores.

// =========================================================
// errors.As — extrae el tipo concreto de la cadena
// =========================================================
// errors.As(err, &target) encuentra el primer error de ese tipo.

type ErrorConCodigo struct {
	Codigo  int
	Mensaje string
	Causa   error
}

func (e *ErrorConCodigo) Error() string {
	return fmt.Sprintf("[%d] %s", e.Codigo, e.Mensaje)
}

func (e *ErrorConCodigo) Unwrap() error {
	return e.Causa // necesario para que errors.Is/As traverseen la cadena
}

func operacionHTTP(codigo int) error {
	var causa error
	switch codigo {
	case 404:
		causa = ErrNotFound
	case 401:
		causa = ErrUnauthorized
	default:
		return nil
	}
	return &ErrorConCodigo{Codigo: codigo, Mensaje: "petición fallida", Causa: causa}
}

func main() {
	// --- Sentinel errors con errors.Is ---
	fmt.Println("=== errors.Is ===")

	_, err := obtenerUsuario(99)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound)) // true
		fmt.Println("Is ErrUnauthorized:", errors.Is(err, ErrUnauthorized)) // false
	}

	err = crearUsuario(1, "duplicado")
	fmt.Println("\nConflicto:", err)
	fmt.Println("Is ErrConflict:", errors.Is(err, ErrConflict))

	err = crearUsuario(3, "") // nombre vacío
	fmt.Println("\nBad request:", err)
	fmt.Println("Is ErrBadRequest:", errors.Is(err, ErrBadRequest))

	// --- Propagación de errores ---
	fmt.Println("\n=== Propagación ===")
	err = eliminarUsuario(2, 1) // Carlos (no admin) intenta eliminar a Ana
	fmt.Println("Error:", err)
	fmt.Println("Is ErrUnauthorized:", errors.Is(err, ErrUnauthorized))

	err = eliminarUsuario(1, 99) // Ana (admin) intenta eliminar id 99 (no existe)
	fmt.Println("\nError:", err)
	fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound)) // true (en la cadena)

	// --- errors.As con Unwrap ---
	fmt.Println("\n=== errors.As ===")
	err = operacionHTTP(404)
	if err != nil {
		fmt.Println("Error:", err)

		var errCodigo *ErrorConCodigo
		if errors.As(err, &errCodigo) {
			fmt.Printf("  Código HTTP: %d\n", errCodigo.Codigo)
		}
		fmt.Println("  Is ErrNotFound:", errors.Is(err, ErrNotFound)) // true por Unwrap
	}

	// --- Creación exitosa después del conflicto ---
	fmt.Println("\n=== Operación exitosa ===")
	if err := crearUsuario(4, "María"); err != nil {
		fmt.Println("Error:", err)
	} else {
		u, _ := obtenerUsuario(4)
		fmt.Println("Creado:", u)
	}
}
