package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Panic y Recover ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa: ejecutarConRecover(fn func()) (err error)
	//   → ejecuta fn() de forma segura
	//   → si fn hace panic, intercepta con recover y retorna el error
	//   → si fn es exitosa, retorna nil
	//
	// Prueba llamándola con:
	//   a) Una función que divide por cero (int, no float)
	//   b) Una función que accede a un índice fuera de rango
	//   c) Una función que llama panic("error manual")
	//   d) Una función que ejecuta sin problemas
	//
	// Para cada caso imprime si hubo error o no.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un "middleware" de recuperación para un servidor HTTP simulado:
	//
	//   type Handler func(request string) string
	//
	//   func conRecuperacion(h Handler) Handler
	//     → retorna un nuevo Handler que envuelve h con recover
	//     → si h hace panic, retorna "500 Internal Server Error: <mensaje del panic>"
	//     → si h es exitosa, retorna su resultado normal
	//
	//   func registrarHandler(nombre string, h Handler) Handler
	//     → retorna un Handler que:
	//       1. Imprime "→ [<nombre>] request: <request>"
	//       2. Llama h y captura el resultado
	//       3. Imprime "← [<nombre>] response: <resultado>"
	//       4. Retorna el resultado
	//
	// Crea estos handlers:
	//   saludar: retorna "Hola, <request>!"
	//   vulnerable: si request == "crash", hace panic("handler crashed")
	//               sino retorna "OK: <request>"
	//
	// Compón: registrarHandler("saludar", conRecuperacion(saludar))
	//         registrarHandler("vulnerable", conRecuperacion(vulnerable))
	//
	// Prueba con requests: "mundo", "crash", "Go"

	// TODO
}
