package services

import (
	actividadCliente "api_busquedas/clients/actividades"
	"api_busquedas/dto"
	"api_busquedas/model"
	e "api_busquedas/utils/errors"

	log "github.com/sirupsen/logrus"
)

func GetActividadById(id string) (dto.ActividadDto, e.ApiError) {
	// Buscamos en Solr primero o en el cliente
	actividad := actividadCliente.GetActividadById(id)
	var actividadDto dto.ActividadDto

	if actividad.Id == "" {
		return actividadDto, e.NewBadRequestApiError("actividad not found")
	}

	actividadDto.Nombre = actividad.Nombre
	actividadDto.Id = actividad.Id
	actividadDto.Descripcion = actividad.Descripcion
	actividadDto.Profesor = actividad.Profesor
	actividadDto.OwnerId = actividad.OwnerId

	for _, horario := range actividad.Horarios {
		horarioDto := dto.HorarioDto{
			Id:         horario.Id,
			Dia:        horario.Dia,
			HoraInicio: horario.HoraInicio,
			HoraFin:    horario.HoraFin,
			Cupo:       horario.Cupo,
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
			OwnerId:     actividad.OwnerId,
		}
		for _, horario := range actividad.Horarios {
			horarioDto := dto.HorarioDto{
				Id:         horario.Id,
				Dia:        horario.Dia,
				HoraInicio: horario.HoraInicio,
				HoraFin:    horario.HoraFin,
				Cupo:       horario.Cupo,
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
	//actividad.Cupo = actividadDto.Cupo
	actividad.Profesor = actividadDto.Profesor

	for _, horarioDto := range actividadDto.Horario {
		horario := model.Horario{
			Id:          horarioDto.Id,
			ActividadID: actividad.Id,
			Dia:         horarioDto.Dia,
			HoraInicio:  horarioDto.HoraInicio,
			HoraFin:     horarioDto.HoraFin,
			Cupo:        horarioDto.Cupo,
		}
		actividad.Horarios = append(actividad.Horarios, horario)
	}

	actividad = actividadCliente.InsertActividad(actividad)
	actividadDto.Id = actividad.Id

	return actividadDto, nil
}

func GetActividadesByNombre(nombre string) (dto.ActividadesDto, e.ApiError) {
	actividades, err := actividadCliente.GetActividadesFiltradas(nombre)
	if err != nil {
		log.Error(err.Error())
		return nil, e.NewInternalServerApiError("Error al obtener actividades", err)
	}

	var actividadesDto dto.ActividadesDto
	for _, actividad := range actividades {
		actividadDto := dto.ActividadDto{
			Id:          actividad.Id,
			Nombre:      actividad.Nombre,
			Descripcion: actividad.Descripcion,
			Profesor:    actividad.Profesor,
			OwnerId:     actividad.OwnerId,
		}

		for _, horario := range actividad.Horarios {
			horarioDto := dto.HorarioDto{
				Id:         horario.Id,
				Dia:        horario.Dia,
				HoraInicio: horario.HoraInicio,
				HoraFin:    horario.HoraFin,
				Cupo:       horario.Cupo,
			}
			actividadDto.Horario = append(actividadDto.Horario, horarioDto)
		}

		actividadesDto = append(actividadesDto, actividadDto)
	}

	return actividadesDto, nil
}

func UpdateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	// 1. Validar existencia de la actividad
	actividadActual := actividadCliente.GetActividadById(actividadDto.Id)
	if actividadActual.Id == "" {
		return dto.ActividadDto{}, e.NewNotFoundApiError("No se encontró la actividad con ese ID")
	}

	// 2. Actualizar campos básicos
	actividadActual.Nombre = actividadDto.Nombre
	actividadActual.Descripcion = actividadDto.Descripcion
	actividadActual.Profesor = actividadDto.Profesor
	actividadActual.OwnerId = actividadDto.OwnerId

	// 3. Eliminar horarios anteriores asociados a la actividad
	err := actividadCliente.DeleteHorariosByActividadID(actividadDto.Id)
	if err != nil {
		log.Print("Error al eliminar horarios anteriores: ", err)
		return dto.ActividadDto{}, e.NewInternalServerApiError("Error al eliminar horarios anteriores", err)
	}

	// 4. Crear nuevos horarios desde el DTO
	var nuevosHorarios []model.Horario
	for _, horarioDto := range actividadDto.Horario {
		nuevoHorario := model.Horario{
			Id:          horarioDto.Id,
			ActividadID: actividadDto.Id,
			Dia:         horarioDto.Dia,
			HoraInicio:  horarioDto.HoraInicio,
			HoraFin:     horarioDto.HoraFin,
			Cupo:        horarioDto.Cupo,
		}
		nuevosHorarios = append(nuevosHorarios, nuevoHorario)
	}

	// 5. Guardar actividad actualizada
	actividadActual.Horarios = nuevosHorarios
	actividadActual = actividadCliente.UpdateActividad(actividadActual)
	if actividadActual.Id == "" {
		log.Print("No se pudo actualizar la actividad")
		return dto.ActividadDto{}, e.NewBadRequestApiError("Error al actualizar la actividad")
	}

	// 6. Armar respuesta DTO
	var actividadActualizada dto.ActividadDto
	actividadActualizada.Id = actividadActual.Id
	actividadActualizada.Nombre = actividadActual.Nombre
	actividadActualizada.Descripcion = actividadActual.Descripcion
	actividadActualizada.Profesor = actividadActual.Profesor
	actividadActualizada.OwnerId = actividadActual.OwnerId

	for _, h := range actividadActual.Horarios {
		actividadActualizada.Horario = append(actividadActualizada.Horario, dto.HorarioDto{
			Id:         h.Id,
			Dia:        h.Dia,
			HoraInicio: h.HoraInicio,
			HoraFin:    h.HoraFin,
			Cupo:       h.Cupo,
		})
	}

	return actividadActualizada, nil
}

func DeleteActividad(id string) e.ApiError {
	err := actividadCliente.DeleteActividad(id)
	if err != nil {
		return e.NewInternalServerApiError("No se pudo eliminar la actividad", err)
	}
	return nil
}
