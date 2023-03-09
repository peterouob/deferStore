package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/logic/orm/dal"
	logicToken "server/logic/token"
	"server/server/service/h"
	"strconv"
	"time"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token" `
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return
	}
	acc, err := dal.Account.Where(dal.Account.Name.Eq(request.Name)).First()
	if err != nil {
		h.Fail(c, err)
		return
	}
	if len(acc.Password) > 0 && acc.Password == request.Password {
		token := logicToken.Token{
			Uid:      strconv.Itoa(int(acc.ID)),
			Nickname: acc.Nickname,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
			},
		}
		if sign, err := logicToken.Sign(&token); err == nil {
			h.SetCookie(c, "token", sign)
			fmt.Println(sign)
			h.Ok(c)
		} else {
			h.Fail(c, err)
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "登入有誤",
		})
	}
}
