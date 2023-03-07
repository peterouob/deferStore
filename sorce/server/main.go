package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/component/config"
	"server/logic/orm/dal"
	serviceOrder "server/server/service/account/order"
	"server/server/service/user"
)

func main() {
	r := gin.New()

	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")))
	if err != nil {
		panic(fmt.Errorf("%s", "connect to database have error"))
	}
	dal.SetDefault(db)

	r.POST("/login", user.Login)
	r.POST("/register", user.Register)

	account := r.Group("/account")
	//account.Use(h.Auth())
	order := account.Group("/order")
	order.POST("/list", serviceOrder.List)
	//order.POST("/list", serviceOrder.GetList)

	r.Run(":8081")

}
