package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", )
	r.Run("0.0.0.0:9002")
}
