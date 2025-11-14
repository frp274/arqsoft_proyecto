package usuario

import (
	"api_usuarios/dto"
	"api_usuarios/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuarioId, token, es_admin, err := services.Login(request.Username, request.Password)
	if err != nil {
		log.Errorf("Login failed for user %s: %v", request.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	log.Infof("User %s logged in successfully (id: %d, admin: %t)", request.Username, usuarioId, es_admin)
	c.JSON(http.StatusOK, dto.LoginResponse{
		Id:      usuarioId,
		Token:   token,
		EsAdmin: es_admin,
	})
}

func GetUsuarioById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Warnf("Invalid user ID format: %s", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	usuarioDto, err := services.GetUsuarioById(id)
	if err != nil {
		log.Errorf("Error getting user %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	log.Infof("User %d retrieved successfully", id)
	c.JSON(http.StatusOK, usuarioDto)
}

func CreateUsuario(c *gin.Context) {
	var request dto.CreateUsuarioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Warnf("Invalid request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si se intenta crear un admin
	if request.EsAdmin {
		// Verificar si el usuario que hace la petición es admin autenticado
		esAdmin, exists := c.Get("es_admin")
		if !exists || !esAdmin.(bool) {
			log.Warnf("Unauthorized attempt to create admin user by %s", request.Username)
			c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear nuevos administradores"})
			return
		}
		log.Infof("Admin user creating new admin: %s", request.Username)
	}

	usuarioDto, err := services.CreateUsuario(request)
	if err != nil {
		log.Errorf("Error creating user %s: %v", request.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("User created successfully: %s (ID: %d)", usuarioDto.Username, usuarioDto.Id)
	c.JSON(http.StatusCreated, usuarioDto)
}
