package app

import(
	actividadController "arqsoft_proyecto/controllers/actividad"

	log "github.com/sirupsen/logrus"
)


func mapUrls(){
	router.GET("/actividades/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)

	log.Info("Finishing mappings configurations")
}


