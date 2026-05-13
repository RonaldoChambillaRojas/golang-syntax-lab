# Introducción a Go

## ¿Qué es Go?

Go (también llamado Golang) es un lenguaje **compilado**, **tipado estáticamente** y **con garbage collector**, creado por Google en 2009. Diseñado para ser simple, eficiente y con concurrencia nativa.

### Lo que lo hace diferente a JavaScript

| Característica | JavaScript | Go |
|---------------|------------|----|
| Ejecución | Interpretado (V8/Node.js) | Compilado a binario nativo |
| Tipado | Dinámico | Estático en compilación |
| Errores | Excepciones (`throw`) | Valores de retorno |
| Concurrencia | Event loop + Promises | Goroutines + Channels |
| Runtime | Necesita Node.js / browser | El binario es autosuficiente |
| null | `undefined` y `null` | Solo `nil` (y tiene zero values) |

## Primer programa

Crea un archivo `hola.go`:

```go
package main  // todo ejecutable necesita el paquete "main"

import "fmt"  // fmt = format, la librería de I/O estándar

func main() { // el punto de entrada, como en C o Java
    fmt.Println("Hola, Go!")
}
```

Ejecútalo:

```bash
/usr/local/go/bin/go run hola.go
```

## Herramientas esenciales

| Comando | Qué hace | Equivalente JS |
|---------|----------|----------------|
| `go run ./carpeta/` | Compila y ejecuta | `node index.js` |
| `go build ./...` | Compila a binario | `tsc` |
| `go fmt ./...` | Formatea el código | Prettier |
| `go vet ./...` | Análisis estático | ESLint |
| `go test ./...` | Ejecuta tests | Jest/Vitest |
| `go mod init <nombre>` | Inicializa módulo | `npm init` |
| `go get <paquete>` | Agrega dependencia | `npm install` |
| `go mod tidy` | Limpia dependencias | `npm prune` |

## Configurar PATH

Go está instalado en `/usr/local/go/bin/go`. Para no escribir la ruta completa:

```bash
# Agregar a ~/.zshrc o ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
```

Luego: `source ~/.zshrc` y ya puedes usar `go` directamente.

## Reglas fundamentales que diferencian a Go de JS

### 1. Las variables declaradas DEBEN usarse
```go
x := 5
// Si no usas x → error de compilación
// En JS solo es un warning (o nada)
```

### 2. Los imports no usados son error
```go
import "fmt"
// Si no usas fmt → error de compilación
// En JS/TS los imports no usados son warnings
```

### 3. No hay conversión implícita de tipos
```go
var x int = 5
var y float64 = x   // ERROR: cannot use x (type int) as type float64
var y float64 = float64(x) // ✓ Conversión explícita requerida
```

### 4. No hay ternario
```go
// En JS: const result = condicion ? "si" : "no"
// En Go debes usar if/else:
result := "no"
if condicion {
    result = "si"
}
```

### 5. Go fmt es el estándar
No hay debates de tabs vs spaces ni de estilos. `go fmt` formatea automáticamente y todos los proyectos Go siguen el mismo estilo.
