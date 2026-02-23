package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(AllowCORS)
}

func StartRoute() {
	mapUrls()

	log.Info("Starting API_Usuarios server on port 8082")
	router.Run(":8082")
}

func AllowCORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if ctx.Request.Method == http.MethodOptions {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.Next()
}
