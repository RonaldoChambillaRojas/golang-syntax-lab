# Punteros en Go

## La diferencia fundamental con JavaScript

En JavaScript **todo objeto y array es una referencia**. Nunca trabajas con punteros explícitamente. En Go, los tipos se dividen en dos categorías:

### Tipos Valor (se copian en asignación)
`int`, `float64`, `bool`, `string`, `array`, `struct`

```go
a := 5
b := a     // b es una COPIA de a
b = 10     // no afecta a a
fmt.Println(a) // 5
```

### Tipos Referencia (comparten la misma memoria)
`slice`, `map`, `channel`, `function`, `interface`, `pointer`

```go
a := []int{1, 2, 3}
b := a           // b apunta al mismo array subyacente
b[0] = 99        // modifica a también
fmt.Println(a)   // [99 2 3]
```

## Qué es un puntero

Un puntero es una variable que almacena la **dirección de memoria** de otra variable.

```
Variable x:     [ 42 ]  ← dirección: 0xc000012345
Puntero p:  [ 0xc000012345 ]  ← apunta a x
```

## Operadores

| Operador | Nombre | Qué hace |
|----------|--------|----------|
| `&x` | Dirección-de | Retorna la dirección de memoria de `x` |
| `*p` | Desreferencia | Accede al valor en la dirección almacenada en `p` |
| `*T` | Tipo puntero | El tipo "puntero a T" |

## ¿Por qué son importantes?

1. **Modificar una variable desde una función** — sin puntero, la función recibe una copia
2. **Evitar copias costosas** — pasar `*BigStruct` es más eficiente que copiar todo el struct
3. **Representar la ausencia de valor** — un `*int` nil significa "sin valor" (vs `int` que siempre tiene un valor)
