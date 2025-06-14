package services

import (
	actividadCliente "arqsoft_proyecto/clients/actividades"
	"arqsoft_proyecto/dto"
	"arqsoft_proyecto/model"
	e "arqsoft_proyecto/utils/errors"
	log "github.com/sirupsen/logrus"
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
	//actividadDto.HorarioInscripcion = actividad.Horario

	for _, horario := range actividad.Horarios {
		horarioDto := dto.HorarioDto{
			Id:			horario.Id,
			Dia:        horario.Dia,
			HoraInicio: horario.HoraInicio,
			HoraFin:    horario.HoraFin,
		}
		actividadDto.Horario = append(actividadDto.Horario, horarioDto)
	}

	return actividadDto, nil
	}

func GetAllActividades() (dto.ActividadesDto, e.ApiError) {
	var actividades model.Actividades
	var actividadesDto dto.ActividadesDto

	// Obtener actividades del "client" (repositorio)
	actividades, err := actividadCliente.GetAllActividades()

	if err != nil {
		log.Error(err.Error())
		return actividadesDto, e.NewInternalServerApiError("Error", err)
	}
	// Mapear modelo → DTO
	for _, actividad := range actividades {
		actividadDto := dto.ActividadDto{
			Id:          actividad.Id,
			Nombre:      actividad.Nombre,
			Descripcion: actividad.Descripcion,
			Profesor:    actividad.Profesor,
			Cupo:        actividad.Cupo,
		}
		for _, horario := range actividad.Horarios {
			horarioDto := dto.HorarioDto{
				Id:			horario.Id,
				Dia:        horario.Dia,
				HoraInicio: horario.HoraInicio,
				HoraFin:    horario.HoraFin,
			}
			actividadDto.Horario = append(actividadDto.Horario, horarioDto)
		}	
		actividadesDto = append(actividadesDto, actividadDto)
	}

	return actividadesDto, nil // asumimos que no hay error por ahora
}

func InsertActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	var actividad model.Actividad

	actividad.Nombre = actividadDto.Nombre
	actividad.Descripcion = actividadDto.Descripcion
	actividad.Cupo = actividadDto.Cupo
	actividad.Profesor = actividadDto.Profesor



	for _, horarioDto := range actividadDto.Horario {
		horario := model.Horario{
			Id:			horarioDto.Id,
			Dia:        horarioDto.Dia,
			HoraInicio: horarioDto.HoraInicio,
			HoraFin:    horarioDto.HoraFin,
		}
		actividad.Horarios = append(actividad.Horarios, horario)
	}

	actividad = actividadCliente.InsertActividad(actividad)
	actividadDto.Id = actividad.Id

	return actividadDto, nil
}
