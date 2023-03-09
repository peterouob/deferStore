package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"server/component/config"
	"server/logic/orm/dal"
	serviceOrder "server/server/service/account/order"
	"server/server/service/h"
	"server/server/service/user"
)

func main() {
	r := gin.New()

	db, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")))
	if err != nil {
		panic(fmt.Errorf("%s", "connect to database have error"))
	}
	dal.SetDefault(db)

	r.Use(Cors)
	r.POST("/login", user.Login)
	r.POST("/register", user.Register)

	account := r.Group("/account")
	account.Use(h.Auth())
	account.POST("/list", serviceOrder.List)
	//order.POST("/list", serviceOrder.GetList)

	r.Run(":8081")

}

func Cors(context *gin.Context) {
	method := context.Request.Method
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	fmt.Println(context.GetHeader("Origin"))
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	context.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
