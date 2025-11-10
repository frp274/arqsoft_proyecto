package usuario

import (
	"api_busquedas/dto"
	usuariosService "api_busquedas/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	usuarioId, token, es_admin, err := usuariosService.Login(request.Username, request.Password)
	if err != nil {
		log.Printf("el error esta aca")
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return	
	}
	c.JSON(http.StatusOK, dto.LoginResponse{
		Id:    usuarioId,
		Token: token,
		Es_admin: es_admin,
	})
	log.Printf("id: %d", usuarioId)
	log.Printf("id: %s", token)
	log.Printf("id: %t", es_admin)
}
