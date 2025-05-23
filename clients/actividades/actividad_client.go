package actividad

import (
	"arqsoft_proyecto/model"

	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetActividadById(id int) model.Actividad {
	var actividad model.Actividad

	Db.Where("id = ?", id).First(&actividad)
	log.Debug("Act: ", actividad)

	return actividad
}