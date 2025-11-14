package services

import (
	inscripcionCliente "api_busquedas/clients/inscripciones"
	"api_busquedas/dto"
	"api_busquedas/model"
	e "api_busquedas/utils/errors"

	log "github.com/sirupsen/logrus"
)

func InscripcionActividad(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, e.ApiError) {
	var inscripcion model.Inscripcion

	inscripcion.UsuarioId = inscripcionDto.UsuarioId
	inscripcion.ActividadId = inscripcionDto.ActividadId
	inscripcion.HorarioId = inscripcionDto.HorarioId

	// Crear la inscripción en PostgreSQL
	inscripcion, err := inscripcionCliente.InscripcionActividad(inscripcion)
	if err != nil {
		return dto.InscripcionDto{}, e.NewBadRequestApiError("ya estas inscripto a la actividad")
	}
	inscripcionDto.Id = inscripcion.Id

	// NOTA: Los horarios están en MongoDB (API_Actividades), no en PostgreSQL
	// Por ahora solo creamos la inscripción, el decremento del cupo se manejará en API_Actividades
	log.Info("Inscripción creada exitosamente:", inscripcionDto.Id)

	return inscripcionDto, nil

}

func GetInscripcionesByUsuarioId(usuarioId int) (dto.InscripcionesDto, e.ApiError) {
	// Obtener inscripciones del usuario (solo IDs)
	inscripciones, er := inscripcionCliente.GetInscripcionesByUsuarioId(usuarioId)

	if er != nil {
		log.Error(er.Error())
		return dto.InscripcionesDto{}, e.NewBadRequestApiError("no se encontraron inscripciones para el usuario")
	}

	// Convertir a DTO
	var inscripcionesDto dto.InscripcionesDto
	for _, insc := range inscripciones {
		inscripcionDto := dto.InscripcionDto{
			Id:          insc.Id,
			UsuarioId:   insc.UsuarioId,
			ActividadId: insc.ActividadId,
			HorarioId:   insc.HorarioId,
		}
		inscripcionesDto = append(inscripcionesDto, inscripcionDto)
	}

	return inscripcionesDto, nil
}
