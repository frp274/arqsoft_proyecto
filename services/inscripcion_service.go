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
	inscripcion.Fecha = inscripcionDto.Fecha

	inscripcion = inscripcionCliente.InscripcionActividad(inscripcion)
	inscripcionDto.Id = inscripcion.Id
	return inscripcionDto, nil
}