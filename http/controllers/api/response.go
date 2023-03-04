package api

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Error(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": msg})
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func ErrorServer(err error, c *gin.Context) {
	glog.Error(err)
	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal Server Error"})
}

func ErrorUnauthorized(msg string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": msg})
}

func ErrorNotFound(err error, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
}
