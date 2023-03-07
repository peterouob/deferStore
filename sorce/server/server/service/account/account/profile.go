package serviceAccount

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	c.String(http.StatusOK, "Welcome login")
}
