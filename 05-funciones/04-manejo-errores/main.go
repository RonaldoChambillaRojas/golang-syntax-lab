package main

import (
	"errors"
	"fmt"
	"os"
)

// =========================================================
// MANEJO DE ERRORES EN GO
// =========================================================
// La diferencia filosófica fundamental con JavaScript:
//
// JavaScript: los errores son excepciones (flujo excepcional)
//   → se lanzan con throw, se capturan con try/catch
//   → si no los capturas, el programa se cae
//
// Go: los errores son VALORES (flujo normal)
//   → se retornan como último valor de retorno
//   → el llamador DECIDE qué hacer con el error
//   → el compilador no te obliga a manejarlos, pero sí a reconocerlos
//
// La interface error es simplemente:
//   type error interface { Error() string }

// --- Creando errores con errors.New ---
var ErrDivisionPorCero = errors.New("división por cero")

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionPorCero
	}
	return a / b, nil
}

// --- Errores con contexto: fmt.Errorf ---
// Como throw new Error(`context: ${message}`) en JS
func abrirArchivo(ruta string) (*os.File, error) {
	f, err := os.Open(ruta)
	if err != nil {
		// Agrega contexto al error original con %w (wrap)
		return nil, fmt.Errorf("abrirArchivo(%q): %w", ruta, err)
	}
	return f, nil
}

// --- Tipo de error personalizado ---
// Para errores que necesitan llevar datos adicionales

type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

// Implementar la interface error
func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("error en campo %q: %s", e.Campo, e.Mensaje)
}

func validarEdad(edad int) error {
	if edad < 0 {
		return &ErrorValidacion{Campo: "edad", Mensaje: "no puede ser negativa"}
	}
	if edad > 150 {
		return &ErrorValidacion{Campo: "edad", Mensaje: "valor fuera de rango"}
	}
	return nil
}

func validarNombre(nombre string) error {
	if nombre == "" {
		return &ErrorValidacion{Campo: "nombre", Mensaje: "no puede estar vacío"}
	}
	if len(nombre) < 2 {
		return &ErrorValidacion{Campo: "nombre", Mensaje: "muy corto (mínimo 2 caracteres)"}
	}
	return nil
}

func main() {
	// --- Patrón estándar: if err != nil ---
	resultado, err := dividir(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("10 / 3 = %.4f\n", resultado)

	// Error case
	_, err = dividir(5, 0)
	if err != nil {
		fmt.Println("Error:", err) // división por cero
	}

	// --- Comparar con error sentinel ---
	if errors.Is(err, ErrDivisionPorCero) {
		fmt.Println("Fue una división por cero específicamente")
	}

	// --- Error con contexto ---
	_, err = abrirArchivo("archivo_inexistente.txt")
	if err != nil {
		fmt.Println("\nError de archivo:", err)
	}

	// --- Errores de tipo personalizado ---
	if err := validarEdad(-5); err != nil {
		fmt.Println("\nValidación edad:", err)

		// Extraer el tipo concreto del error con errors.As
		var errVal *ErrorValidacion
		if errors.As(err, &errVal) {
			fmt.Printf("  Campo: %s\n", errVal.Campo)
			fmt.Printf("  Mensaje: %s\n", errVal.Mensaje)
		}
	}

	if err := validarNombre(""); err != nil {
		fmt.Println("Validación nombre:", err)
	}

	// =========================================================
	// PROPAGACIÓN DE ERRORES — el patrón más común
	// =========================================================
	// En lugar de manejar todos los errores en main,
	// lo más frecuente es propagarlos hacia arriba:
	err = procesarDatos()
	if err != nil {
		fmt.Println("\nError en proceso:", err)
	}
}

// procesarDatos muestra cómo propagas errores hacia arriba.
func procesarDatos() error {
	_, err := dividir(10, 0)
	if err != nil {
		return fmt.Errorf("procesarDatos: %w", err) // envuelve el error con contexto
	}
	return nil
}
