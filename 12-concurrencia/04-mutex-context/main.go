package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// =========================================================
// MUTEX — Proteger datos compartidos
// =========================================================
// Cuando múltiples goroutines necesitan leer/escribir los mismos datos,
// debes protegerlos con un Mutex (MUTual EXclusion).

type ContadorSeguro struct {
	mu    sync.Mutex
	valor int64
}

func (c *ContadorSeguro) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *ContadorSeguro) Valor() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

// --- RWMutex: múltiples lectores, un solo escritor ---
// Más eficiente cuando hay muchas lecturas y pocas escrituras
type Cache struct {
	mu    sync.RWMutex
	datos map[string]string
}

func NewCache() *Cache {
	return &Cache{datos: make(map[string]string)}
}

func (c *Cache) Set(clave, valor string) {
	c.mu.Lock() // escritura: lock exclusivo
	defer c.mu.Unlock()
	c.datos[clave] = valor
}

func (c *Cache) Get(clave string) (string, bool) {
	c.mu.RLock() // lectura: múltiples lectores simultáneos
	defer c.mu.RUnlock()
	v, ok := c.datos[clave]
	return v, ok
}

// =========================================================
// ATOMIC — Operaciones atómicas para valores simples
// =========================================================
// Para contadores y flags simples, sync/atomic es más eficiente que Mutex

func demoAtomic() {
	var contador atomic.Int64

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Add(1) // operación atómica, sin mutex
		}()
	}
	wg.Wait()
	fmt.Println("Contador atómico:", contador.Load()) // siempre 1000
}

// =========================================================
// CONTEXT — Cancelación y timeouts
// =========================================================
// context.Context propaga cancelación, timeouts y valores a través
// de la cadena de llamadas. Es el estándar de Go para control de flujo.
//
// En JS serías similar a: AbortController / AbortSignal

func operacionConContext(ctx context.Context, nombre string) error {
	fmt.Printf("  %s: iniciando\n", nombre)

	select {
	case <-time.After(500 * time.Millisecond): // simula trabajo de 500ms
		fmt.Printf("  %s: completado\n", nombre)
		return nil
	case <-ctx.Done(): // si el contexto se cancela/expira
		fmt.Printf("  %s: cancelado - %v\n", nombre, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// --- Mutex básico ---
	fmt.Println("=== Mutex ===")
	contador := &ContadorSeguro{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Println("Contador seguro:", contador.Valor())

	// --- RWMutex (Cache) ---
	fmt.Println("\n=== RWMutex ===")
	cache := NewCache()
	cache.Set("usuario:1", "Ana")
	cache.Set("usuario:2", "Carlos")

	// Múltiples lecturas concurrentes
	var lecturas sync.WaitGroup
	for i := 1; i <= 3; i++ {
		lecturas.Add(1)
		go func(id int) {
			defer lecturas.Done()
			if v, ok := cache.Get(fmt.Sprintf("usuario:%d", id)); ok {
				fmt.Printf("  Lector %d: %s\n", id, v)
			}
		}(i)
	}
	lecturas.Wait()

	// --- Atomic ---
	fmt.Println("\n=== Atomic ===")
	demoAtomic()

	// --- Context con timeout ---
	fmt.Println("\n=== Context con timeout ===")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	err := operacionConContext(ctx, "peticion-lenta")
	if err != nil {
		fmt.Println("Error:", err) // context deadline exceeded
	}

	// --- Context con cancelación manual ---
	fmt.Println("\n=== Context cancelación manual ===")
	ctx2, cancel2 := context.WithCancel(context.Background())

	var wg2 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			operacionConContext(ctx2, fmt.Sprintf("tarea-%d", id))
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	cancel2() // cancela todas las operaciones
	wg2.Wait()
	fmt.Println("Todas las tareas canceladas")

	// --- Context con valor (para pasar datos transversales) ---
	fmt.Println("\n=== Context con valor ===")
	type claveCtx string
	ctx3 := context.WithValue(context.Background(), claveCtx("requestID"), "req-abc-123")

	procesarRequest(ctx3)
}

func procesarRequest(ctx context.Context) {
	type claveCtx string
	if id, ok := ctx.Value(claveCtx("requestID")).(string); ok {
		fmt.Printf("Procesando request: %s\n", id)
	}
}
