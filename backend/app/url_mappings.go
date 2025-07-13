package app

import (
	actividadController "arqsoft_proyecto/controllers/actividad"
	inscripcionController "arqsoft_proyecto/controllers/inscripcion"
	usuarioController "arqsoft_proyecto/controllers/usuario"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	router.GET("/actividad/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)
	router.GET("/actividad", actividadController.GetAllActividades)

	router.POST("/inscripcion", inscripcionController.InscripcionActividad)

	router.POST("/login", usuarioController.Login)
	log.Info("Finishing mappings configurations")
}
