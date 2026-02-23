package services

import (
	"api_actividades/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarActividadConcurrently_Exito(t *testing.T) {
	// GIVEN: Una actividad perfectamente válida
	actividad := dto.ActividadDto{
		Nombre:   "Crossfit Extremo",
		Profesor: "Juan Perez",
		Horario: []dto.HorarioDto{
			{
				Dia:        "Lunes",
				HoraInicio: "10:00",
				HoraFin:    "11:00",
				Cupo:       20,
			},
		},
	}

	// WHEN: Ejecutamos la validación concurrente
	err := ValidarActividadConcurrently(actividad)

	// THEN: No debería retornar ningún error
	assert.Nil(t, err)
}

func TestValidarActividadConcurrently_FaltaNombre(t *testing.T) {
	// GIVEN: Una actividad sin nombre
	actividad := dto.ActividadDto{
		Nombre:   "",
		Profesor: "Juan Perez",
		Horario: []dto.HorarioDto{
			{Dia: "Lunes", HoraInicio: "10:00", HoraFin: "11:00", Cupo: 20},
		},
	}

	// WHEN
	err := ValidarActividadConcurrently(actividad)

	// THEN
	assert.NotNil(t, err)
	assert.Equal(t, "el nombre de la actividad no puede estar vacío", err.Error())
}

func TestValidarActividadConcurrently_FaltaProfesor(t *testing.T) {
	// GIVEN: Una actividad sin profesor
	actividad := dto.ActividadDto{
		Nombre:   "Yoga",
		Profesor: "",
		Horario: []dto.HorarioDto{
			{Dia: "Lunes", HoraInicio: "10:00", HoraFin: "11:00", Cupo: 20},
		},
	}

	// WHEN
	err := ValidarActividadConcurrently(actividad)

	// THEN
	assert.NotNil(t, err)
	assert.Equal(t, "el nombre del profesor no puede estar vacío", err.Error())
}

func TestValidarActividadConcurrently_SinHorarios(t *testing.T) {
	// GIVEN: Una actividad sin horarios definidos
	actividad := dto.ActividadDto{
		Nombre:   "Zumba",
		Profesor: "Marta",
		Horario:  []dto.HorarioDto{}, // Vacío
	}

	// WHEN
	err := ValidarActividadConcurrently(actividad)

	// THEN
	assert.NotNil(t, err)
	assert.Equal(t, "la actividad debe tener al menos un horario", err.Error())
}
