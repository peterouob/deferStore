package main

import (
	"github.com/gin-gonic/gin"
	"server/server/account/account"
	serviceOrder "server/server/account/order"
	"server/server/service"
)

func main() {
	r := gin.New()

	r.POST("/login", service.Login)
	r.POST("/register", service.Register)

	account := r.Group("/account")
	{
		account.POST("/profile", serviceAccount.Profile)
		order := account.Group("/order")
		{
			order.POST("/list", serviceOrder.List)
		}
	}

	r.Run(":8081")
}
