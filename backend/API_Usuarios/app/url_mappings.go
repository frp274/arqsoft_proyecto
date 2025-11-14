package app

import (
	inscripcionController "api_usuarios/controllers/inscripcion"
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

	// ============================================
	// INSCRIPCIONES - Gestión de inscripciones a actividades
	// ============================================
	
	// POST /inscripcion - Crear nueva inscripción (requiere autenticación)
	router.POST("/inscripcion",
		middlewares.AuthMiddleware(),
		inscripcionController.CreateInscripcion)
	
	// GET /inscripciones/usuario/:id - Obtener inscripciones de un usuario
	router.GET("/inscripciones/usuario/:id",
		middlewares.AuthMiddleware(),
		middlewares.RequireOwnerOrAdmin(),
		inscripcionController.GetInscripcionesByUsuarioId)
	
	// DELETE /inscripcion/:id - Eliminar inscripción
	router.DELETE("/inscripcion/:id",
		middlewares.AuthMiddleware(),
		inscripcionController.DeleteInscripcion)

	// Endpoints futuros (comentados por ahora)
	//router.PUT("/usuario/:id",
	//	middlewares.AuthMiddleware(),
	//	middlewares.RequireOwnerOrAdmin(),
	//	usuarioController.Update)
	// router.DELETE("/usuario/:id", middlewares.AuthMiddleware(), middlewares.RequireAdmin(), usuarioController.DeleteUsuario)

	log.Info("API_Usuarios: Finishing mappings configurations")
}
