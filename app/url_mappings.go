package app

import(
	actividadController "arqsoft_proyecto/controllers/actividad"
	inscripcionController "arqsoft_proyecto/controllers/inscripcion"
	log "github.com/sirupsen/logrus"
)


func mapUrls(){
	router.GET("/actividades/:id", actividadController.GetActividadById)
	router.POST("/actividad", actividadController.InsertActividad)

	router.POST("/inscripcion", inscripcionController.InscripcionActividad)
	
	log.Info("Finishing mappings configurations")
}


