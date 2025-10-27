package services

import (
	"api_actividades/dto"
	actividadRepositories "api_actividades/repositories/actividades"

	//"api_actividades/model"
	e "api_actividades/utils/errors"
	"errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetActividadById(id string) (dto.ActividadDto, e.ApiError) {

	actividad, err := actividadRepositories.GetActividadById(id)

	var actividadDto dto.ActividadDto

	// 1. Validar si hubo un error general en la DB
	if err != nil {
		// 2. Validar si el error fue específicamente 'documento no encontrado'
		log.Errorf("Error retrieving actividad: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Documento no encontrado
			return actividadDto, e.NewNotFoundApiError("actividad not found")
		}
		// Otro error de la DB
		return actividadDto, e.NewInternalServerApiError("Error retrieving actividad", err)
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error converting id to mongo ID: %v", err)
		return dto.ActividadDto{}, e.NewBadRequestApiError("Invalid ID format")
	}

	actividadDto.Id = objectID.Hex()
	actividadDto.Nombre = actividad.Nombre
	actividadDto.Descripcion = actividad.Descripcion
	actividadDto.Profesor = actividad.Profesor
	//actividadDto.Cupo = actividad.Cupo
	//if actividad.tags != nil {
	//	actividadDto.Tags = actividad.tags
	//}
	//actividadDto.HorarioInscripcion = actividad.Horario

	for _, horario := range actividad.Horarios {
		horarioDto := dto.HorarioDto{
			Dia:        horario.Dia,
			HoraInicio: horario.HoraInicio,
			HoraFin:    horario.HoraFin,
			Cupo:       horario.Cupo,
		}
		actividadDto.Horario = append(actividadDto.Horario, horarioDto)
	}

	log.Infof("Actividad DTO to return: %+v", actividadDto)
	return actividadDto, nil
}

/*
func GetAllActividades() (dto.ActividadesDto, e.ApiError) {
	var actividades model.Actividades
	var actividadesDto dto.ActividadesDto

	// Obtener actividades del "client" (repositorio)
	actividades, err := actividadRepositories.GetAllActividades()

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

	actividad = actividadRepositories.InsertActividad(actividad)
	actividadDto.Id = actividad.Id

	return actividadDto, nil
}

func GetActividadesByNombre(nombre string) (dto.ActividadesDto, e.ApiError) {
	actividades, err := actividadRepositories.GetActividadesFiltradas(nombre)
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
	err := actividadRepositories.DeleteActividad(id)
	if err != nil {
		return e.NewInternalServerApiError("No se pudo eliminar la actividad", err)
	}
	return nil
}


func UpdateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	// 1. Validar existencia de la actividad
	actividadActual := actividadRepositories.GetActividadById(actividadDto.Id)
	if actividadActual.Id == 0 {
		return dto.ActividadDto{}, e.NewNotFoundApiError("No se encontró la actividad con ese ID")
	}

	// 2. Actualizar campos básicos
	actividadActual.Nombre = actividadDto.Nombre
	actividadActual.Descripcion = actividadDto.Descripcion
	actividadActual.Profesor = actividadDto.Profesor

	// 3. Eliminar horarios anteriores asociados a la actividad
	err := actividadRepositories.DeleteHorariosByActividadID(actividadDto.Id)
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
	actividadActual = actividadRepositories.UpdateActividad(actividadActual)
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

*/
