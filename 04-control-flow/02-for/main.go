package main

import "fmt"

func main() {
	// =========================================================
	// FOR EN GO — La única estructura de bucle
	// =========================================================
	// Go tiene UN solo keyword de bucle: for
	// Con él replicas for, while, do-while y forEach de JS.

	// --- Forma 1: for clásico ---
	// Idéntico a JS: for (let i = 0; i < 5; i++)
	fmt.Println("=== For clásico ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}

	// --- Forma 2: for como while ---
	// Go no tiene while, pero for sin init y post actúa igual
	fmt.Println("\n=== For como while ===")
	n := 1
	for n < 64 {
		fmt.Printf("  n = %d\n", n)
		n *= 2
	}

	// --- Forma 3: for infinito ---
	// En JS: while(true) o for(;;)
	// En Go: for { }
	// Necesitas break o return para salir
	fmt.Println("\n=== For infinito con break ===")
	count := 0
	for {
		if count >= 3 {
			break // salida del bucle
		}
		fmt.Println("  count:", count)
		count++
	}

	// --- break y continue ---
	fmt.Println("\n=== break y continue ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // salta al siguiente ciclo (como JS)
		}
		if i > 7 {
			break // sale del bucle
		}
		fmt.Println("  impar:", i) // 1, 3, 5, 7
	}

	// =========================================================
	// FOR RANGE — Iteración idiomática
	// =========================================================
	// Equivalente a forEach/for-of de JS, pero más versátil.

	// Range sobre slice: retorna (índice, valor)
	fmt.Println("\n=== Range sobre slice ===")
	lenguajes := []string{"Go", "Rust", "Python", "TS"}
	for i, lang := range lenguajes {
		fmt.Printf("  [%d] %s\n", i, lang)
	}

	// Solo el índice
	for i := range lenguajes {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Solo el valor (ignorar índice con _)
	for _, lang := range lenguajes {
		fmt.Print(lang, " ")
	}
	fmt.Println()

	// Range sobre map: retorna (clave, valor)
	fmt.Println("\n=== Range sobre map ===")
	capitales := map[string]string{
		"México": "CDMX", "España": "Madrid", "Francia": "París",
	}
	for pais, capital := range capitales {
		fmt.Printf("  %s → %s\n", pais, capital)
	}

	// Range sobre string: retorna (índice de byte, rune)
	fmt.Println("\n=== Range sobre string ===")
	for i, r := range "café" {
		fmt.Printf("  índice=%d, rune=%c (%d)\n", i, r, r)
	}
	// Nota: el índice 3 salta a 4 porque 'é' ocupa 2 bytes en UTF-8

	// Range sobre canal (lo veremos en concurrencia)

	// =========================================================
	// LABELS: control de bucles anidados
	// =========================================================
	// En JS no hay labels nativos para break/continue. En Go sí.
	// Útil para salir de múltiples niveles de loops.
	fmt.Println("\n=== Labels ===")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // sale del bucle EXTERNO
			}
			fmt.Printf("  (%d,%d)\n", i, j)
		}
	}
	fmt.Println("Después del bucle etiquetado")
}
