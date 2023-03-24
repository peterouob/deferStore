package serviceCart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logicCart "server/logic/cart"
	"server/server/service/h"
)

type AddRequest struct {
	GoodsId string `json:"goodsId"`
}

func Add(c *gin.Context) {

	var req AddRequest
	c.ShouldBindJSON(&req)
	tk := h.GetToken(c)
	fmt.Println(tk.Uid)
	_, err := logicCart.Add(tk.Uid, req.GoodsId, 1)
	h.OkFail(c, true, err)
}
