package goodsService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/logic/orm/dal"
	"server/server/service/h"
)

type GetRequest struct {
	Id string `json:"id"`
}

func Get(c *gin.Context) {
	req := GetRequest{}
	c.ShouldBindJSON(&req)
	goods, err := dal.Good.Where(dal.Good.ID.Eq(req.Id)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		h.FailMessage(c, "商品不存在")
		return
	} else if err != nil {
		h.Fail(c, err)
		return
	}
	h.OKMessage(c, &HomePageFeed{
		Id:            goods.ID,
		Name:          goods.Name,
		Img:           goods.Img,
		Price:         goods.Price.String(),
		OriginalPrice: goods.OriginalPrice.String(),
		Desc:          goods.Desc,
	})
}
