package main

import (
	"aid.dev/etherms/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", internal.Login)
	r.Run("0.0.0.0:9002")
}
