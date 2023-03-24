package logicCart

import (
	"encoding/json"
)

var (
	RedisPrefix = "cart"
)

type Cart struct {
	GoodsId string `json:"goods_id"`
	Numbers int    `json:"numbers"`
}

func (c *Cart) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Cart) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *Cart) FromRedis(data interface{}) error {
	str := data.(string)
	return json.Unmarshal([]byte(str), c)
}

func KeyCard(uid string) string {
	return RedisPrefix + "." + uid
}
func KeyCardItem(uid, goodsId string) string {
	return RedisPrefix + "." + uid + "." + goodsId
}
