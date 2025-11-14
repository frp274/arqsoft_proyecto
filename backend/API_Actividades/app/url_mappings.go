package app

import (
	actividadController "api_actividades/controllers/actividad"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Operaciones CRUD de actividades
	router.GET("/actividad/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)
	router.PUT("/actividad/:id", actividadController.UpdateActividad)
	router.DELETE("/actividad/:id", actividadController.DeleteActividad)

	// Endpoint de acción con procesamiento concurrente
	router.POST("/actividad/:id/calcular-disponibilidad", actividadController.CalcularDisponibilidad)
	router.POST("/actividad/:id/borar-cupo", actividadController.BorrarCupo)
	// Endpoints futuros (comentados)
	//router.GET("/actividad", actividadController.GetActividadesByNombre)
	//router.POST("/inscripcion", inscripcionController.InscripcionActividad)
	//router.GET("/inscripcion/:id", inscripcionController.GetInscripcionesByUsuarioId)
	//router.POST("/login", usuarioController.Login)
	
	log.Info("Finishing mappings configurations")
}
