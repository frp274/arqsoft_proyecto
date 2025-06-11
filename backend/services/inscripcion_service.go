package services

import (
	inscripcionCliente "arqsoft_proyecto/clients/inscripciones"
	actividadCliente "arqsoft_proyecto/clients/actividades"
	"arqsoft_proyecto/dto"
	"arqsoft_proyecto/model"
	e "arqsoft_proyecto/utils/errors"
)


func InscripcionActividad(inscripcionDto dto.InscripcionDto)(dto.InscripcionDto, e.ApiError){
	var inscripcion model.Inscripcion

	inscripcion.UsuarioId = inscripcionDto.UsuarioId
	inscripcion.ActividadId = inscripcionDto.ActividadId
	inscripcion.HorarioId = inscripcionDto.HorarioId

	inscripcion = inscripcionCliente.InscripcionActividad(inscripcion)
	inscripcionDto.Id = inscripcion.Id
    actividad := actividadCliente.GetActividadById(inscripcionDto.ActividadId)
    if actividad.Id == 0 {
        // Manejá error
    } else if actividad.Cupo > 0 {
        actividad.Cupo -= 1
        actividadCliente.UpdateActividad(actividad)
    }

	return inscripcionDto, nil

}