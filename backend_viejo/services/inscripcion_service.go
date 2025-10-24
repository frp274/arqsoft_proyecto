package services

import (
	actividadCliente "arqsoft_proyecto/clients/actividades"
	inscripcionCliente "arqsoft_proyecto/clients/inscripciones"
	"arqsoft_proyecto/dto"
	"arqsoft_proyecto/model"
	e "arqsoft_proyecto/utils/errors"

	log "github.com/sirupsen/logrus"
)

func InscripcionActividad(inscripcionDto dto.InscripcionDto) (dto.InscripcionDto, e.ApiError) {
	var inscripcion model.Inscripcion

	inscripcion.UsuarioId = inscripcionDto.UsuarioId
	inscripcion.ActividadId = inscripcionDto.ActividadId
	inscripcion.HorarioId = inscripcionDto.HorarioId

	inscripcion, err := inscripcionCliente.InscripcionActividad(inscripcion)
	if err != nil {
		return dto.InscripcionDto{}, e.NewBadRequestApiError("ya estas inscripto a la actividad")
	}
	inscripcionDto.Id = inscripcion.Id
	actividad := actividadCliente.GetActividadById(inscripcionDto.ActividadId)
	horario, errror := inscripcionCliente.GetCupoByHorarioId(inscripcion.HorarioId)
	if errror != nil {
		return dto.InscripcionDto{}, e.NewBadRequestApiError("no se encontro el horario")
	}
	if actividad.Id == 0 {
		// Manejá error
	} else if horario.Cupo > 0 {
		horario.Cupo -= 1
		inscripcionCliente.UpdateInscripcion(horario)
	}

	return inscripcionDto, nil

}

func GetInscripcionesByUsuarioId(usuarioId int) (dto.ActividadesDto, e.ApiError) {
	var actividadesDto dto.ActividadesDto

	actividades, er := inscripcionCliente.GetInscripcionesByUsuarioId(usuarioId)

	if er != nil {
		log.Error(er.Error())
		return dto.ActividadesDto{}, e.NewBadRequestApiError("no se encontro el usuario id")
	}

	for _, actividad := range actividades {
		actividadDto := dto.ActividadDto{
			Id:          actividad.Id,
			Nombre:   actividad.Nombre,
			Descripcion: actividad.Descripcion,
			Profesor:   actividad.Profesor,
		}
		actividadesDto = append(actividadesDto, actividadDto)
	}

	return actividadesDto, nil
}
