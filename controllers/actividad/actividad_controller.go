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

func InsertActividad(c *gin.Context){
	var actividadDto dto.ActividadDto
	err := c.BindJSON(&actividadDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	actividadDto, er := service.InsertActividad(actividadDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, actividadDto)
}