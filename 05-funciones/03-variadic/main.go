package main

import "fmt"

// =========================================================
// FUNCIONES VARIÁDICAS — Cantidad variable de argumentos
// =========================================================
// En JavaScript: function suma(...nums) { }  (rest parameters)
// En Go: func suma(nums ...int)              (mismo concepto)
//
// Dentro de la función, el parámetro variádico es un SLICE del tipo.

// --- Función variádica simple ---
func sumar(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

// --- Parámetros normales + variádico (el variádico SIEMPRE es el último) ---
func imprimir(prefijo string, valores ...interface{}) {
	fmt.Print(prefijo, ": ")
	for i, v := range valores {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

// --- Variádico con tipo personalizado ---
type Opcion func(*Configuracion)

type Configuracion struct {
	Host    string
	Puerto  int
	Timeout int
	Debug   bool
}

// Patrón "Functional Options" — muy común en librerías Go
// En JS sería: function crear({ host, puerto } = {})
func nuevaConexion(opciones ...Opcion) *Configuracion {
	config := &Configuracion{
		Host:    "localhost",
		Puerto:  8080,
		Timeout: 30,
	}
	for _, opcion := range opciones {
		opcion(config)
	}
	return config
}

func conHost(host string) Opcion {
	return func(c *Configuracion) { c.Host = host }
}

func conPuerto(puerto int) Opcion {
	return func(c *Configuracion) { c.Puerto = puerto }
}

func conDebug() Opcion {
	return func(c *Configuracion) { c.Debug = true }
}

func main() {
	// --- Llamar función variádica ---
	fmt.Println(sumar())           // 0
	fmt.Println(sumar(1))          // 1
	fmt.Println(sumar(1, 2, 3))    // 6
	fmt.Println(sumar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	imprimir("INFO", "servidor iniciado", 42, true)

	// =========================================================
	// EXPANDIR UN SLICE CON ... (spread)
	// =========================================================
	// En JS: sumar(...miArray)
	// En Go: sumar(miSlice...)
	//
	// Cuando ya tienes un slice y quieres pasarlo a una función variádica:

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Suma del slice:", sumar(numeros...)) // spread

	// Combinar slices con append también usa ...:
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	combinado := append(a, b...) // b... expande el slice
	fmt.Println("Combinado:", combinado)

	// =========================================================
	// FUNCTIONAL OPTIONS PATTERN
	// =========================================================
	// Permite llamadas flexibles sin tener que pasar nil ni structs opcionales

	// Sin opciones → defaults
	conn1 := nuevaConexion()
	fmt.Printf("\nConn1: %+v\n", *conn1)

	// Con opciones selectivas
	conn2 := nuevaConexion(
		conHost("produccion.api.com"),
		conPuerto(443),
		conDebug(),
	)
	fmt.Printf("Conn2: %+v\n", *conn2)

	// =========================================================
	// TRASPASO DE VARIÁDICOS entre funciones
	// =========================================================
	wrapper(1, 2, 3, 4)
}

func wrapper(nums ...int) {
	fmt.Println("\nWrapper recibió:", nums)
	// Para pasar el variádico a otra función variádica, usa ...
	fmt.Println("Suma total:", sumar(nums...))
}
