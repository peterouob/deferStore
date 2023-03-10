package main

import (
	"github.com/gin-gonic/gin"
	"server/component/mysql"
	"server/server/service/router"
)

func main() {
	r := gin.New()
	go mysql.InitDB()
	router.Router(r)
	r.Run(":8081")
}
