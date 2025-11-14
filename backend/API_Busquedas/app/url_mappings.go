package app

import (
	searchController "api_busquedas/controllers/search"
	inscripcionController "api_busquedas/controllers/inscripcion"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Endpoint principal de búsqueda paginada con filtros y ordenamiento
	// Query params: q (query), page, size, sort, order, filters
	router.GET("/search/actividades", searchController.SearchActividades)

	// Endpoint para obtener actividad específica (con cache)
	router.GET("/actividad/:id", searchController.GetActividadById)

	// Endpoint para obtener inscripciones por usuario (solo consulta)
	// NOTA: La creación de inscripciones está en API_Usuarios
	router.GET("/inscripciones/usuario/:id", inscripcionController.GetInscripcionesByUsuarioId)

	// Health check endpoint
	router.GET("/health", searchController.HealthCheck)

	// Stats endpoint para monitoreo de cache
	router.GET("/stats", searchController.GetCacheStats)

	log.Info("API_Busquedas: Finishing mappings configurations")
}
