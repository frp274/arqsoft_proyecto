package inscripcion

import (
	"arqsoft_proyecto/dto"
	service "arqsoft_proyecto/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InscripcionActividad(c *gin.Context) {
	var inscripcionDto dto.InscripcionDto
	err := c.BindJSON(&inscripcionDto)

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
