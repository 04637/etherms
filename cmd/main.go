package main

import (
	"aid.dev/etherms/config"
	"aid.dev/etherms/internal"
	"aid.dev/etherms/internal/etherpad"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.Load()
	etherpad.InitPad(conf)
	r := gin.Default()
	r.GET("/ping", internal.Login)
	_ = r.Run(":9002")
}
