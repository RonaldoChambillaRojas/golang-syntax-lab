// Package store gestiona la persistencia en memoria de las tareas.
package store

import (
	"fmt"
	"sync"

	"golang-syntax-lab/11-paquetes/proyecto-taskmanager/task"
)

// Store es un repositorio de tareas en memoria.
// Usa sync.RWMutex para ser seguro en concurrencia (veremos esto en sec. 12).
type Store struct {
	mu      sync.RWMutex
	tareas  map[int]*task.Tarea
	nextID  int
}

// New crea un nuevo Store vacío.
func New() *Store {
	return &Store{
		tareas: make(map[int]*task.Tarea),
		nextID: 1,
	}
}

// Crear agrega una nueva tarea al store.
func (s *Store) Crear(titulo, descripcion string, prioridad task.Prioridad) (*task.Tarea, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, err := task.New(s.nextID, titulo, descripcion, prioridad)
	if err != nil {
		return nil, fmt.Errorf("store.Crear: %w", err)
	}
	s.tareas[t.ID] = t
	s.nextID++
	return t, nil
}

// ObtenerPorID busca una tarea por su ID.
func (s *Store) ObtenerPorID(id int) (*task.Tarea, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.tareas[id]
	if !ok {
		return nil, fmt.Errorf("store.ObtenerPorID(%d): %w", id, task.ErrTareaNoEncontrada)
	}
	return t, nil
}

// Listar retorna todas las tareas.
func (s *Store) Listar() []*task.Tarea {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resultado := make([]*task.Tarea, 0, len(s.tareas))
	for _, t := range s.tareas {
		resultado = append(resultado, t)
	}
	return resultado
}

// ListarPorEstado filtra tareas por estado.
func (s *Store) ListarPorEstado(estado task.Estado) []*task.Tarea {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var resultado []*task.Tarea
	for _, t := range s.tareas {
		if t.Estado == estado {
			resultado = append(resultado, t)
		}
	}
	return resultado
}

// Eliminar borra una tarea del store.
func (s *Store) Eliminar(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tareas[id]; !ok {
		return fmt.Errorf("store.Eliminar(%d): %w", id, task.ErrTareaNoEncontrada)
	}
	delete(s.tareas, id)
	return nil
}

// Estadisticas retorna un resumen del store.
func (s *Store) Estadisticas() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := map[string]int{
		"total":      len(s.tareas),
		"pendiente":  0,
		"en_progreso": 0,
		"completada": 0,
		"cancelada":  0,
	}
	for _, t := range s.tareas {
		switch t.Estado {
		case task.Pendiente:
			stats["pendiente"]++
		case task.EnProgreso:
			stats["en_progreso"]++
		case task.Completada:
			stats["completada"]++
		case task.Cancelada:
			stats["cancelada"]++
		}
	}
	return stats
}
