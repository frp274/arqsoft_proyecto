package inscripcion

import (
	"arqsoft_proyecto/model"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
	
	e "arqsoft_proyecto/utils/errors"
)

var Db *gorm.DB

func InscripcionActividad(inscripcion model.Inscripcion) (model.Inscripcion, error) {
	var verificadorInscripcion model.Inscripcion
	Db.Where("usuario_id = ? AND actividad_id = ? AND horario_id = ?", inscripcion.UsuarioId, inscripcion.ActividadId, inscripcion.HorarioId).First(&verificadorInscripcion)
	
	if verificadorInscripcion.Id != 0{
		return inscripcion, e.NewBadRequestApiError("ya esta inscripto a la actividad")
	}
	//Db.First(&verificadorInscripcion, "UsuarioId = ?", inscripcion.UsuarioId)
	result := Db.Create(&inscripcion)

	if result.Error != nil {

		log.Info(result.Error.Error())
	}
	log.Info("Inscripcion realizada", inscripcion.Id)
	return inscripcion, nil
}
