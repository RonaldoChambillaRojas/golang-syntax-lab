# Interfaces en Go

## La diferencia más grande con TypeScript

En TypeScript, implementar una interface es **explícito**:
```typescript
interface Animal {
    hablar(): string;
}
class Perro implements Animal {  // ← "implements" explícito
    hablar() { return "Guau"; }
}
```

En Go, la implementación es **implícita** — si un tipo tiene los métodos que la interface requiere, automáticamente la implementa. Sin `implements`, sin herencia declarada:

```go
type Animal interface {
    Hablar() string
}

type Perro struct{}
func (p Perro) Hablar() string { return "Guau" }
// Perro implementa Animal automáticamente — sin declararlo
```

Esto es **duck typing**: "si camina como un pato y grazna como un pato, es un pato".

## Por qué esto es poderoso

Permite que código de terceros implemente tus interfaces sin modificar su código. Es la base del principio de "Dependency Inversion" en Go.

## Embedding — composición sobre herencia

Go no tiene herencia. En su lugar usa **embedding**:

```go
type Base struct { ID int }
func (b Base) Identificar() string { return fmt.Sprintf("ID: %d", b.ID) }

type Extendido struct {
    Base         // embed: los campos y métodos de Base son promovidos
    Nombre string
}

e := Extendido{Base: Base{ID: 1}, Nombre: "foo"}
e.Identificar()  // accede al método de Base directamente
```
