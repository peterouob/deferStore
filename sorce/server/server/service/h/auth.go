package h

import (
	"github.com/gin-gonic/gin"
	logictoken "server/logic/token"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenSign, err := c.Cookie("token")
		if err != nil {
			Forbidden(c)
			return
		}
		token, err := logictoken.Parse(tokenSign)
		if err != nil {
			Forbidden(c)
			return
		}
		c.Set("token", token)
	}
}
