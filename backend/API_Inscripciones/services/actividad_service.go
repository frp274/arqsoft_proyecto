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
	//actividadDto.Cupo = actividad.Cupo
	actividadDto.Profesor = actividad.Profesor
	//actividadDto.HorarioInscripcion = actividad.Horario

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

func DeleteActividad(id int) e.ApiError {
	err := actividadCliente.DeleteActividad(id)
	if err != nil {
		return e.NewInternalServerApiError("No se pudo eliminar la actividad", err)
	}
	return nil
}

/*func UpdateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {

	// Traemos la actividad actual para asegurarnos de que existe
	actividadActual := actividadCliente.GetActividadById(actividadDto.Id)
	if actividadActual.Id == 0 {
		return dto.ActividadDto{}, e.NewNotFoundApiError("No se encontró la actividad con ese ID")
	}

	// Actualizamos los campos base
	//actividadActual.Id = actividadDto.Id
	actividadActual.Nombre = actividadDto.Nombre
	actividadActual.Descripcion = actividadDto.Descripcion
	actividadActual.Profesor = actividadDto.Profesor

	// Si vienen horarios nuevos, reemplazamos los anteriores
	// var nuevosHorarios []model.Horario
	for _, horarioDto := range actividadDto.Horario {
		horario := model.Horario{
			Id:          horarioDto.Id,
			ActividadID: actividadDto.Id,
			Dia:         horarioDto.Dia,
			HoraInicio:  horarioDto.HoraInicio,
			HoraFin:     horarioDto.HoraFin,
			Cupo:        horarioDto.Cupo,
		}
	// 	// nuevosHorarios = append(nuevosHorarios, horario)
		actividadActual.Horarios = append(actividadActual.Horarios, horario)
	}
	// actividadActual.Horarios = nuevosHorarios

	// Guardamos la actividad actualizada en la base de datos
	actividadActual = actividadCliente.UpdateActividad(actividadActual)
	if actividadActual.Id == 0{
		log.Print("no existe el id a actualizar")
		return dto.ActividadDto{}, e.NewBadRequestApiError("error, el id de la actividad no existe")
	}
	// Armamos el DTO de respuesta
	var actividadActualizada dto.ActividadDto
	actividadActualizada.Id = actividadActual.Id
	actividadActualizada.Nombre = actividadActual.Nombre
	actividadActualizada.Descripcion = actividadActual.Descripcion
	actividadActualizada.Profesor = actividadActual.Profesor

	for _, horario := range actividadActual.Horarios {
		horarioDto := dto.HorarioDto{
			Id:         horario.Id,
			Dia:        horario.Dia,
			HoraInicio: horario.HoraInicio,
			HoraFin:    horario.HoraFin,
			Cupo:       horario.Cupo,
		}
		actividadActualizada.Horario = append(actividadActualizada.Horario, horarioDto)
	}

	return actividadActualizada, nil
}*/

func UpdateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	// 1. Validar existencia de la actividad
	actividadActual := actividadCliente.GetActividadById(actividadDto.Id)
	if actividadActual.Id == 0 {
		return dto.ActividadDto{}, e.NewNotFoundApiError("No se encontró la actividad con ese ID")
	}

	// 2. Actualizar campos básicos
	actividadActual.Nombre = actividadDto.Nombre
	actividadActual.Descripcion = actividadDto.Descripcion
	actividadActual.Profesor = actividadDto.Profesor

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
	if actividadActual.Id == 0 {
		log.Print("No se pudo actualizar la actividad")
		return dto.ActividadDto{}, e.NewBadRequestApiError("Error al actualizar la actividad")
	}

	// 6. Armar respuesta DTO
	var actividadActualizada dto.ActividadDto
	actividadActualizada.Id = actividadActual.Id
	actividadActualizada.Nombre = actividadActual.Nombre
	actividadActualizada.Descripcion = actividadActual.Descripcion
	actividadActualizada.Profesor = actividadActual.Profesor

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
