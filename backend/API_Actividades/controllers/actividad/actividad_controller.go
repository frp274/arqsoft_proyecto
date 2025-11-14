package actividadController

import (
	"api_actividades/dto"
	service "api_actividades/services"

	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetActividadById(c *gin.Context) {
	log.Debug("Actividad id to load: " + c.Param("id"))
	id:= c.Param("id")

	var actividadDto dto.ActividadDto

	actividadDto, err := service.GetActividadById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, actividadDto)
}


/*
func GetActividadesByNombre(c *gin.Context) {
	nombre := c.Query("nombre")
	log.Infof(">> Filtro recibido en el controller: '%s'", nombre)

	actividadesDto, err := service.GetActividadesByNombre(nombre)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, actividadesDto)
}
*/

func InsertActividad(c *gin.Context) {
	var actividadDto dto.ActividadDto

	err := c.BindJSON(&actividadDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Extraer token del header Authorization
	token := ""
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	actividadDto, er := service.InsertActividad(actividadDto, token)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, actividadDto)
}

func DeleteActividad(c *gin.Context) {
	log.Debug("Actividad id to delete: " + c.Param("id"))
	id:= c.Param("id")

	err := service.DeleteActividad(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Actividad eliminada correctamente"})
}


func UpdateActividad(c *gin.Context) {
	log.Debug("Actividad id to update: " + c.Param("id"))
	id := c.Param("id")

	var actividadDto dto.ActividadDto
	err := c.BindJSON(&actividadDto)
	if err != nil {
		log.Error("Error al parsear el cuerpo JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	actividadDto.Id = id

	// Extraer token del header Authorization
	token := ""
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	actividadActualizada, updateErr := service.UpdateActividad(actividadDto, token)
	if updateErr != nil {
		c.JSON(updateErr.Status(), updateErr)
		return
	}

	c.JSON(http.StatusOK, actividadActualizada)
}

// CalcularDisponibilidad endpoint de acción que calcula disponibilidad con concurrencia
func CalcularDisponibilidad(c *gin.Context) {
	log.Debug("Calculando disponibilidad para actividad: " + c.Param("id"))
	id := c.Param("id")

	resultado, err := service.CalcularDisponibilidad(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resultado)
}

func BorrarCupo (c *gin.Context) {
	log.Debug("Borrando cupo para actividad: " + c.Param("id"))
	id := c.Param("id")

	err := service.BorrarCupo(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cupo borrado correctamente"})
}