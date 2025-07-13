package inscripcion

import (
	"arqsoft_proyecto/dto"
	service "arqsoft_proyecto/services"
	"net/http"
	_"strconv"

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
