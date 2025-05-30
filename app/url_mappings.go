package app

import (
	actividadController "arqsoft_proyecto/controllers/actividad"
	inscripcionController "arqsoft_proyecto/controllers/inscripcion"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.GET("/actividad/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)
	router.GET("/actividad", actividadController.GetAllActividades)
	router.POST("/inscripcion", inscripcionController.InscripcionActividad)

	log.Info("Finishing mappings configurations")
}
