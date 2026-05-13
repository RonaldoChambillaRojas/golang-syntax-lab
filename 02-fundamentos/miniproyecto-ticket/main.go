package main

import "fmt"

// Producto representa un artículo en el ticket de compra.
type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func main() {
	// =========================================================
	// MINIPROYECTO: Ticket de compra con impuestos
	// =========================================================
	// Este mini proyecto integra: variables, tipos, operadores y fmt.
	// No te preocupes si aún no entiendes structs o slices en detalle,
	// los veremos a fondo en la sección 3.

	const tasaImpuesto = 0.16 // 16% IVA

	// Lista de productos (slice de Producto — veremos slices en sec. 3)
	productos := []Producto{
		{Nombre: "Laptop",   Precio: 1200.00, Cantidad: 1},
		{Nombre: "Mouse",    Precio: 25.50,   Cantidad: 2},
		{Nombre: "Teclado",  Precio: 89.99,   Cantidad: 1},
		{Nombre: "Monitor",  Precio: 350.00,  Cantidad: 1},
		{Nombre: "USB-C Hub",Precio: 45.75,   Cantidad: 3},
	}

	// Imprimir encabezado del ticket
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║           TICKET DE COMPRA                   ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Printf("║ %-15s %6s %10s %10s ║\n", "Producto", "Cant.", "P.Unit.", "Subtotal")
	fmt.Println("╠══════════════════════════════════════════════╣")

	// Calcular subtotales e imprimir cada línea
	var subtotal float64
	for _, p := range productos {
		lineaTotal := p.Precio * float64(p.Cantidad)
		subtotal += lineaTotal
		fmt.Printf("║ %-15s %6d %10.2f %10.2f ║\n",
			p.Nombre, p.Cantidad, p.Precio, lineaTotal)
	}

	// Calcular impuesto y total
	impuesto := subtotal * tasaImpuesto
	total := subtotal + impuesto

	// Imprimir resumen
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Printf("║ %-34s %10.2f ║\n", "Subtotal:", subtotal)
	fmt.Printf("║ %-34s %10.2f ║\n", fmt.Sprintf("IVA (%.0f%%):", tasaImpuesto*100), impuesto)
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Printf("║ %-34s %10.2f ║\n", "TOTAL:", total)
	fmt.Println("╚══════════════════════════════════════════════╝")

	// =========================================================
	// Qué conceptos usamos aquí:
	// - const: tasaImpuesto
	// - struct: Producto (sec. 3)
	// - slice: []Producto (sec. 3)
	// - for range: iteración (sec. 4)
	// - operadores: *, +, float64()
	// - fmt.Printf: formato de tabla con verbos %-15s, %10.2f
	// =========================================================
}
