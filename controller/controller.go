package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func missingParam(name string, c *gin.Context) bool {
	if c.PostForm(name) == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "Missing data POST : " + name})
		return true
	}

	return false
}
