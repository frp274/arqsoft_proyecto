package inscripcion

import (
	"api_usuarios/dto"
	"api_usuarios/services"
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CreateInscripcion maneja la creación de una nueva inscripción
func CreateInscripcion(c *gin.Context) {
	var inscripcionDto dto.InscripcionDto
	
	// Parsear JSON del body
	if err := c.ShouldBindJSON(&inscripcionDto); err != nil {
		log.Errorf("Error al parsear JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	
	log.Infof("Solicitud de inscripción - Usuario: %d, Actividad: %d", 
		inscripcionDto.UsuarioId, inscripcionDto.ActividadId)
	
	// Crear inscripción
	inscripcionDto, err := services.CreateInscripcion(inscripcionDto)
	if err != nil {
		c.JSON(err.Status(), gin.H{"error": err.Message()})
		return
	}
	
	// Retornar inscripción creada
	c.JSON(http.StatusCreated, inscripcionDto)
}

// GetInscripcionesByUsuarioId obtiene todas las inscripciones de un usuario
func GetInscripcionesByUsuarioId(c *gin.Context) {
	usuarioIdStr := c.Param("id")
	usuarioId, err := strconv.Atoi(usuarioIdStr)
	
	if err != nil {
		log.Errorf("ID de usuario inválido: %s", usuarioIdStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}
	
	log.Infof("Obteniendo inscripciones del usuario: %d", usuarioId)
	
	// Obtener inscripciones
	inscripcionesDto, apiErr := services.GetInscripcionesByUsuarioId(usuarioId)
	if apiErr != nil {
		c.JSON(apiErr.Status(), gin.H{"error": apiErr.Message()})
		return
	}
	
	// Retornar inscripciones
	c.JSON(http.StatusOK, inscripcionesDto)
}

// DeleteInscripcion elimina una inscripción
func DeleteInscripcion(c *gin.Context) {
	inscripcionIdStr := c.Param("id")
	inscripcionId, err := strconv.Atoi(inscripcionIdStr)
	
	if err != nil {
		log.Errorf("ID de inscripción inválido: %s", inscripcionIdStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de inscripción inválido"})
		return
	}
	
	log.Infof("Eliminando inscripción: %d", inscripcionId)
	
	// Eliminar inscripción
	apiErr := services.DeleteInscripcion(inscripcionId)
	if apiErr != nil {
		c.JSON(apiErr.Status(), gin.H{"error": apiErr.Message()})
		return
	}
	
	// Retornar éxito
	c.JSON(http.StatusOK, gin.H{"message": "Inscripción eliminada exitosamente"})
}
