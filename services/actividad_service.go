package services

import (
	actividadCliente "arqsoft_proyecto/clients/actividades"
	"arqsoft_proyecto/dto"
	"arqsoft_proyecto/model"
	e "arqsoft_proyecto/utils/errors"
)

func GetActividadById(id int) (dto.ActividadDto, e.ApiError) {
	var actividad model.Actividad = actividadCliente.GetActividadById(id)
	var actividadDto dto.ActividadDto

	if actividad.Id == 0 {
		return actividadDto, e.NewBadRequestApiError("actividad not found")
	}

	actividadDto.Nombre = actividad.Nombre
	actividadDto.Id = actividad.Id
	actividadDto.Descripcion = actividad.Descripcion
	actividadDto.Cupo = actividad.Cupo
	actividadDto.Profesor = actividad.Profesor

	return actividadDto, nil
}


func InsertActividad(actividadDto dto.ActividadDto)(dto.ActividadDto, e.ApiError){
	var actividad model.Actividad
	
	actividad.Nombre = actividadDto.Nombre
	actividad.Descripcion = actividadDto.Descripcion
	actividad.Cupo = actividadDto.Cupo
	actividad.Profesor = actividadDto.Profesor

	actividad = actividadCliente.InsertActividad(actividad)
	actividadDto.Id = actividad.Id

	return actividadDto, nil
}