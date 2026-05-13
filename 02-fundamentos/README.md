# Fundamentos de Go

## Cambios de mentalidad para desarrolladores JS/TS

### Variables: el tipo es permanente

```javascript
// JavaScript: el tipo puede cambiar
let x = 5;
x = "hola"; // válido ✓
```

```go
// Go: el tipo se fija en la declaración
x := 5      // x es int
x = "hola"  // ERROR de compilación ✗
```

### No existe `undefined`

En JavaScript, una variable declarada sin valor es `undefined`. En Go cada tipo tiene un **zero value** — un valor por defecto que nunca es "indefinido":

| Tipo | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` (cadena vacía) |
| `bool` | `false` |
| `pointer`, `slice`, `map`, `interface` | `nil` |

### No hay coerción implícita de tipos

```javascript
// JavaScript
5 + "5"   // "55" (coerción silenciosa)
5 == "5"  // true (coerción silenciosa)
```

```go
// Go
5 + "5"  // ERROR: mismatched types int and string
```

## Temas de esta sección

1. **01-variables** — `var`, `:=`, múltiples variables, swap
2. **02-constantes** — `const`, `iota`, constantes tipadas/no tipadas
3. **03-tipos-primitivos** — `int`, `float64`, `string`, `bool`, `rune`, `byte`
4. **04-zero-values** — valores por defecto de cada tipo
5. **05-conversiones** — conversiones explícitas, `strconv`
6. **06-operadores** — aritmética, comparación, lógica
7. **07-entrada-salida** — `fmt.Println`, `fmt.Printf`, `fmt.Scan`
8. **miniproyecto-ticket** — ticket de compra con impuestos

## Cómo ejecutar

```bash
go run ./02-fundamentos/01-variables/
go run ./02-fundamentos/01-variables/ejercicios/
```
