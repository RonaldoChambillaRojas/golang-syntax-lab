package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicios: Sentinel Errors y Wrapping ===")

	// EJERCICIO 1 — Básico ─────────────────────────────────────
	//
	// Implementa un sistema de caché simple con sentinel errors:
	//
	//   var ErrCacheMiss = errors.New("cache miss")
	//   var ErrCacheExpired = errors.New("entrada expirada")
	//
	//   type Cache struct {
	//       datos map[string]entrada
	//   }
	//   type entrada struct {
	//       valor     string
	//       expiracion time.Time
	//   }
	//
	//   func NuevoCache() *Cache
	//   func (c *Cache) Set(clave, valor string, ttl time.Duration)
	//   func (c *Cache) Get(clave string) (string, error)
	//     → ErrCacheMiss si la clave no existe
	//     → ErrCacheExpired si la entrada existe pero ya expiró
	//     → (valor, nil) si existe y está vigente
	//
	// Prueba:
	//   cache := NuevoCache()
	//   cache.Set("user:1", "Ana", 100*time.Millisecond)
	//   Get("user:1")  → "Ana", nil
	//   time.Sleep(200ms)
	//   Get("user:1")  → "", ErrCacheExpired
	//   Get("user:9")  → "", ErrCacheMiss
	//
	// Importa "time" y "errors".

	// TODO

	// EJERCICIO 2 — Intermedio ─────────────────────────────────
	//
	// Implementa un sistema de reintentos automáticos:
	//
	//   type ErrTransitorio struct { Motivo string }
	//     → representa un error que puede resolverse reintentando
	//     → implementa Error() string
	//
	//   var ErrPermanente = errors.New("error permanente, no reintentar")
	//
	//   func reintentar(intentos int, fn func() error) error
	//     → ejecuta fn hasta `intentos` veces
	//     → si fn retorna nil: retorna nil
	//     → si fn retorna ErrTransitorio (errors.As): reintenta
	//     → si fn retorna cualquier otro error: retorna inmediatamente (no reintenta)
	//     → si agotó los intentos: retorna fmt.Errorf("agotados %d intentos: %w", intentos, últimoError)
	//
	// Simula con un contador:
	//   intentoActual := 0
	//   fn := func() error {
	//       intentoActual++
	//       if intentoActual < 3 {
	//           return &ErrTransitorio{fmt.Sprintf("intento %d fallido", intentoActual)}
	//       }
	//       return nil  // éxito en el tercer intento
	//   }
	//   err := reintentar(5, fn)  → nil (éxito en intento 3)
	//   err := reintentar(2, fn)  → error (agotó intentos)

	// TODO
}
