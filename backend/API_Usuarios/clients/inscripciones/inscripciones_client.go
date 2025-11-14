package inscripciones

import (
	"api_usuarios/model"
	"api_usuarios/utils/errors"
	
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

// InsertInscripcion crea una nueva inscripción en la base de datos
func InsertInscripcion(inscripcion model.Inscripcion) (model.Inscripcion, errors.ApiError) {
	// Verificar si ya existe una inscripción para este usuario, actividad y horario
	var existente model.Inscripcion
	result := Db.Where("usuario_id = ? AND actividad_id = ? AND horario_id = ?", 
		inscripcion.UsuarioId, inscripcion.ActividadId, inscripcion.HorarioId).
		First(&existente)
	
	if result.Error == nil {
		// Ya existe una inscripción
		log.Warnf("Usuario %d ya está inscrito en actividad %d, horario %d", 
			inscripcion.UsuarioId, inscripcion.ActividadId, inscripcion.HorarioId)
		return model.Inscripcion{}, errors.NewBadRequestApiError("Ya estás inscrito en este horario")
	}
	
	// Crear la nueva inscripción
	result = Db.Create(&inscripcion)
	if result.Error != nil {
		log.Errorf("Error al crear inscripción: %v", result.Error)
		return model.Inscripcion{}, errors.NewInternalServerApiError("Error al crear la inscripción", result.Error)
	}
	
	log.Infof("Inscripción creada exitosamente - ID: %d, Usuario: %d, Actividad: %d, Horario: %d", 
		inscripcion.Id, inscripcion.UsuarioId, inscripcion.ActividadId, inscripcion.HorarioId)
	
	return inscripcion, nil
}

// GetInscripcionesByUsuarioId obtiene todas las inscripciones de un usuario
func GetInscripcionesByUsuarioId(usuarioId int) (model.Inscripciones, errors.ApiError) {
	var inscripciones model.Inscripciones
	
	result := Db.Where("usuario_id = ?", usuarioId).Find(&inscripciones)
	if result.Error != nil {
		log.Errorf("Error al obtener inscripciones del usuario %d: %v", usuarioId, result.Error)
		return nil, errors.NewInternalServerApiError("Error al obtener inscripciones", result.Error)
	}
	
	log.Infof("Se encontraron %d inscripciones para el usuario %d", len(inscripciones), usuarioId)
	return inscripciones, nil
}

// DeleteInscripcion elimina una inscripción por ID
func DeleteInscripcion(inscripcionId int) errors.ApiError {
	result := Db.Delete(&model.Inscripcion{}, inscripcionId)
	
	if result.Error != nil {
		log.Errorf("Error al eliminar inscripción %d: %v", inscripcionId, result.Error)
		return errors.NewInternalServerApiError("Error al eliminar la inscripción", result.Error)
	}
	
	if result.RowsAffected == 0 {
		return errors.NewNotFoundApiError("Inscripción no encontrada")
	}
	
	log.Infof("Inscripción %d eliminada exitosamente", inscripcionId)
	return nil
}
