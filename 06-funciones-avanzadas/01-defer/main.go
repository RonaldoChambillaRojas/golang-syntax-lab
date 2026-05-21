package main

import (
	"fmt"
	"os"
)

// =========================================================
// DEFER — ejecución postergada al retorno de la función
// =========================================================
// defer garantiza que una función se ejecute cuando la función
// que la contiene retorne, sin importar si retorna normalmente
// o por un error/panic.
//
// En JavaScript el equivalente más cercano es try/finally:
//   try { ... } finally { cleanup() }
//
// Pero defer es más limpio: lo declaras JUSTO DESPUÉS de adquirir
// el recurso, junto a su liberación — sin necesidad de bloques.

func ejemploBasico() {
	fmt.Println("1: inicio de función")
	defer fmt.Println("3: esto se ejecuta al final (defer)")
	fmt.Println("2: mitad de función")
	// Al retornar, Go ejecuta todos los defers en orden LIFO
}

// --- Múltiples defers: orden LIFO (como una pila) ---
func multipleDefer() {
	fmt.Println("\n=== Múltiples defers ===")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  defer %d\n", i)
	}
	fmt.Println("Fin de función (los defers se ejecutan después)")
}
// Output:
//   Fin de función...
//   defer 3   ← último en declararse, primero en ejecutarse
//   defer 2
//   defer 1

// --- Caso de uso real: liberar recursos ---
func leerArchivo(ruta string) error {
	f, err := os.Open(ruta)
	if err != nil {
		return fmt.Errorf("leerArchivo: %w", err)
	}
	defer f.Close() // garantizado sin importar qué pase debajo

	// simulación de lectura
	info, _ := f.Stat()
	fmt.Printf("Leyendo: %s (%d bytes)\n", info.Name(), info.Size())
	return nil
}

// --- Defer con función anónima para modificar el valor de retorno ---
// Este es un uso avanzado: la función anónima puede acceder a los
// named return values y modificarlos antes de retornar.
func operacionConLog(nombre string) (err error) {
	fmt.Printf("Iniciando: %s\n", nombre)

	defer func() {
		if err != nil {
			fmt.Printf("Operación '%s' falló: %v\n", nombre, err)
		} else {
			fmt.Printf("Operación '%s' exitosa\n", nombre)
		}
	}()

	// Simulamos una operación que puede fallar
	if nombre == "" {
		err = fmt.Errorf("nombre no puede estar vacío")
		return
	}
	return
}

// --- Trampas comunes con defer ---
func trampaDefer() {
	fmt.Println("\n=== Trampa: defer evalúa argumentos AHORA, no después ===")

	x := 0
	// El valor de x se captura en el momento del defer (x=0)
	defer fmt.Println("x en defer (capturado ahora):", x)

	x = 10
	fmt.Println("x al final de la función:", x) // 10
	// Cuando se ejecute el defer, imprimirá 0 (el valor capturado)
}

func main() {
	ejemploBasico()
	multipleDefer()

	fmt.Println("\n=== Lectura de archivo ===")
	leerArchivo("go.mod") // existe en el proyecto
	leerArchivo("inexistente.txt")

	fmt.Println("\n=== Operaciones con log ===")
	operacionConLog("procesarPago")
	operacionConLog("")

	trampaDefer()
}
