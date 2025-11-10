package usuario

import (
	"api_usuarios/dto"
	usuariosService "api_usuarios/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	usuarioId, token, es_admin, err := usuariosService.Login(request.Username, request.Password)
	if err != nil {
		log.Errorf("Login failed for user %s: %v", request.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	}

	log.Infof("User %s logged in successfully (id: %d, admin: %t)", request.Username, usuarioId, es_admin)
	c.JSON(http.StatusOK, dto.LoginResponse{
		Id:       usuarioId,
		Token:    token,
		Es_admin: es_admin,
	})
}
