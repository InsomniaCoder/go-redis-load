package controllers

import (
	"github.com/InsomniaCoder/go-redis-load/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type KeysController struct{}

// Push stores 1 new key into redis
func (u KeysController) Push(c *gin.Context) {

	err := redis.Set(c, "test","test")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to persist key", "error": err})
		c.Abort()
		return
	}
	c.String(http.StatusOK, "ok")
	return
}
