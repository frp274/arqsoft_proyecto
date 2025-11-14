package app

import (
	usuarioController "api_usuarios/controllers/usuario"
	"api_usuarios/middlewares"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Autenticación (sin protección - acceso público)
	router.POST("/login", usuarioController.Login)
	
	// Registro público (sin protección - solo usuarios normales)
	// Si intentan crear admin sin autenticación, será rechazado en el controller
	router.POST("/usuario", usuarioController.CreateUsuario)

	// Gestión de usuarios (protegidos)
	// GET /usuario/:id - Requiere autenticación y ser owner o admin
	router.GET("/usuario/:id",
		middlewares.AuthMiddleware(),
		middlewares.RequireOwnerOrAdmin(),
		usuarioController.GetUsuarioById)

	// POST /admin/usuario - Solo admins pueden crear nuevos admins
	router.POST("/admin/usuario",
		middlewares.AuthMiddleware(),
		middlewares.RequireAdmin(),
		usuarioController.CreateUsuario)

	// Endpoints futuros (comentados por ahora)
	// router.PUT("/usuario/:id", middlewares.AuthMiddleware(), middlewares.RequireOwnerOrAdmin(), usuarioController.UpdateUsuario)
	// router.DELETE("/usuario/:id", middlewares.AuthMiddleware(), middlewares.RequireAdmin(), usuarioController.DeleteUsuario)

	log.Info("API_Usuarios: Finishing mappings configurations")
}
