package services

import (
	inscripcionCliente "arqsoft_proyecto/clients/inscripciones"
	"arqsoft_proyecto/dto"
	"arqsoft_proyecto/model"
	e "arqsoft_proyecto/utils/errors"
)


func InscripcionActividad(inscripcionDto dto.InscripcionDto)(dto.InscripcionDto, e.ApiError){
	var inscripcion model.Inscripcion

	inscripcion.UsuarioId = inscripcionDto.UsuarioId
	inscripcion.ActividadId = inscripcionDto.ActividadId
	inscripcion.HorarioInscripcion.Dia = inscripcionDto.HorarioInscripcion.Dia
	inscripcion.HorarioInscripcion.HoraInicio = inscripcionDto.HorarioInscripcion.HoraInicio
	inscripcion.HorarioInscripcion.HoraFin = inscripcionDto.HorarioInscripcion.HoraFin

	inscripcion = inscripcionCliente.InscripcionActividad(inscripcion)
	inscripcionDto.Id = inscripcion.Id
	return inscripcionDto, nil

}