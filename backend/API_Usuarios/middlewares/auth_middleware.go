package middlewares

import (
	"api_usuarios/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// AuthMiddleware valida el token JWT en el header Authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Warn("Missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Verificar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Warn("Invalid Authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Use: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validar y parsear el token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			log.Warnf("Invalid token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Guardar claims en el contexto para uso posterior
		c.Set("userId", claims.ID)
		c.Set("es_admin", claims.Es_admin)
		
		log.Debugf("User authenticated: ID=%s, Admin=%t", claims.ID, claims.Es_admin)
		c.Next()
	}
}

// RequireAdmin middleware que requiere que el usuario sea administrador
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		esAdmin, exists := c.Get("es_admin")
		if !exists {
			log.Error("es_admin not found in context - AuthMiddleware may not be applied")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication context error"})
			c.Abort()
			return
		}

		if !esAdmin.(bool) {
			log.Warnf("Access denied: user is not admin")
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin privileges required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireOwnerOrAdmin middleware que requiere que el usuario sea el dueño del recurso o admin
func RequireOwnerOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener ID del recurso desde la URL
		resourceId := c.Param("id")
		
		// Obtener datos del usuario autenticado del contexto
		userId, userExists := c.Get("userId")
		esAdmin, adminExists := c.Get("es_admin")

		if !userExists || !adminExists {
			log.Error("Authentication context not found - AuthMiddleware may not be applied")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication context error"})
			c.Abort()
			return
		}

		// Si es admin, permitir acceso
		if esAdmin.(bool) {
			log.Debugf("Access granted: user is admin")
			c.Next()
			return
		}

		// Si es el owner del recurso, permitir acceso
		if userId.(string) == resourceId {
			log.Debugf("Access granted: user is owner of resource")
			c.Next()
			return
		}

		log.Warnf("Access denied: user %s tried to access resource %s", userId, resourceId)
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only modify your own resources"})
		c.Abort()
	}
}
