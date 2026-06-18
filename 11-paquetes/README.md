# Módulos, Paquetes e Imports en Go

## Comparación con Node.js / npm

| Concepto | Node.js | Go |
|----------|---------|----|
| Archivo de módulo | `package.json` | `go.mod` |
| Gestor de dependencias | `npm` | `go get` / `go mod tidy` |
| Import | `require()` / `import` | `import "path/paquete"` |
| Exportar | `module.exports` / `export` | Nombre con **Mayúscula** |
| Módulo privado | `_nombre` (convención) | paquete `internal/` |

## Exported vs Unexported — la regla más simple de Go

**Mayúscula inicial = público (exported)**
**minúscula inicial = privado (unexported)**

```go
// Visible desde otros paquetes:
func Calcular() int { ... }
type Usuario struct { ... }
var Config = "valor"

// Solo visible dentro del mismo paquete:
func calcularInterno() int { ... }
type usuario struct { ... }
var config = "valor"
```

No hay `public`, `private`, `protected`. Solo mayúscula vs minúscula.

## Estructura de un módulo

```
mi-proyecto/
├── go.mod              ← define el módulo
├── main.go             ← package main
├── usuario/
│   └── usuario.go      ← package usuario
└── internal/
    └── db/
        └── db.go       ← package db (solo accesible dentro del módulo)
```

## Importar paquetes

```go
import (
    "fmt"                              // stdlib
    "strings"                          // stdlib
    "golang-syntax-lab/11-paquetes/task"  // paquete local del módulo
)
```
