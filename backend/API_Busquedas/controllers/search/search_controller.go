package search

import (
	"fmt"
	"net/http"

	"api_busquedas/cache"
	"api_busquedas/clients"
	"api_busquedas/model"
	"api_busquedas/search"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vanng822/go-solr/solr"
)

// SearchActividades handles paginated search with filters
func SearchActividades(c *gin.Context) {
	var params model.SearchParams

	// Set defaults
	params.Page = 1
	params.PageSize = 10

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate and set defaults
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	if params.Query == "" {
		params.Query = "*:*" // Match all
	}
	if params.Order != "asc" && params.Order != "desc" {
		params.Order = "asc"
	}

	// Generate cache key
	cacheKey := fmt.Sprintf("search:%s:p%d:s%d:sort%s:%s",
		params.Query, params.Page, params.PageSize, params.Sort, params.Order)

	// Try cache first
	var result model.ActividadSearchResult
	if cache.GetJSON(cacheKey, &result) {
		log.Infof("Returning cached search results for: %s", params.Query)
		c.JSON(http.StatusOK, result)
		return
	}

	// Build Solr query
	query := solr.NewQuery()
	query.Q(params.Query)
	query.Rows((params.Page - 1) * params.PageSize)
	query.Start(params.PageSize)

	// Add sorting
	if params.Sort != "" {
		sortField := params.Sort + " " + params.Order
		query.Sort(sortField)
	}

	// Execute search
	si := search.SolrClient
	searchObj := si.Search(query)
	response, err := searchObj.Result(nil)
	if err != nil {
		log.Errorf("Solr search failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	// Parse results
	var actividades []model.ActividadSolr
	for _, doc := range response.Results.Docs {
		actividad := model.ActividadSolr{
			ID:          getString(doc, "id"),
			Nombre:      getString(doc, "nombre"),
			Descripcion: getString(doc, "descripcion"),
			Profesor:    getString(doc, "profesor"),
			Horarios:    getString(doc, "horarios"),
		}

		if tags, ok := doc["tags"].([]interface{}); ok {
			for _, tag := range tags {
				if str, ok := tag.(string); ok {
					actividad.Tags = append(actividad.Tags, str)
				}
			}
		}

		actividades = append(actividades, actividad)
	}

	// Build result
	totalResults := int64(response.Results.NumFound)
	totalPages := int(totalResults) / params.PageSize
	if int(totalResults)%params.PageSize > 0 {
		totalPages++
	}

	result = model.ActividadSearchResult{
		Actividades: actividades,
		Total:       totalResults,
		Page:        params.Page,
		PageSize:    params.PageSize,
		TotalPages:  totalPages,
	}

	// Cache the result
	if err := cache.SetJSON(cacheKey, result); err != nil {
		log.Warnf("Failed to cache search results: %v", err)
	}

	log.Infof("Search completed: %d results for query '%s'", len(actividades), params.Query)
	c.JSON(http.StatusOK, result)
}

// GetActividadById retrieves a single actividad with caching
func GetActividadById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	cacheKey := fmt.Sprintf("actividad:%s", id)

	// Try cache first
	var actividad clients.Actividad
	if cache.GetJSON(cacheKey, &actividad) {
		log.Debugf("Returning cached actividad: %s", id)
		c.JSON(http.StatusOK, actividad)
		return
	}

	// Fetch from API_Actividades
	actividadPtr, err := clients.GetActividadFromAPI(id)
	if err != nil {
		log.Errorf("Failed to fetch actividad %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Actividad not found"})
		return
	}

	// Cache the result
	if err := cache.SetJSON(cacheKey, actividadPtr); err != nil {
		log.Warnf("Failed to cache actividad: %v", err)
	}

	c.JSON(http.StatusOK, actividadPtr)
}

// HealthCheck returns the health status
func HealthCheck(c *gin.Context) {
	health := gin.H{
		"status": "healthy",
		"solr":   "connected",
		"cache":  "active",
	}

	// Test Solr connection
	query := solr.NewQuery()
	query.Q("*:*")
	query.Rows(0)

	si := search.SolrClient
	searchObj := si.Search(query)
	if _, err := searchObj.Result(nil); err != nil {
		health["solr"] = "disconnected"
		health["status"] = "degraded"
	}

	statusCode := http.StatusOK
	if health["status"] == "degraded" {
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, health)
}

// GetCacheStats returns cache statistics
func GetCacheStats(c *gin.Context) {
	stats := cache.GetStats()
	c.JSON(http.StatusOK, stats)
}

// Helper function
func getString(doc map[string]interface{}, key string) string {
	if val, ok := doc[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}
