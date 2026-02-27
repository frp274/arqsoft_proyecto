package inscripcion

import (
	"api_busquedas/dto"
	service "api_busquedas/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InscripcionActividad(c *gin.Context) {
	var inscripcionDto dto.InscripcionDto
	//id, _ := strconv.Atoi(c.Param("actividad_id"))
	err := c.BindJSON(&inscripcionDto)
	//inscripcionDto.ActividadId = id
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	inscripcionDto, er := service.InscripcionActividad(inscripcionDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, inscripcionDto)
}

func GetInscripcionesByUsuarioId(c *gin.Context) {
	idParam := c.Param("id")
	log.Debug("Usuario id to load: " + idParam)

	// Manejar formato defensivo (ej: "1:1" -> "1")
	if strings.Contains(idParam, ":") {
		idParam = strings.Split(idParam, ":")[0]
		log.Warnf("ID de usuario con formato inválido detectado, usando: %s", idParam)
	}

	id, atoiErr := strconv.Atoi(idParam)
	if atoiErr != nil {
		log.Errorf("Error al convertir ID '%s': %v", idParam, atoiErr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	actividadesDto, err := service.GetInscripcionesByUsuarioId(id)

	if err != nil {
		log.Error(err.Error())
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, actividadesDto)
}
