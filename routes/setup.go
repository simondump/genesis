package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/simonwep/genisis/core"
)

func SetupRoutes() *gin.Engine {

	// Set mode
	gin.SetMode(core.Config().GinMode)

	// Create router
	router := gin.New()

	// Auth endpoints
	router.POST("/login", Login)
	router.POST("/register", Register)

	// Data endpoints
	router.PUT("/data/:key", SetData)
	router.DELETE("/data/:key", DeleteData)
	router.GET("/data/:key", DataByKey)
	router.GET("/data", Data)

	return router
}
