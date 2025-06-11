package usuario

import (
	"arqsoft_proyecto/dto"
	usuariosService "arqsoft_proyecto/services"
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

	usuarioId, token, err := usuariosService.Login(request.Username, request.Password)
	if err != nil {
		log.Printf("el error esta aca")
		c.JSON(http.StatusForbidden, gin.H{"Error": err.Error()})
		return	
	}
	c.JSON(http.StatusOK, dto.LoginResponse{
		Id:    usuarioId,
		Token: token,
	})
}
