package h

import (
	"fmt"
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
	c.SetCookie(key, val, 365*3600, "/", config.Config.GetString("server.host"), false, true)
}
func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "/", config.Config.GetString("server.host"), false, true)
}

func Cors(context *gin.Context) {
	method := context.Request.Method
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	fmt.Println(context.GetHeader("Origin"))
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	context.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
