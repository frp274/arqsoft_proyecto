package services

import (
	"api_actividades/cache"
	"api_actividades/clients/usuarios"
	"api_actividades/dto"
	"api_actividades/queue"
	actividadRepositories "api_actividades/repositories/actividades"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"api_actividades/model"
	"api_actividades/utils"
	e "api_actividades/utils/errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ValidarActividadConcurrently(a dto.ActividadDto) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 3) // una por cada validación

	// VALIDACIÓN 1: Nombre
	wg.Add(1)
	go func() {
		defer wg.Done()
		if a.Nombre == "" {
			errChan <- errors.New("el nombre de la actividad no puede estar vacío")
		}
	}()

	// VALIDACIÓN 2: Horarios
	wg.Add(1)
	go func() {
		defer wg.Done()
		if len(a.Horario) == 0 {
			errChan <- errors.New("la actividad debe tener al menos un horario")
		}
	}()

	// VALIDACIÓN 3: Profesor
	wg.Add(1)
	go func() {
		defer wg.Done()
		if a.Profesor == "" {
			errChan <- errors.New("el nombre del profesor no puede estar vacío")
		}
	}()

	// VALIDACIÓN 4: Owner ID
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	if a.OwnerId == 0 {
	//		errChan <- errors.New("owner_id es requerido")
	//	}
	//}()

	// Esperar a que todas terminen y cerrar canal
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Retornar el primer error encontrado
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func GetActividadById(id string) (dto.ActividadDto, e.ApiError) {
	var actividad model.Actividad

	// 1. Intentar obtener de la DOBLE CAPA DE CACHÉ (L1 + L2)
	cacheKey := fmt.Sprintf("actividad:%s", id)
	if found := cache.GetJSON(cacheKey, &actividad); found {
		log.Infof("Actividad %s encontrada en caché", id)
	} else {
		// 2. Si no está en caché, buscar en MongoDB
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Errorf("Error converting id to mongo ID: %v", err)
			return dto.ActividadDto{}, e.NewBadRequestApiError("Invalid ID format")
		}

		actividad, err = actividadRepositories.GetActividadById(objectID)
		if err != nil {
			log.Errorf("Error retrieving actividad: %v", err)
			if errors.Is(err, mongo.ErrNoDocuments) {
				return dto.ActividadDto{}, e.NewNotFoundApiError("actividad not found")
			}
			return dto.ActividadDto{}, e.NewInternalServerApiError("Error retrieving actividad", err)
		}

		// 3. Guardar en la doble capa de caché (L1 + L2) para futuras consultas
		if err := cache.SetJSON(cacheKey, actividad); err != nil {
			log.Warnf("Error guardando actividad en caché: %v", err)
		}
	}

	var actividadDto dto.ActividadDto
	actividadDto.Id = id
	log.Debugf("Nombre de la actividad: %v", actividad.Nombre)
	actividadDto.Nombre = actividad.Nombre
	actividadDto.Descripcion = actividad.Descripcion
	actividadDto.Profesor = actividad.Profesor
	actividadDto.OwnerId = actividad.OwnerId
	actividadDto.ImagenURL = actividad.ImagenURL
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
func InsertActividad(actividadDto dto.ActividadDto, token string) (dto.ActividadDto, e.ApiError) {
	var actividad model.Actividad

	err := ValidarActividadConcurrently(actividadDto)
	if err != nil {
		return dto.ActividadDto{}, e.NewBadRequestApiError(err.Error())
	}
	//if actividadDto.Nombre == "" {
	//	log.Error("El nombre de la actividad no puede estar vacío")
	//	return dto.ActividadDto{}, e.NewBadRequestApiError("El nombre de la actividad no puede estar vacío")
	//}
	//if len(actividadDto.Horario) == 0 {
	//	log.Error("La actividad debe tener al menos un horario")
	//	return dto.ActividadDto{}, e.NewBadRequestApiError("La actividad debe tener al menos un horario")
	//}
	//if actividadDto.Profesor == "" {
	//	log.Error("El nombre del profesor no puede estar vacío")
	//	return dto.ActividadDto{}, e.NewBadRequestApiError("El nombre del profesor no puede estar vacío")
	//}
	//if actividadDto.OwnerId == 0 {
	//	log.Error("El owner_id es requerido")
	//	return dto.ActividadDto{}, e.NewBadRequestApiError("owner_id is required")
	//}

	// VALIDAR EXISTENCIA DEL OWNER CONTRA API_USUARIOS
	log.Infof("Validating owner user %d against API_Usuarios", actividadDto.OwnerId)
	if err := usuarios.ValidateUser(actividadDto.OwnerId, token); err != nil {
		log.Errorf("Owner validation failed for user %d: %v", actividadDto.OwnerId, err)
		return dto.ActividadDto{}, e.NewBadRequestApiError(fmt.Sprintf("Invalid owner_id: %v", err))
	}

	actividad.Nombre = actividadDto.Nombre
	actividad.Descripcion = actividadDto.Descripcion
	actividad.Profesor = actividadDto.Profesor
	actividad.OwnerId = actividadDto.OwnerId
	actividad.ImagenURL = actividadDto.ImagenURL

	for _, horarioDto := range actividadDto.Horario {
		horario := model.Horario{
			Dia:        horarioDto.Dia,
			HoraInicio: horarioDto.HoraInicio,
			HoraFin:    horarioDto.HoraFin,
			Cupo:       horarioDto.Cupo,
		}
		actividad.Horarios = append(actividad.Horarios, horario)
	}

	actividadInsertada, err := actividadRepositories.InsertActividad(actividad)
	if err != nil {
		log.Errorf("Error al insertar actividad: %v", err)
		return dto.ActividadDto{}, e.NewInternalServerApiError("Error al insertar actividad", err)
	}

	// Publicar en caché
	cacheKey := fmt.Sprintf("actividad:%s", actividadDto.Id)
	cache.SetJSON(cacheKey, actividadInsertada)

	log.Infof("Actividad %s indexada en caché L1+L2", actividadDto.Id)

	actividadDto.Id = actividadInsertada.Id.Hex()

	// Publicar evento CREATE en RabbitMQ
	if err := queue.PublishEvent(queue.EventCreate, actividadDto.Id); err != nil {
		log.Errorf("Error al publicar evento CREATE en RabbitMQ: %v", err)
		// No retornamos error porque la actividad ya fue creada exitosamente
	}

	return actividadDto, nil // Mandar aca el error de la cache en caso de que haya?? ==============================^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
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
	// Invalida caché L1+L2
	cacheKey := fmt.Sprintf("actividad:%s", id)
	cache.Delete(cacheKey)

	// Publicar evento DELETE en RabbitMQ
	if err := queue.PublishEvent(queue.EventDelete, id); err != nil {
		log.Errorf("Error al publicar evento DELETE en RabbitMQ: %v", err)
		// No retornamos error porque la actividad ya fue eliminada exitosamente
	}

	return nil
}

func UpdateActividad(actividadDto dto.ActividadDto, token string) (dto.ActividadDto, e.ApiError) {
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

	// 1.2 VALIDACIÓN DE SEGURIDAD (JWT)
	claims, err := utils.ParseJWT(token)
	if err != nil {
		log.Errorf("Security check failed: Invalid token: %v", err)
		return dto.ActividadDto{}, e.NewUnauthorizedApiError("Sesión inválida o expirada")
	}

	userIdStr := claims.Subject
	var currentUserId int
	fmt.Sscanf(userIdStr, "%d", &currentUserId)

	if !claims.Es_admin && currentUserId != actividadActual.OwnerId {
		log.Warnf("Unauthorized edit attempt by user %d on activity %s (owner: %d)", currentUserId, actividadDto.Id, actividadActual.OwnerId)
		return dto.ActividadDto{}, e.NewForbiddenApiError("No tienes permisos para editar esta actividad")
	}

	// 1.5 VALIDAR OWNER SI SE PROPORCIONA (para cambio de propietario)
	if actividadDto.OwnerId != 0 && actividadDto.OwnerId != actividadActual.OwnerId {
		log.Infof("Validating new owner user %d against API_Usuarios", actividadDto.OwnerId)
		if err := usuarios.ValidateUser(actividadDto.OwnerId, token); err != nil {
			log.Errorf("Owner validation failed for user %d: %v", actividadDto.OwnerId, err)
			return dto.ActividadDto{}, e.NewBadRequestApiError(fmt.Sprintf("Invalid owner_id: %v", err))
		}
		actividadActual.OwnerId = actividadDto.OwnerId
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
	if actividadDto.ImagenURL != "" {
		actividadActual.ImagenURL = actividadDto.ImagenURL
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

	// Actualiza caché L1+L2
	cacheKey := fmt.Sprintf("actividad:%s", actividadActual.Id.Hex())
	cache.SetJSON(cacheKey, actividadActual)

	log.Infof("Actividad %s actualizada en caché L1+L2", cacheKey)

	// 6. Armar respuesta DTO
	var actividadActualizada dto.ActividadDto
	actividadActualizada.Id = actividadActual.Id.Hex()
	actividadActualizada.Nombre = actividadActual.Nombre
	actividadActualizada.Descripcion = actividadActual.Descripcion
	actividadActualizada.Profesor = actividadActual.Profesor
	actividadActualizada.OwnerId = actividadActual.OwnerId
	actividadActualizada.ImagenURL = actividadActual.ImagenURL

	for _, h := range actividadActual.Horarios {
		actividadActualizada.Horario = append(actividadActualizada.Horario, dto.HorarioDto{
			Dia:        h.Dia,
			HoraInicio: h.HoraInicio,
			HoraFin:    h.HoraFin,
			Cupo:       h.Cupo,
		})
	}

	// Publicar evento UPDATE en RabbitMQ
	if err := queue.PublishEvent(queue.EventUpdate, actividadActualizada.Id); err != nil {
		log.Errorf("Error al publicar evento UPDATE en RabbitMQ: %v", err)
		// No retornamos error porque la actividad ya fue actualizada exitosamente
	}

	return actividadActualizada, nil
}

// CalcularDisponibilidad calcula la disponibilidad de horarios de forma concurrente
// Utiliza GoRoutines, Channels y WaitGroups según enunciado
func CalcularDisponibilidad(id string) (dto.DisponibilidadResponse, e.ApiError) {
	startTime := time.Now()

	// 1. Obtener la actividad
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error converting id to mongo ID: %v", err)
		return dto.DisponibilidadResponse{}, e.NewBadRequestApiError("Invalid ID format")
	}

	actividad, err := actividadRepositories.GetActividadById(objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dto.DisponibilidadResponse{}, e.NewNotFoundApiError("actividad not found")
		}
		return dto.DisponibilidadResponse{}, e.NewInternalServerApiError("Error retrieving actividad", err)
	}

	// 2. Configurar procesamiento concurrente
	numHorarios := len(actividad.Horarios)
	if numHorarios == 0 {
		return dto.DisponibilidadResponse{}, e.NewBadRequestApiError("La actividad no tiene horarios")
	}

	// Channel para recibir resultados de las GoRoutines
	resultsChan := make(chan dto.DisponibilidadResult, numHorarios)

	// WaitGroup para sincronizar las GoRoutines
	var wg sync.WaitGroup

	log.Infof("Iniciando cálculo concurrente de disponibilidad para %d horarios", numHorarios)

	// 3. Lanzar una GoRoutine por cada horario
	for i, horario := range actividad.Horarios {
		wg.Add(1)
		go calcularDisponibilidadHorario(i, horario, resultsChan, &wg)
	}

	// 4. GoRoutine para cerrar el channel cuando todas las goroutines terminen
	go func() {
		wg.Wait()
		close(resultsChan)
		log.Info("Todas las GoRoutines completadas, channel cerrado")
	}()

	// 5. Recolectar resultados del channel
	var resultados []dto.DisponibilidadResult
	totalCupo := 0
	totalOcupado := 0

	for result := range resultsChan {
		resultados = append(resultados, result)
		totalCupo += result.Cupo
		totalOcupado += (result.Cupo - result.Disponibles)
	}

	elapsed := time.Since(startTime)

	response := dto.DisponibilidadResponse{
		ActividadId:  id,
		Nombre:       actividad.Nombre,
		Horarios:     resultados,
		TotalCupo:    totalCupo,
		TotalOcupado: totalOcupado,
		Tiempo:       elapsed.String(),
	}

	log.Infof("Cálculo completado en %s: %d horarios procesados", elapsed, numHorarios)
	return response, nil
}

// calcularDisponibilidadHorario es ejecutada como GoRoutine para procesar cada horario
func calcularDisponibilidadHorario(index int, horario model.Horario, resultsChan chan<- dto.DisponibilidadResult, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Debugf("GoRoutine %d: Procesando horario %s", index, horario.Dia)

	// Simular cálculo complejo (en producción aquí consultarías inscripciones, etc.)
	time.Sleep(time.Millisecond * time.Duration(50+rand.Intn(100)))

	// Simular ocupación aleatoria para demostración
	// En producción, esto vendría de consultar inscripciones reales
	ocupados := rand.Intn(horario.Cupo + 1)
	disponibles := horario.Cupo - ocupados
	porcentajeOcupado := 0.0
	if horario.Cupo > 0 {
		porcentajeOcupado = float64(ocupados) / float64(horario.Cupo) * 100
	}

	result := dto.DisponibilidadResult{
		Dia:               horario.Dia,
		Cupo:              horario.Cupo,
		Disponibles:       disponibles,
		PorcentajeOcupado: porcentajeOcupado,
	}

	log.Debugf("GoRoutine %d: Completado - %s: %d/%d disponibles (%.1f%% ocupado)",
		index, horario.Dia, disponibles, horario.Cupo, porcentajeOcupado)

	// Enviar resultado por el channel
	resultsChan <- result
}

func BorrarCupo(id string, horarioId string) e.ApiError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error converting id to mongo ID: %v", err)
		return e.NewBadRequestApiError("Invalid ID format")
	}

	// horarioId format is expected to be "Dia-HoraInicio" (e.g., "Lunes-08:00")
	parts := strings.Split(horarioId, "-")
	if len(parts) < 2 {
		log.Errorf("Invalid horarioId format: %s", horarioId)
		return e.NewBadRequestApiError("Invalid horarioId format")
	}
	dia := parts[0]
	horaInicio := parts[1]
	log.Infof("Borrando cupo de la actividad:%s para el horario %s - %s en la BD", objectID.Hex(), dia, horaInicio)
	err = actividadRepositories.BorrarCupo(objectID, dia, horaInicio)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return e.NewNotFoundApiError("actividad or horario not found")
		}
		return e.NewInternalServerApiError("Error updating cupo", err)
	}
	// Invalida caché local
	log.Infof("Borrando de cache la actividad:%s", objectID.Hex())
	cacheKey := fmt.Sprintf("actividad:%s", objectID.Hex())
	cache.Delete(cacheKey)

	// Publicar evento UPDATE en RabbitMQ para sincronizar API_Busquedas (Solr y Cache)
	if err := queue.PublishEvent(queue.EventUpdate, id); err != nil {
		log.Errorf("Error al publicar evento UPDATE en RabbitMQ despues de borrar cupo: %v", err)
	}

	return nil
}
