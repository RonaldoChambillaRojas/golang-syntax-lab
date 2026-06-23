# Goroutines y Channels en Go

## La diferencia fundamental con JavaScript

JavaScript es **single-threaded** con un event loop. La concurrencia se logra con `async/await` y Promises, pero solo una cosa se ejecuta a la vez.

Go tiene **verdadero paralelismo**: múltiples goroutines pueden ejecutarse simultáneamente en distintos núcleos del CPU.

## Goroutines vs async/await

```javascript
// JavaScript: async/await - una a la vez
async function main() {
    const result1 = await fetchData("url1");
    const result2 = await fetchData("url2"); // espera que termine la anterior
    // o con Promise.all:
    const [r1, r2] = await Promise.all([fetchData("url1"), fetchData("url2")]);
}
```

```go
// Go: goroutines - realmente paralelas
func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() { defer wg.Done(); fetchData("url1") }()
    go func() { defer wg.Done(); fetchData("url2") }()
    wg.Wait()
}
```

## Channels vs Promises

```javascript
// JS: Promise resuelve un valor
const promesa = new Promise((resolve) => resolve(42));
const valor = await promesa;
```

```go
// Go: channel transmite valores entre goroutines
ch := make(chan int)
go func() { ch <- 42 }()  // envía
valor := <-ch              // recibe
```

## El modelo de Go: "No comuniques compartiendo memoria; comparte memoria comunicando"

La diferencia con threads de otros lenguajes: las goroutines se comunican principalmente a través de channels, no a través de variables compartidas protegidas por mutexes (aunque los mutexes también existen cuando son necesarios).

## Resumen de herramientas

| Herramienta | Cuándo usar |
|-------------|-------------|
| `go func()` | Ejecutar tarea concurrente |
| `sync.WaitGroup` | Esperar múltiples goroutines |
| `chan T` | Comunicar valores entre goroutines |
| `select` | Esperar múltiples operaciones de canal |
| `sync.Mutex` | Proteger datos compartidos mutables |
| `context.Context` | Cancelación y timeouts |
