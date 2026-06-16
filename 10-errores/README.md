# Manejo de Errores Avanzado en Go

## La filosofía de errores en Go

En Go, los errores son **valores** — se retornan, se inspeccionan, se propagan. No son excepciones que rompen el flujo de ejecución.

```
JavaScript:               Go:
throw Error("oops")  →  return nil, errors.New("oops")
try { } catch { }    →  if err != nil { return err }
```

## Cadena de errores (Error Wrapping)

Desde Go 1.13, puedes **envolver** errores para agregar contexto sin perder el error original:

```go
original := errors.New("conexión rechazada")
envuelto := fmt.Errorf("conectarBD: %w", original)  // %w = wrap

errors.Is(envuelto, original)  // true — busca en la cadena
```

## Sentinel Errors — errores "centinela"

Son errores declarados como variables de paquete que representan condiciones específicas:

```go
var ErrNotFound = errors.New("not found")

// En el llamador:
if errors.Is(err, ErrNotFound) { ... }
```

## Panic — solo para errores irrecuperables

`panic` es como `throw` en JS, pero debe usarse solo cuando el programa no puede continuar:
- Índice fuera de rango
- Desreferencia de nil
- Invariante roto en tiempo de desarrollo

Para errores normales (archivos no encontrados, validaciones, red) → usa error como valor de retorno.
