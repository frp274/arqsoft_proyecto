package app

import (
	actividadController "api_actividades/controllers/actividad"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.GET("/actividad/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)
	router.GET("/actividad", actividadController.GetAllActividades)
	router.DELETE("/actividad/:id", actividadController.DeleteActividad)
	router.PUT("/actividad/:id", actividadController.UpdateActividad)

	//router.POST("/inscripcion", inscripcionController.InscripcionActividad)
	//router.GET("/inscripcion/:id", inscripcionController.GetInscripcionesByUsuarioId)
//
	//router.POST("/login", usuarioController.Login)
	log.Info("Finishing mappings configurations")

}
