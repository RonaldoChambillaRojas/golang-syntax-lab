# Funciones Avanzadas en Go

## Conceptos de esta sección

### Defer — como `finally` pero más flexible

`defer` ejecuta una función cuando la función que la contiene retorna (sin importar cómo). Se acumulan en pila (LIFO): el último defer declarado es el primero en ejecutarse.

```go
func procesarArchivo(nombre string) error {
    f, err := os.Open(nombre)
    if err != nil { return err }
    defer f.Close()  // se ejecuta cuando procesarArchivo retorne, sin importar qué pase
    // ... usa f
}
```

En JavaScript el equivalente sería `try/finally`, pero defer es más conciso y no requiere anidar bloques.

### Closures — igual que en JavaScript

```go
func contador() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}
c := contador()
c() // 1
c() // 2
```

Esto es **idéntico** al comportamiento en JavaScript. Los closures en Go capturan variables por referencia.

### Métodos — funciones con receptor

```go
// En JavaScript: class Rectangulo { area() { return this.base * this.altura } }
// En Go: sin class, el método se define fuera del struct
type Rectangulo struct { Base, Altura float64 }

func (r Rectangulo) Area() float64 {  // r es el receptor (como "this")
    return r.Base * r.Altura
}
```

### Recursividad

Go soporta recursividad. **No** optimiza tail-call (diferente a algunos lenguajes funcionales), así que bucles profundos pueden causar stack overflow.
