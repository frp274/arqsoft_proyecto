package app

import (
	usuarioController "api_usuarios/controllers/usuario"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Autenticación
	router.POST("/login", usuarioController.Login)

	// Gestión de usuarios (endpoints a implementar)
	// router.GET("/usuario/:id", usuarioController.GetUsuarioById)
	// router.POST("/usuario", usuarioController.CreateUsuario)
	// router.PUT("/usuario/:id", usuarioController.UpdateUsuario)
	// router.DELETE("/usuario/:id", usuarioController.DeleteUsuario)

	log.Info("API_Usuarios: Finishing mappings configurations")
}
