package main

import (
	"errors"
	"fmt"
)

// =========================================================
// ERRORES COMO VALORES
// =========================================================
// La interface error es solo:
//   type error interface { Error() string }
// Cualquier tipo que implemente Error() string es un error.

// --- errors.New — el error más simple ---
var ErrVacio = errors.New("el slice está vacío")

func promediar(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, ErrVacio
	}
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total / float64(len(nums)), nil
}

// --- fmt.Errorf — error con formato y contexto ---
// Usa %w para ENVOLVER el error original (preserva la cadena)
// Usa %v para formatear sin envolver

func leerConfig(archivo string) (string, error) {
	if archivo == "" {
		return "", fmt.Errorf("leerConfig: archivo no puede estar vacío")
	}
	if archivo == "inexistente.yaml" {
		return "", fmt.Errorf("leerConfig(%q): %w", archivo, ErrVacio)
	}
	return "contenido", nil
}

// --- Tipo de error personalizado con datos adicionales ---
type ErrorHTTP struct {
	Codigo  int
	Mensaje string
	URL     string
}

func (e *ErrorHTTP) Error() string {
	return fmt.Sprintf("HTTP %d en %s: %s", e.Codigo, e.URL, e.Mensaje)
}

func hacerPeticion(url string) (string, error) {
	if url == "" {
		return "", &ErrorHTTP{Codigo: 400, Mensaje: "URL vacía", URL: url}
	}
	if url == "http://no-existe.com" {
		return "", &ErrorHTTP{Codigo: 404, Mensaje: "No encontrado", URL: url}
	}
	return `{"status": "ok"}`, nil
}

// --- Múltiples errores: validación de formulario ---
type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("%s: %s", e.Campo, e.Mensaje)
}

type ErroresMultiples struct {
	Errores []*ErrorValidacion
}

func (e *ErroresMultiples) Error() string {
	msg := "errores de validación:"
	for _, err := range e.Errores {
		msg += "\n  - " + err.Error()
	}
	return msg
}

func (e *ErroresMultiples) Agregar(campo, mensaje string) {
	e.Errores = append(e.Errores, &ErrorValidacion{Campo: campo, Mensaje: mensaje})
}

func (e *ErroresMultiples) TieneErrores() bool {
	return len(e.Errores) > 0
}

func validarFormulario(nombre, email string, edad int) error {
	errs := &ErroresMultiples{}
	if nombre == "" {
		errs.Agregar("nombre", "es requerido")
	}
	if len(nombre) > 0 && len(nombre) < 2 {
		errs.Agregar("nombre", "muy corto (mínimo 2 caracteres)")
	}
	if email == "" {
		errs.Agregar("email", "es requerido")
	}
	if edad < 0 || edad > 150 {
		errs.Agregar("edad", fmt.Sprintf("fuera de rango: %d", edad))
	}
	if errs.TieneErrores() {
		return errs
	}
	return nil
}

func main() {
	// --- errors.New ---
	if prom, err := promediar([]float64{1, 2, 3, 4, 5}); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Promedio: %.2f\n", prom)
	}

	if _, err := promediar([]float64{}); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("errors.Is(ErrVacio):", errors.Is(err, ErrVacio))
	}

	// --- fmt.Errorf con contexto ---
	fmt.Println()
	_, err := leerConfig("")
	fmt.Println("Error 1:", err)

	_, err = leerConfig("inexistente.yaml")
	fmt.Println("Error 2:", err)
	fmt.Println("Is ErrVacio:", errors.Is(err, ErrVacio)) // true — %w preserva la cadena

	// --- Tipo de error personalizado ---
	fmt.Println()
	_, err = hacerPeticion("http://no-existe.com")
	if err != nil {
		fmt.Println("Error HTTP:", err)

		// errors.As: extrae el tipo concreto del error
		var errHTTP *ErrorHTTP
		if errors.As(err, &errHTTP) {
			fmt.Printf("  Código: %d, URL: %s\n", errHTTP.Codigo, errHTTP.URL)
		}
	}

	// --- Múltiples errores ---
	fmt.Println()
	err = validarFormulario("", "no-es-email", -5)
	if err != nil {
		fmt.Println(err)
	}

	err = validarFormulario("Ana", "ana@mail.com", 28)
	if err == nil {
		fmt.Println("\nFormulario válido!")
	}
}
