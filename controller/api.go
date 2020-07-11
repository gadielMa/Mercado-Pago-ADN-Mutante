package controller

import (
	_ "github.com/gadielMa/test/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter inicializa las rutas
func SetupRouter(router *gin.Engine) {
	router.POST("/mutant", Mutant)
	router.GET("/stats", Stats)
	router.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
