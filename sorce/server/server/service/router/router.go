package router

import (
	"github.com/gin-gonic/gin"
	serviceCart "server/server/service/account/cart"
	goodsService "server/server/service/account/goods"
	serviceOrder "server/server/service/account/order"
	"server/server/service/h"
	"server/server/service/user"
)

func Router(g *gin.Engine) {
	g.Use(h.Cors)
	v1 := g.Group("/user/v1")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.POST("/logout", user.Logout)
	}

	account := g.Group("/account")
	account.Use(h.Auth())
	{
		order := account.Group("/order")
		order.POST("/list", serviceOrder.List)
	}
	{
		goods := account.Group("/goods")
		//展示首頁
		goods.GET("/homepage", goodsService.HomePage)
		//傳入id後顯示的畫面
		goods.GET("/get", goodsService.Get)
	}
	{
		cart := account.Group("/cart")
		cart.POST("/add", serviceCart.Add)
	}
	//goods := account.Group("/goods")
	//{
	//
	//}
}
