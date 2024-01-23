package main

import (
	"Learn/initializers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	fmt.Println("yo")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "code testing",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error in loading server")
	}
}
