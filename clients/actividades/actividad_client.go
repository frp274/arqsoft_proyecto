package actividad

import (
	"arqsoft_proyecto/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad

	result := Db.Where("id = ?", id).First(&actividad)
	if result.Error != nil {

	}
	log.Debugf("Act: %v", actividad)

	return actividad
}

func GetAllActividades() (model.Actividades, error) {
	var actividades model.Actividades
	result := Db.Find(&actividades)
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
