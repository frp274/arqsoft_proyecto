package actividad

import (
	"arqsoft_proyecto/model"

	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
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

// func GetAllActividades() []model.Actividades{
// 	var actividades []model.Actividades
// 	Db.Find(&actividades)

// 	return actividades
// }

func InsertActividad(actividad model.Actividad) model.Actividad{
	result := Db.Create(&actividad)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debugf("Actividad Created: %v", actividad.Id)
	return actividad
}
