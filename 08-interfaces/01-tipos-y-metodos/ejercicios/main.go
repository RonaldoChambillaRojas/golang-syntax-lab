package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Tipos y Métodos ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Crea un tipo Temperatura basado en float64.
	// Agrega estos métodos:
	//   (t Temperatura) Celsius() string    → "XX.XX°C"
	//   (t Temperatura) Fahrenheit() Temperatura → convierte asumiendo que t está en °C
	//   (t Temperatura) Kelvin() Temperatura    → convierte de °C a Kelvin (+273.15)
	//   (t Temperatura) String() string    → igual que Celsius()
	//   (t Temperatura) EsBajoCero() bool  → true si t < 0
	//
	// Conversiones:
	//   F = C * 9/5 + 32
	//   K = C + 273.15
	//
	// Prueba:
	//   temp := Temperatura(100)
	//   fmt.Println(temp)                    → "100.00°C"
	//   fmt.Println(temp.Fahrenheit())       → "212.00°C" (¿debería mostrar °F?)
	//   Nota: decide si Fahrenheit() retorna Temperatura o un tipo propio.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Crea un sistema de semáforo con iota:
	//
	//   type Semaforo int
	//   const (Rojo, Amarillo, Verde Semaforo = iota)
	//
	//   (s Semaforo) String() string
	//     → "🔴 Rojo" / "🟡 Amarillo" / "🟢 Verde"
	//     (usa rune literals o strings normales si prefieres sin emoji)
	//
	//   (s Semaforo) EsSeguroAvanzar() bool
	//     → true solo si es Verde
	//
	//   (s *Semaforo) Siguiente() Semaforo
	//     → Rojo → Verde, Amarillo → Rojo, Verde → Amarillo
	//     → actualiza el estado del semáforo Y retorna el nuevo estado
	//
	// Simula 6 cambios de estado partiendo desde Rojo e imprime
	// el estado y si es seguro avanzar en cada paso.

	// TODO
}
