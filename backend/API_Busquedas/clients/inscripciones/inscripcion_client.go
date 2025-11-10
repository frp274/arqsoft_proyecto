package inscripcion

import (
	"api_busquedas/model"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
	
	e "api_busquedas/utils/errors"
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


func GetCupoByHorarioId(horarioId int) (model.Horario, error){
	var horario model.Horario
	Db.Where("Id = ?", horarioId).First(&horario)
	if horario.Id != 0{
		return horario, nil
	}

	return horario, e.NewNotFoundApiError("no se encontro el horario")
}

func UpdateInscripcion(horario model.Horario) model.Horario {
    result := Db.Save(&horario)
    if result.Error != nil {
		log.Error("")
    }
    return horario
}

func GetInscripcionesByUsuarioId(usuarioId int) (model.Actividades, error) {
	var actividades model.Actividades
	result := Db.Preload("Horarios").Where("id = ?", usuarioId).Find(&actividades)
	//result := Db.Where("usuario_id = ?", usuarioId).Find(&inscripciones)
	if result.Error != nil {
		return model.Actividades{}, result.Error
	}
	log.Debugf("Ins: %v", actividades)

	return actividades, nil
}
