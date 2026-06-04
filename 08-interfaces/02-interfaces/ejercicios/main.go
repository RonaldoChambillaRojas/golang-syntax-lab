package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Interfaces ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Define una interface Notificador:
	//   type Notificador interface {
	//       Enviar(mensaje string) error
	//       Nombre() string
	//   }
	//
	// Implementa tres tipos que la satisfagan:
	//
	//   type Email struct { Destinatario string }
	//     Enviar → imprime "📧 Email a <Destinatario>: <mensaje>" y retorna nil
	//     Nombre → "Email"
	//
	//   type SMS struct { Telefono string }
	//     Enviar → imprime "📱 SMS a <Telefono>: <mensaje>" y retorna nil
	//     Nombre → "SMS"
	//
	//   type Push struct { DispositivoID string; Activo bool }
	//     Enviar → si !Activo, retorna error "dispositivo inactivo"
	//              si Activo, imprime "🔔 Push a <DispositivoID>: <mensaje>"
	//     Nombre → "Push"
	//
	// Implementa: notificarTodos(ns []Notificador, mensaje string)
	//   → envía a todos, imprime éxito o el error para cada uno
	//
	// Prueba con los tres tipos, incluyendo un Push inactivo.

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un sistema de plugins con interface:
	//
	//   type Procesador interface {
	//       Procesar(entrada string) (string, error)
	//       Nombre() string
	//   }
	//
	// Crea tres procesadores:
	//   MayusculasProc: convierte a mayúsculas
	//   TrimProc: elimina espacios al inicio y final (strings.TrimSpace)
	//   LongitudProc: retorna "<entrada> (len=N)"
	//
	// Implementa: cadena(entrada string, procs ...Procesador) (string, error)
	//   → aplica los procesadores en secuencia
	//   → si alguno retorna error, detiene y retorna ese error
	//   → retorna el resultado final
	//
	// Prueba:
	//   cadena("  hola mundo  ", TrimProc{}, MayusculasProc{}, LongitudProc{})
	//   Esperado: "HOLA MUNDO (len=10)"

	// TODO
}
