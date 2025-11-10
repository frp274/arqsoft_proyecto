package app

import (
	usuarioController "api_usuarios/controllers/usuario"
	"api_usuarios/middlewares"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Autenticación (sin protección)
	router.POST("/login", usuarioController.Login)

	// Gestión de usuarios
	// GET /usuario/:id - Requiere autenticación y ser owner o admin
	router.GET("/usuario/:id",
		middlewares.AuthMiddleware(),
		middlewares.RequireOwnerOrAdmin(),
		usuarioController.GetUsuarioById)

	// POST /usuario - Solo admins pueden crear usuarios
	router.POST("/usuario",
		middlewares.AuthMiddleware(),
		middlewares.RequireAdmin(),
		usuarioController.CreateUsuario)

	// Endpoints futuros (comentados por ahora)
	// router.PUT("/usuario/:id", middlewares.AuthMiddleware(), middlewares.RequireOwnerOrAdmin(), usuarioController.UpdateUsuario)
	// router.DELETE("/usuario/:id", middlewares.AuthMiddleware(), middlewares.RequireAdmin(), usuarioController.DeleteUsuario)

	log.Info("API_Usuarios: Finishing mappings configurations")
}
