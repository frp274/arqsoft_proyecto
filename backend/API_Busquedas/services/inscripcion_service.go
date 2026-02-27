package services

import (
	inscripcionCliente "api_busquedas/clients/inscripciones"
	"api_busquedas/dto"
	"api_busquedas/model"
	"api_busquedas/search"
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
		log.Warnf("No se encontraron inscripciones para el usuario %d: %v", usuarioId, er)
		return dto.InscripcionesDto{}, nil
	}

	// Convertir a DTO y enriquecer con datos de Solr
	var inscripcionesDto dto.InscripcionesDto
	for _, insc := range inscripciones {
		inscripcionDto := dto.InscripcionDto{
			Id:          insc.Id,
			UsuarioId:   insc.UsuarioId,
			ActividadId: insc.ActividadId,
			HorarioId:   insc.HorarioId,
		}

		// Buscar detalles de la actividad en Solr
		query := "id:" + insc.ActividadId
		result, err := search.SolrClient.Search(query, "", 0, 1)
		if err == nil && len(result.Response.Docs) > 0 {
			doc := result.Response.Docs[0]
			if val, ok := doc["nombre"].(string); ok {
				inscripcionDto.Nombre = val
			} else if valArr, ok := doc["nombre"].([]interface{}); ok && len(valArr) > 0 {
				inscripcionDto.Nombre = valArr[0].(string)
			}

			if val, ok := doc["descripcion"].(string); ok {
				inscripcionDto.Descripcion = val
			} else if valArr, ok := doc["descripcion"].([]interface{}); ok && len(valArr) > 0 {
				inscripcionDto.Descripcion = valArr[0].(string)
			}

			if val, ok := doc["profesor"].(string); ok {
				inscripcionDto.Profesor = val
			} else if valArr, ok := doc["profesor"].([]interface{}); ok && len(valArr) > 0 {
				inscripcionDto.Profesor = valArr[0].(string)
			}
		}

		inscripcionesDto = append(inscripcionesDto, inscripcionDto)
	}

	return inscripcionesDto, nil
}
