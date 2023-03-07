package h

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/component/config"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "尚無權限"})
}

func SetCookie(c *gin.Context, key, val string) {
	c.SetCookie(key, val, 365*3600, "/", config.Config.GetString("server.host"), false, false)
}
func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "/", config.Config.GetString("server.host"), false, true)
}
