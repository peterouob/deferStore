package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"server/logic/orm/dal"
	"server/logic/orm/model"
	logictoken "server/logic/token"
	"server/server/service/h"
	"strconv"
	"time"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail(c, err)
		return
	}
	_, err := dal.Account.Where(dal.Account.Name.Eq(req.Name)).First()
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "-1",
			"msg":    "相同名稱已存在",
		})
	} else {
		fmt.Println("ok")
	}

	if req.NickName == "" || req.Name == "" || req.Password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "-1",
			"msg":  "請輸入完整訊息",
		})
		return
	}

	acc := model.Account{
		Name:     req.Name,
		Nickname: req.NickName,
		Password: req.Password,
	}
	err = dal.Account.Create(&acc)
	if err != nil {
		h.Fail(c, err)
	}
	token := logictoken.Token{
		Uid:      strconv.Itoa(int(acc.ID)),
		Nickname: acc.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UnixMilli(),
		},
	}

	signToken, err := logictoken.Sign(&token)
	if err != nil {
		h.Fail(c, err)
	} else {
		h.SetCookie(c, "token", signToken)
		fmt.Println(signToken)
		h.Ok(c)
	}
}
