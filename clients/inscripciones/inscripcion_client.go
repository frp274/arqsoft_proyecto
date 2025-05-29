package inscripcion

import (
	"arqsoft_proyecto/model"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
	
)

var Db *gorm.DB

func InscripcionActividad(inscripcion model.Inscripcion) model.Inscripcion {
	result := Db.Create(&inscripcion)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Inscripcion realizada", inscripcion.Id)
	return inscripcion
}
