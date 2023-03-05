package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.String(http.StatusOK, "Welcome login")
}
