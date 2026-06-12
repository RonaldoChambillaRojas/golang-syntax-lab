# Generics en Go (desde Go 1.18)

## Comparación con TypeScript Generics

Los generics de Go son muy similares a los de TypeScript, con algunas diferencias de sintaxis:

```typescript
// TypeScript
function primero<T>(arr: T[]): T | undefined {
    return arr[0];
}

function filtrar<T>(arr: T[], pred: (x: T) => boolean): T[] {
    return arr.filter(pred);
}
```

```go
// Go
func primero[T any](arr []T) (T, bool) {
    if len(arr) == 0 {
        var zero T
        return zero, false
    }
    return arr[0], true
}

func filtrar[T any](arr []T, pred func(T) bool) []T {
    resultado := make([]T, 0)
    for _, v := range arr {
        if pred(v) {
            resultado = append(resultado, v)
        }
    }
    return resultado
}
```

## Constraints (restricciones de tipo)

En TypeScript usas `extends`:
```typescript
function mayor<T extends number | string>(a: T, b: T): T { ... }
```

En Go usas constraints (del paquete `constraints` o inline):
```go
type Ordenable interface {
    ~int | ~float64 | ~string
}
func mayor[T Ordenable](a, b T) T { ... }
```

El `~` significa "cualquier tipo cuyo tipo subyacente sea int" — incluye tipos definidos como `type MiInt int`.

## Cuándo usar generics

- **Sí**: colecciones genéricas (Stack, Queue, Set), funciones utilitarias (Map, Filter, Reduce)
- **No**: cuando interfaces normales son suficientes, cuando solo tienes 1-2 tipos concretos
