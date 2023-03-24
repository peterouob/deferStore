package h

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/component/config"
	logicToken "server/logic/token"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func OKMessage(c *gin.Context, any interface{}) {
	c.JSON(http.StatusOK, gin.H{"msg": any})
}

func OkFail(c *gin.Context, body any, err error) {
	if err != nil {
		Fail(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "body": body})
	}
}

func FailMessage(c *gin.Context, any interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": any})
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

func GetToken(c *gin.Context) *logicToken.Token {
	value, exists := c.Get("token")
	if !exists {
		return nil
	} else {
		return value.(*logicToken.Token)
	}
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
