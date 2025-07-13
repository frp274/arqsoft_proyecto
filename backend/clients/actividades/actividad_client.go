package actividad

import (
	"arqsoft_proyecto/model"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad

	result := Db.Preload("Horarios").Where("id = ?", id).First(&actividad)
	if result.Error != nil {
		//return actividad, result.Error
	}
	log.Debugf("Act: %v", actividad)

	return actividad
}

func GetAllActividades() (model.Actividades, error) {
	var actividades model.Actividades
	result := Db.Preload("Horarios").Find(&actividades)
	if result.Error != nil {
		return actividades, result.Error
	}

	return actividades, nil
}

func InsertActividad(actividad model.Actividad) model.Actividad {
	result := Db.Create(&actividad)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debugf("Actividad Created: %v", actividad.Id)
	return actividad
}

func UpdateActividad(actividad model.Actividad) model.Actividad {
	result := Db.Save(&actividad)
	if result.Error != nil {
		log.Error("")
	}
	return actividad
}

func DeleteActividad(id int) error {
	// 1. Borrar los horarios asociados a la actividad
	err := Db.Where("actividad_id = ?", id).Delete(&model.Horario{}).Error
	if err != nil {
		return fmt.Errorf("error al eliminar los horarios asociados: %w", err)
	}

	// 2. Borrar la actividad en sí
	actividad := model.Actividad{}
	result := Db.Delete(&actividad, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no se encontró actividad con ID %d", id)
	}

	return nil
}
