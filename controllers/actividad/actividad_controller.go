package actividadController

import (
	"arqsoft_proyecto/dto"
	service "arqsoft_proyecto/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetActividadById(c *gin.Context) {
	log.Debug("Actividad id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	
	var actividadDto dto.ActividadDto

	actividadDto, err := service.GetActividadById(id)


	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, actividadDto)
}