package goodsService

import (
	"github.com/gin-gonic/gin"
	"server/logic/orm/dal"
	"server/logic/orm/model"
	"server/server/helper"
	"server/server/service/h"
)

type HomePageFeed struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Img           string `json:"img"`
	Price         string `json:"price"`
	OriginalPrice string `json:"original_price"`
	Desc          string `json:"desc"`
}

type HomePageResponse struct {
	Block map[string][]*HomePageFeed `json:"block"`
}

func HomePage(c *gin.Context) {
	blocks, err := dal.Block.Find()
	if err != nil {
		panic(err)
	}
	resp := HomePageResponse{Block: make(map[string][]*HomePageFeed)}

	var ids []string
	for _, block := range blocks {
		ids = append(ids, block.GoodID)
		//fmt.Println(block.GoodID)
	}
	ids = helper.RemoveDuplicateInt64(ids)

	items, err := dal.Good.Where(dal.Good.ID.In(ids...)).Find()
	if err != nil {
		panic(err)
	}
	m := make(map[string]*model.Good)
	for _, item := range items {
		m[item.ID] = item
	}
	for _, block := range blocks {
		if goods, ok := m[block.GoodID]; ok {
			resp.Block[block.Key] = append(resp.Block[block.Key], &HomePageFeed{
				Id:            goods.ID,
				Name:          goods.Name,
				Img:           goods.Img,
				Price:         goods.Price.String(),
				OriginalPrice: goods.OriginalPrice.String(),
				Desc:          goods.Desc,
			})
		}
	}
	h.OKMessage(c, &resp)
}
