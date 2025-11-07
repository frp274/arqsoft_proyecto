package services

import (
	"api_actividades/dto"
	actividadRepositories "api_actividades/repositories/actividades"

	"api_actividades/model"
	e "api_actividades/utils/errors"
	"errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetActividadById(id string) (dto.ActividadDto, e.ApiError) {
	var actividadDto dto.ActividadDto
	//Aplicar Busqueda en la cache =====================================================================================================================================
	actividad, er := actividadRepositories.GetActividadByIdCache(id)
	if er != nil {
		log.Errorf("Error retrieving actividad from cache: %v", er)
	}
	if actividad.Id.IsZero() {

		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Errorf("Error converting id to mongo ID: %v", err)
			return dto.ActividadDto{}, e.NewBadRequestApiError("Invalid ID format")
		}

		actividad, err = actividadRepositories.GetActividadById(objectID)
		
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
		// Insertar en la cache local
		actividadCache := actividadRepositories.InsertActividadCache(actividad)
		log.Infof("Actividad insertada en la cache local: %v", actividadCache)
	}


	actividadDto.Id = id //objectID.Hex()
	log.Debugf("Nombre de la actividad: %v", actividad.Nombre)
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
func GetActividadesByNombre(nombre string) (dto.ActividadesDto, e.ApiError) {
	// Llamar a la API_Busquedas, para encontrar las actividades que coincidan con el nombre ===================================================================================
	if len(nombre) == 0 {
		return dto.ActividadesDto{}, e.NewBadRequestApiError("Invalid name format")
	}

	actividades, err := actividadRepositories.GetActividadesByNombre(nombre)
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
*/
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
*/
func InsertActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	var actividad model.Actividad

	if actividadDto.Nombre == "" {
		log.Error("El nombre de la actividad no puede estar vacío")
		return dto.ActividadDto{}, e.NewBadRequestApiError("El nombre de la actividad no puede estar vacío")
	}
	if len(actividadDto.Horario) == 0 {
		log.Error("La actividad debe tener al menos un horario")
		return dto.ActividadDto{}, e.NewBadRequestApiError("La actividad debe tener al menos un horario")
	}
	if actividadDto.Profesor == "" {
		log.Error("El nombre del profesor no puede estar vacío")
		return dto.ActividadDto{}, e.NewBadRequestApiError("El nombre del profesor no puede estar vacío")
	}
	actividad.Nombre = actividadDto.Nombre
	actividad.Descripcion = actividadDto.Descripcion
	//actividad.Cupo = actividadDto.Cupo
	actividad.Profesor = actividadDto.Profesor

	for _, horarioDto := range actividadDto.Horario {
		horario := model.Horario{
			Dia:         horarioDto.Dia,
			HoraInicio:  horarioDto.HoraInicio,
			HoraFin:     horarioDto.HoraFin,
			Cupo:        horarioDto.Cupo,
		}
		actividad.Horarios = append(actividad.Horarios, horario)
	}

	actividadInsertada, err := actividadRepositories.InsertActividad(actividad)
	if err != nil {
		log.Errorf("Error al insertar actividad: %v", err)
		return dto.ActividadDto{}, e.NewInternalServerApiError("Error al insertar actividad", err)
	}

	//Luego de insertar y que haya salido todo bien, falta insertarla en localcache e indexarla en soler (para busquedas)================================================================
	actividadCache := actividadRepositories.InsertActividadCache(actividadInsertada)
	// Preguntar al profe si conviene en vez de retornar,			==================================================================================================
	// simplemente cuando devuelva el error en la funcion general, 	==================================================================================================
	// mostrar el error que hubo									==================================================================================================
	
	log.Infof("Actividad insertada en la cache local: %v", actividadCache)

	actividadDto.Id = actividadInsertada.Id.Hex()

	return actividadDto, nil  // Mandar aca el error de la cache en caso de que haya?? ==============================^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
}


func DeleteActividad(id string) e.ApiError {
	//Falta eliminar en Solr y en la cache local===============================================================================================================================
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error converting id to mongo ID: %v", err)
		return e.NewBadRequestApiError("Invalid ID format")
	}

	err = actividadRepositories.DeleteActividad(objectID)
	if err != nil {
		return e.NewInternalServerApiError("No se pudo eliminar la actividad", err)
	}
	//Eliminar de la cache local
	er := actividadRepositories.DeleteActividadCache(id)
	if er != nil {
		return e.NewInternalServerApiError("No se pudo eliminar la actividad de la cache local", er)
	}
	return nil
}


func UpdateActividad(actividadDto dto.ActividadDto) (dto.ActividadDto, e.ApiError) {
	//Falta actualizar en Solr ===============================================================================================================================
	objetID, err := primitive.ObjectIDFromHex(actividadDto.Id)
	if err != nil {
		log.Errorf("Error converting id to mongo ID: %v", err)
		return dto.ActividadDto{}, e.NewBadRequestApiError("Invalid ID format")
	}

	// 1. Validar existencia de la actividad
	actividadActual, err := actividadRepositories.GetActividadById(objetID)
	if err != nil {
		log.Print("Error al obtener la actividad: ", err)
		return dto.ActividadDto{}, e.NewNotFoundApiError("Actividad no encontrada")
	}

	// 2. Actualizar campos básicos
	if actividadDto.Nombre != "" {
		actividadActual.Nombre = actividadDto.Nombre
	}
	if actividadDto.Descripcion != "" {
		actividadActual.Descripcion = actividadDto.Descripcion
	}
	if actividadDto.Profesor != "" {
		actividadActual.Profesor = actividadDto.Profesor
	}
	if actividadDto.Horario != nil {
		nuevosHorarios := make([]model.Horario, 0, len(actividadDto.Horario))
		for _, horarioDto := range actividadDto.Horario {
			nuevosHorarios = append(nuevosHorarios, model.Horario{
				Dia:        horarioDto.Dia,
				HoraInicio: horarioDto.HoraInicio,
				HoraFin:    horarioDto.HoraFin,
				Cupo:       horarioDto.Cupo,
			})
			actividadActual.Horarios = nuevosHorarios
		}
	}

	// 3. Guardar cambios en la base de datos
	actividadActual, err = actividadRepositories.UpdateActividad(actividadActual)
	if err != nil {
		log.Print("Error al actualizar la actividad: ", err)
		return dto.ActividadDto{}, e.NewInternalServerApiError("Error al actualizar la actividad", err)
	}

	//Actualizar en la cache local 
	actividadCache := actividadRepositories.InsertActividadCache(actividadActual)
	
	log.Infof("Actividad actualizada en la cache local: %v", actividadCache)

	// 6. Armar respuesta DTO
	var actividadActualizada dto.ActividadDto
	actividadActualizada.Id = actividadActual.Id.Hex()
	actividadActualizada.Nombre = actividadActual.Nombre
	actividadActualizada.Descripcion = actividadActual.Descripcion
	actividadActualizada.Profesor = actividadActual.Profesor

	for _, h := range actividadActual.Horarios {
		actividadActualizada.Horario = append(actividadActualizada.Horario, dto.HorarioDto{
			Dia:        h.Dia,
			HoraInicio: h.HoraInicio,
			HoraFin:    h.HoraFin,
			Cupo:       h.Cupo,
		})
	}

	return actividadActualizada, nil
}

