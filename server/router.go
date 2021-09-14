package server

import (
	"github.com/InsomniaCoder/go-redis-load/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("keys")
		{
			keys := new(controllers.KeysController)
			userGroup.GET("/", keys.Push)

		}
	}
	return router

}
