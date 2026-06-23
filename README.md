# Golang Syntax Lab

Laboratorio de aprendizaje de Go para desarrolladores con experiencia en **JavaScript/TypeScript**.

## Estructura del curso

| Sección | Tema |
|---------|------|
| `01-introduccion` | ¿Qué es Go?, herramientas esenciales |
| `02-fundamentos` | Variables, tipos, operadores, I/O |
| `03-tipos-compuestos` | Arrays, slices, maps, structs |
| `04-control-flow` | If/else, for, switch |
| `05-funciones` | Declaración, retornos múltiples, errores |
| `06-funciones-avanzadas` | Defer, closures, métodos, recursividad |
| `07-punteros` | Punteros, valor vs referencia, mutabilidad |
| `08-interfaces` | Tipos, interfaces implícitas, embedding |
| `09-generics` | Genéricos, constraints, pipelines |
| `10-errores` | Sentinel errors, wrapping, panic/recover |
| `11-paquetes` | Módulos, paquetes, imports, proyecto TaskManager |
| `12-concurrencia` | Goroutines, channels, select, mutex, context |

## Cómo ejecutar ejemplos

```bash
# Agregar Go al PATH si no está configurado
export PATH=$PATH:/usr/local/go/bin

# Ejecutar un ejemplo
go run ./02-fundamentos/01-variables/

# Ejecutar un ejercicio
go run ./02-fundamentos/01-variables/ejercicios/

# Formatear código (equivalente a Prettier)
go fmt ./...

# Análisis estático (equivalente a ESLint)
go vet ./...
```

## Convención de archivos

Cada tema tiene dos archivos:

```
topic/
├── main.go            → Ejemplos explicados con comentarios
└── ejercicios/
    └── main.go        → Ejercicios para resolver (sin soluciones)
```

Las secciones conceptualmente complejas también tienen un `README.md`.

## Go vs JavaScript/TypeScript — diferencias clave

| Concepto | JavaScript/TypeScript | Go |
|----------|-----------------------|----|
| Tipado | Dinámico (JS) / Estático (TS) | Estático siempre |
| Declaración | `let x = 5` | `x := 5` |
| Nulo | `undefined`, `null` | Zero values (`0`, `""`, `false`, `nil`) |
| Errores | `try/catch/throw` | Retorno explícito `(value, error)` |
| Concurrencia | `async/await`, `Promise` | goroutines + channels |
| Clases | `class` + `extends` | `struct` + métodos + interfaces |
| Herencia | `extends` | Composición (embedding) |
| Genéricos | `<T>` en TypeScript | `[T any]` desde Go 1.18 |
| Módulos | `npm` + `package.json` | `go mod` + `go.mod` |
| Interfaces | Explícitas (`implements`) | **Implícitas** (duck typing) |
