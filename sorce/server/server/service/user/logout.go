package user

import (
	"github.com/gin-gonic/gin"
	"server/server/service/h"
)

type LogoutRequest struct{}
type LogoutResponse struct{}

func Logout(c *gin.Context) {
	h.RemoveCookie(c, "token")
	h.Ok(c)
}
