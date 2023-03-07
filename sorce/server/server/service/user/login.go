package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/logic/orm/dal"
	logictoken "server/logic/token"
	"server/server/service/h"
	"strconv"
	"time"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	acc, err := dal.Account.Where(dal.Account.Name.Eq(req.Name)).First()
	if err != nil {
		return
	}

	if len(acc.Password) > 0 && acc.Password == req.Password {
		token := logictoken.Token{
			Uid:      strconv.Itoa(int(acc.ID)),
			Nickname: acc.Nickname,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(7 * 24 * 3066).Unix(),
			},
		}
		if sign, err := logictoken.Sign(&token); err != nil {
			h.SetCookie(c, "token", sign)
			fmt.Println(token)
			h.Ok(c)
		} else {
			h.Fail(c, err)
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "登入錯誤",
		})
	}
}
