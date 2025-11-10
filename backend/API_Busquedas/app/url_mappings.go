package app

import (
	searchController "api_busquedas/controllers/search"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Endpoint principal de búsqueda paginada con filtros y ordenamiento
	// Query params: q (query), page, size, sort, order, filters
	router.GET("/search/actividades", searchController.SearchActividades)

	// Endpoint para obtener actividad específica (con cache)
	router.GET("/actividad/:id", searchController.GetActividadById)

	// Health check endpoint
	router.GET("/health", searchController.HealthCheck)

	// Stats endpoint para monitoreo de cache
	router.GET("/stats", searchController.GetCacheStats)

	log.Info("API_Busquedas: Finishing mappings configurations")
}
