# Funciones en Go

## Lo que más diferencia a Go de JavaScript

### 1. Múltiples valores de retorno

```go
// Go: una función puede retornar dos (o más) valores
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

resultado, err := dividir(10, 3)
```

```javascript
// JavaScript: necesitas un objeto o destructuring
function dividir(a, b) {
    if (b === 0) throw new Error("división por cero");
    return { resultado: a / b, error: null };
}
const { resultado } = dividir(10, 3);
```

### 2. Manejo de errores sin excepciones

```go
// Go: el error es un valor de retorno, no una excepción
resultado, err := operacion()
if err != nil {
    // maneja el error
    return err
}
// usa resultado
```

```javascript
// JavaScript: excepciones con try/catch
try {
    const resultado = operacion();
} catch (err) {
    // maneja el error
}
```

### 3. Funciones como valores de primera clase

```go
// Go: igual que JS, las funciones son valores
suma := func(a, b int) int { return a + b }
aplicar := func(f func(int, int) int, a, b int) int { return f(a, b) }
```

## Temas de esta sección

1. **01-basico** — Declaración, parámetros, tipos de retorno
2. **02-multiples-retornos** — Múltiples valores, named returns, blank identifier
3. **03-variadic** — Parámetros variables `...T`
4. **04-manejo-errores** — Error como valor de retorno, `error` interface
