// Command taskmanager es la aplicación de gestión de tareas.
// Demuestra el uso de múltiples paquetes dentro de un mismo módulo.
package main

import (
	"errors"
	"fmt"

	"golang-syntax-lab/11-paquetes/proyecto-taskmanager/store"
	"golang-syntax-lab/11-paquetes/proyecto-taskmanager/task"
)

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║         TASK MANAGER             ║")
	fmt.Println("╚══════════════════════════════════╝")

	s := store.New()

	// --- Crear tareas ---
	fmt.Println("\n=== Creando tareas ===")
	tareas := []struct {
		titulo, desc string
		prioridad    task.Prioridad
	}{
		{"Implementar API REST", "Endpoints CRUD para usuarios", task.Alta},
		{"Escribir tests", "Cobertura mínima 80%", task.Media},
		{"Deploy a producción", "CI/CD pipeline", task.Critica},
		{"Documentar API", "Swagger/OpenAPI", task.Baja},
		{"Code review", "Revisar PRs pendientes", task.Media},
	}

	for _, td := range tareas {
		t, err := s.Crear(td.titulo, td.desc, td.prioridad)
		if err != nil {
			fmt.Println("Error al crear:", err)
			continue
		}
		fmt.Println("  Creada:", t)
	}

	// --- Intentar crear con título vacío ---
	_, err := s.Crear("", "sin título", task.Baja)
	if errors.Is(err, task.ErrTituloVacio) {
		fmt.Println("\nError esperado:", err)
	}

	// --- Listar todas ---
	fmt.Println("\n=== Todas las tareas ===")
	for _, t := range s.Listar() {
		fmt.Println(" ", t)
	}

	// --- Transiciones de estado ---
	fmt.Println("\n=== Transiciones de estado ===")

	t1, _ := s.ObtenerPorID(1)
	if err := t1.Iniciar(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Iniciada:", t1)
	}

	if err := t1.Completar(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Completada:", t1)
	}

	// Transición inválida: completar lo que ya está completado
	if err := t1.Completar(); err != nil {
		fmt.Println("Error esperado:", err)
		fmt.Println("Is ErrTransicionInvalida:", errors.Is(err, task.ErrTransicionInvalida))
	}

	// Cancelar la tarea 2
	t2, _ := s.ObtenerPorID(2)
	t2.Cancelar()
	fmt.Println("Cancelada:", t2)

	// --- Filtrar por estado ---
	fmt.Println("\n=== Pendientes ===")
	for _, t := range s.ListarPorEstado(task.Pendiente) {
		fmt.Println(" ", t)
	}

	// --- Estadísticas ---
	fmt.Println("\n=== Estadísticas ===")
	stats := s.Estadisticas()
	for clave, valor := range stats {
		fmt.Printf("  %-15s: %d\n", clave, valor)
	}

	// --- Eliminar tarea ---
	fmt.Println("\n=== Eliminar ===")
	if err := s.Eliminar(4); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Tarea 4 eliminada")
	}

	// Eliminar tarea inexistente
	if err := s.Eliminar(99); err != nil {
		fmt.Println("Error esperado:", err)
		fmt.Println("Is ErrTareaNoEncontrada:", errors.Is(err, task.ErrTareaNoEncontrada))
	}

	fmt.Printf("\nTotal final: %d tareas\n", s.Estadisticas()["total"])
}
