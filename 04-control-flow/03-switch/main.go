package main

import (
	"fmt"
	"time"
)

func main() {
	// =========================================================
	// SWITCH EN GO
	// =========================================================
	// Diferencias clave con JS:
	// 1. NO hay fallthrough por defecto (cada case tiene break implícito)
	// 2. Los cases pueden ser expresiones, no solo constantes
	// 3. Switch sin expresión actúa como if/else if encadenado
	// 4. Puedes poner múltiples valores en un case con coma

	// --- Switch básico ---
	dia := "Lunes"
	switch dia {
	case "Sábado", "Domingo": // múltiples valores en un case
		fmt.Println("Fin de semana")
	case "Lunes", "Martes", "Miércoles", "Jueves", "Viernes":
		fmt.Println("Día laboral")
	default:
		fmt.Println("Día desconocido")
	}

	// --- Switch con int ---
	mes := 3
	switch mes {
	case 12, 1, 2:
		fmt.Println("Invierno")
	case 3, 4, 5:
		fmt.Println("Primavera")
	case 6, 7, 8:
		fmt.Println("Verano")
	case 9, 10, 11:
		fmt.Println("Otoño")
	}

	// --- Switch con sentencia de inicialización ---
	switch hora := time.Now().Hour(); {
	case hora < 12:
		fmt.Println("Buenos días")
	case hora < 18:
		fmt.Println("Buenas tardes")
	default:
		fmt.Println("Buenas noches")
	}

	// =========================================================
	// SWITCH EN BLANCO — equivale a if/else if
	// =========================================================
	// Cuando el switch no tiene expresión, cada case es una condición bool.
	// Más limpio que un if/else if largo.

	temperatura := 28.5
	switch {
	case temperatura < 0:
		fmt.Println("Bajo cero")
	case temperatura < 15:
		fmt.Println("Frío")
	case temperatura < 25:
		fmt.Println("Agradable")
	case temperatura < 35:
		fmt.Println("Caluroso")
	default:
		fmt.Println("Muy caluroso")
	}

	// =========================================================
	// FALLTHROUGH — explícito cuando lo necesitas
	// =========================================================
	// Al contrario de JS (donde el fallthrough es el DEFECTO),
	// en Go debes pedirlo explícitamente con la palabra fallthrough.

	x := 1
	switch x {
	case 1:
		fmt.Println("case 1")
		fallthrough // cae al siguiente caso
	case 2:
		fmt.Println("case 2 (por fallthrough)")
		// NO cae al 3 porque no tiene fallthrough
	case 3:
		fmt.Println("case 3")
	}

	// =========================================================
	// TYPE SWITCH — un feature único de Go
	// =========================================================
	// Permite distinguir el tipo dinámico de una interface{}
	// Muy útil cuando trabajas con interfaces (sec. 8)

	valores := []interface{}{42, "hola", 3.14, true, []int{1, 2}}
	for _, v := range valores {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d\n", t)
		case string:
			fmt.Printf("string: %q\n", t)
		case float64:
			fmt.Printf("float64: %f\n", t)
		case bool:
			fmt.Printf("bool: %t\n", t)
		default:
			fmt.Printf("tipo desconocido: %T\n", t)
		}
	}
}
