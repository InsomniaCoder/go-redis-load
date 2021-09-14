package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type KeysController struct{}

// Push stores 1 new key into redis
func (u KeysController) Push(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
