package services

import (
	"api_usuarios/clients/actividades"
	inscripcionesClient "api_usuarios/clients/inscripciones"
	"api_usuarios/dto"
	"api_usuarios/model"
	"api_usuarios/utils/errors"

	log "github.com/sirupsen/logrus"
)

// CreateInscripcion crea una nueva inscripción
func CreateInscripcion(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, errors.ApiError) {
	log.Infof("Creando inscripción - Usuario: %d, Actividad: %s, Horario: %s", inscripcionDto.UsuarioId, inscripcionDto.ActividadId, inscripcionDto.HorarioId)

	// Validar datos
	if inscripcionDto.UsuarioId == 0 {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de usuario inválido")
	}
	if inscripcionDto.ActividadId == "" {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de actividad inválido")
	}
	if inscripcionDto.HorarioId == "" {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de horario inválido")
	}

	// Validar existencia de actividad
	actividad, err := actividades.GetActividadById(inscripcionDto.ActividadId, inscripcionDto.HorarioId)
	if err != nil {
		log.Errorf("Error getting actividad %s: %v", inscripcionDto.ActividadId, err)
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("Actividad no encontrada en la base de datos")
	}
	if actividad.Id == "" {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("Actividad no encontrada")
	}
	// Convertir DTO a Model
	inscripcion := model.Inscripcion{
		UsuarioId:   inscripcionDto.UsuarioId,
		ActividadId: inscripcionDto.ActividadId,
		HorarioId:   inscripcionDto.HorarioId,
	}

	// Crear inscripción en BD
	inscripcion, apiErr := inscripcionesClient.InsertInscripcion(inscripcion)
	if apiErr != nil {
		return dto.InscripcionDto{}, apiErr
	}

	// Decrementar cupo en API_Actividades
	if err := actividades.DecrementarCupo(inscripcionDto.ActividadId, inscripcionDto.HorarioId); err != nil {
		log.Warnf("Error al decrementar cupo para actividad %s: %v", inscripcionDto.ActividadId, err)
		// No retornamos error de inscripción porque la inscripción ya se guardó.
		// Podríamos considerar rollback, pero por ahora logueamos el aviso.
	}

	// Convertir Model a DTO
	inscripcionDto.Id = inscripcion.Id

	log.Infof("Inscripción creada exitosamente con ID: %d", inscripcionDto.Id)
	return inscripcionDto, nil
}

// GetInscripcionesByUsuarioId obtiene todas las inscripciones de un usuario
func GetInscripcionesByUsuarioId(usuarioId int) (dto.InscripcionesDto, errors.ApiError) {
	log.Infof("Obteniendo inscripciones del usuario: %d", usuarioId)

	if usuarioId <= 0 {
		return nil, errors.NewBadRequestApiError("ID de usuario inválido")
	}

	// Obtener inscripciones de BD
	inscripciones, err := inscripcionesClient.GetInscripcionesByUsuarioId(usuarioId)
	if err != nil {
		return nil, err
	}

	// Convertir Models a DTOs
	var inscripcionesDto dto.InscripcionesDto
	for _, inscripcion := range inscripciones {
		inscripcionDto := dto.InscripcionDto{
			Id:          inscripcion.Id,
			UsuarioId:   inscripcion.UsuarioId,
			ActividadId: inscripcion.ActividadId,
			HorarioId:   inscripcion.HorarioId,
		}
		inscripcionesDto = append(inscripcionesDto, inscripcionDto)
	}

	log.Infof("Se encontraron %d inscripciones para el usuario %d", len(inscripcionesDto), usuarioId)
	return inscripcionesDto, nil
}

// DeleteInscripcion elimina una inscripción
func DeleteInscripcion(inscripcionId int) errors.ApiError {
	log.Infof("Eliminando inscripción: %d", inscripcionId)

	if inscripcionId <= 0 {
		return errors.NewBadRequestApiError("ID de inscripción inválido")
	}

	err := inscripcionesClient.DeleteInscripcion(inscripcionId)
	if err != nil {
		return err
	}

	log.Infof("Inscripción %d eliminada exitosamente", inscripcionId)
	return nil
}
