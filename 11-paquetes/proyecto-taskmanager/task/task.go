// Package task define el dominio de tareas del TaskManager.
package task

import (
	"errors"
	"fmt"
	"time"
)

// Estado representa el ciclo de vida de una tarea.
type Estado int

const (
	Pendiente  Estado = iota
	EnProgreso        // no hay espacio entre el nombre y el valor en iota seguido
	Completada
	Cancelada
)

func (e Estado) String() string {
	switch e {
	case Pendiente:
		return "Pendiente"
	case EnProgreso:
		return "En Progreso"
	case Completada:
		return "Completada"
	case Cancelada:
		return "Cancelada"
	default:
		return fmt.Sprintf("Estado(%d)", int(e))
	}
}

// Prioridad define la urgencia de una tarea.
type Prioridad int

const (
	Baja Prioridad = iota + 1
	Media
	Alta
	Critica
)

func (p Prioridad) String() string {
	switch p {
	case Baja:
		return "Baja"
	case Media:
		return "Media"
	case Alta:
		return "Alta"
	case Critica:
		return "Crítica"
	default:
		return "Desconocida"
	}
}

// Sentinel errors del dominio.
var (
	ErrTareaNoEncontrada = errors.New("tarea no encontrada")
	ErrTransicionInvalida = errors.New("transición de estado inválida")
	ErrTituloVacio       = errors.New("el título no puede estar vacío")
)

// Tarea es la entidad central del sistema.
type Tarea struct {
	ID          int
	Titulo      string
	Descripcion string
	Estado      Estado
	Prioridad   Prioridad
	CreadaEn    time.Time
	ActualizadaEn time.Time
}

// New crea una nueva tarea validada.
func New(id int, titulo, descripcion string, prioridad Prioridad) (*Tarea, error) {
	if titulo == "" {
		return nil, ErrTituloVacio
	}
	ahora := time.Now()
	return &Tarea{
		ID:            id,
		Titulo:        titulo,
		Descripcion:   descripcion,
		Estado:        Pendiente,
		Prioridad:     prioridad,
		CreadaEn:      ahora,
		ActualizadaEn: ahora,
	}, nil
}

// Iniciar transiciona la tarea a EnProgreso.
func (t *Tarea) Iniciar() error {
	if t.Estado != Pendiente {
		return fmt.Errorf("iniciar tarea %d: %w (estado actual: %s)", t.ID, ErrTransicionInvalida, t.Estado)
	}
	t.Estado = EnProgreso
	t.ActualizadaEn = time.Now()
	return nil
}

// Completar transiciona la tarea a Completada.
func (t *Tarea) Completar() error {
	if t.Estado != EnProgreso {
		return fmt.Errorf("completar tarea %d: %w (estado actual: %s)", t.ID, ErrTransicionInvalida, t.Estado)
	}
	t.Estado = Completada
	t.ActualizadaEn = time.Now()
	return nil
}

// Cancelar transiciona la tarea a Cancelada.
func (t *Tarea) Cancelar() error {
	if t.Estado == Completada || t.Estado == Cancelada {
		return fmt.Errorf("cancelar tarea %d: %w (ya está %s)", t.ID, ErrTransicionInvalida, t.Estado)
	}
	t.Estado = Cancelada
	t.ActualizadaEn = time.Now()
	return nil
}

// EstaActiva reporta si la tarea puede recibir trabajo.
func (t *Tarea) EstaActiva() bool {
	return t.Estado == Pendiente || t.Estado == EnProgreso
}

func (t *Tarea) String() string {
	return fmt.Sprintf("[%d] %s [%s/%s]", t.ID, t.Titulo, t.Estado, t.Prioridad)
}
