package services

import (
	inscripcionesClient "api_usuarios/clients/inscripciones"
	"api_usuarios/dto"
	"api_usuarios/model"
	"api_usuarios/utils/errors"
	
	log "github.com/sirupsen/logrus"
)

// CreateInscripcion crea una nueva inscripción
func CreateInscripcion(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, errors.ApiError) {
	log.Infof("Creando inscripción - Usuario: %d, Actividad: %d, Horario: %d", 
		inscripcionDto.UsuarioId, inscripcionDto.ActividadId, inscripcionDto.HorarioId)
	
	// Validar datos
	if inscripcionDto.UsuarioId <= 0 {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de usuario inválido")
	}
	if inscripcionDto.ActividadId <= 0 {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de actividad inválido")
	}
	if inscripcionDto.HorarioId < 0 {
		return dto.InscripcionDto{}, errors.NewBadRequestApiError("ID de horario inválido")
	}
	
	// Convertir DTO a Model
	inscripcion := model.Inscripcion{
		UsuarioId:   inscripcionDto.UsuarioId,
		ActividadId: inscripcionDto.ActividadId,
		HorarioId:   inscripcionDto.HorarioId,
	}
	
	// Crear inscripción en BD
	inscripcion, err := inscripcionesClient.InsertInscripcion(inscripcion)
	if err != nil {
		return dto.InscripcionDto{}, err
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
