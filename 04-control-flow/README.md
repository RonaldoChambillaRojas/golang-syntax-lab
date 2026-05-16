# Estructuras de Control en Go

## Diferencias clave con JavaScript

### Solo existe `for` — no hay `while` ni `do-while`

Go usa `for` para todos los tipos de bucle:

```go
for i := 0; i < 10; i++ { }       // for clásico
for condicion { }                   // equivale a while
for { }                             // bucle infinito (equivale a while true)
for i, v := range slice { }         // for-each
```

### `switch` NO tiene fallthrough por defecto

En JavaScript, los `case` caen al siguiente si no hay `break`. En Go es al revés: cada case tiene su propio `break` implícito.

```go
// Go: debes poner `fallthrough` explícitamente si lo quieres
switch x {
case 1:
    fmt.Println("uno")  // se detiene aquí automáticamente
case 2:
    fmt.Println("dos")
}
```

### `if` puede tener una sentencia de inicialización

```go
// La variable v solo existe dentro del bloque if/else
if v, err := operacion(); err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Resultado:", v)
}
// v y err no existen aquí
```

### No existe el operador ternario

```go
// ❌ No existe: resultado := condicion ? "a" : "b"
// ✅ Usa if/else:
resultado := "b"
if condicion {
    resultado = "a"
}
```
