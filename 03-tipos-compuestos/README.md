# Tipos Compuestos en Go

## Arrays, Slices, Maps y Structs

Go tiene 4 tipos compuestos fundamentales. La diferencia más importante con JavaScript:

### Arrays — Valor fijo, copiados en asignación

```go
// Go: [3]int es un tipo diferente a [4]int
var a [3]int = [3]int{1, 2, 3}
b := a           // b es una COPIA de a, no una referencia
b[0] = 99        // no afecta a a
```

```javascript
// JavaScript: los arrays son siempre referencias
const a = [1, 2, 3];
const b = a;     // b apunta al mismo array
b[0] = 99;       // TAMBIÉN cambia a[0]
```

### Slices — El array dinámico de Go

Los slices son similares a los arrays de JavaScript: dinámicos, con `append`. Son una **vista** sobre un array subyacente.

```go
s := []int{1, 2, 3}
s = append(s, 4)  // como push()
```

### Maps — Diccionarios clave-valor

Equivalente a los objetos `{}` de JavaScript o `Map` de ES6, pero:
- El tipo de la clave debe ser comparable (`==`)
- El orden de iteración es **aleatorio** (intencional)
- Acceder a una clave inexistente retorna el zero value, no `undefined`

### Structs — Tipos personalizados con campos tipados

Son la alternativa Go a los objetos con forma fija de JavaScript/TypeScript. No son clases, no tienen herencia, pero sí pueden tener métodos.

## Resumen de sintaxis

| Tipo | JS equivalente | Declaración Go |
|------|----------------|----------------|
| Array | Array fijo | `[3]int{1,2,3}` |
| Slice | Array dinámico | `[]int{1,2,3}` |
| Map | Object / Map | `map[string]int{"a": 1}` |
| Struct | Object tipado | `struct{ X, Y int }` |
