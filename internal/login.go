package internal

import (

	"github.com/gin-gonic/gin"
)

// Login : User login from OA
// 1. Get ticket from OA
// 2. Get code from OA
// 3. Get username from database
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	}) 
}