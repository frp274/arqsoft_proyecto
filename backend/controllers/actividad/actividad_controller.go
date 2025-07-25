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

/*func GetAllActividades(c *gin.Context) {
	actividadesDto, err := service.GetAllActividades()

	if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }
	c.JSON(http.StatusOK, actividadesDto)
}*/

/*func GetAllActividades(c *gin.Context) {
	nombre := c.Query("nombre")
	actividadesDto, err := service.GetActividadesByNombre(nombre)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, actividadesDto)
}*/

func GetAllActividades(c *gin.Context) {
	nombre := c.Query("nombre")
	log.Infof(">> Filtro recibido en el controller: '%s'", nombre)

	actividadesDto, err := service.GetActividadesByNombre(nombre)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, actividadesDto)
}

func InsertActividad(c *gin.Context) {
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

func DeleteActividad(c *gin.Context) {
	log.Debug("Actividad id to delete: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.DeleteActividad(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actividad eliminada correctamente"})
}

/*func UpdateActividad(c *gin.Context) {
	log.Debug("Id de la actividad a modificar: " + c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Error("ID inválido: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var actividadDto dto.ActividadDto
	if err := c.BindJSON(&actividadDto); err != nil {
		log.Error("Error al parsear JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Aseguramos que el ID del DTO sea el mismo que el del path
	actividadDto.Id = id

	updatedActividad, serviceErr := service.UpdateActividad(actividadDto)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	c.JSON(http.StatusOK, updatedActividad)
}*/

func UpdateActividad(c *gin.Context) {
	log.Debug("Actividad id to update: " + c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Error("ID inválido: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var actividadDto dto.ActividadDto
	err = c.BindJSON(&actividadDto)
	if err != nil {
		log.Error("Error al parsear el cuerpo JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	actividadDto.Id = id

	actividadActualizada, updateErr := service.UpdateActividad(actividadDto)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	c.JSON(http.StatusOK, actividadActualizada)
}
