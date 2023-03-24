package logicCart

import (
	"fmt"
	"server/component/redis"
)

func Add(uid, goodsId string, number int) (*Cart, error) {
	var cart *Cart
	get := redis.Get(KeyCardItem(uid, goodsId))

	if len(get.Val()) > 0 {
		cart = &Cart{}
		err := get.Scan(cart)
		if err != nil {
			return nil, err
		}
	}

	if cart == nil {
		cart = &Cart{
			GoodsId: goodsId,
			Numbers: 0,
		}
	}
	cart.Numbers += number
	set := redis.Set(KeyCardItem(uid, goodsId), cart, 0)
	fmt.Println(KeyCardItem(uid, goodsId), set.Err())
	redis.SAdd(KeyCard(uid), goodsId)
	return cart, nil
}
