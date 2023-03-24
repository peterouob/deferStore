package main

import (
	"github.com/gin-gonic/gin"
	"server/component/mysql"
	"server/component/redis"
	"server/server/service/router"
)

func main() {
	r := gin.New()
	go mysql.InitDB()
	go redis.RedisInit()
	router.Router(r)
	r.Run(":8081")
}
