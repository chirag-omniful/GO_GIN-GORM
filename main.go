package main

import (
	"Learn/Controllers"
	"Learn/initializers"
	"Learn/migrate"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDB()
	migrate.MigrateDB()
}

func main() {
	fmt.Println("yo")

	r := gin.Default()

	r.POST("/post", Controllers.CreateSinglePost)
	r.POST("/createSelect", Controllers.CreatePostSelectedFields)
	r.POST("/createInBatches", Controllers.CreateInbatches)
	r.POST("posts", Controllers.CreateMultiplePost)
	r.GET("/posts", Controllers.GetAllPosts)
	r.GET("/posts/:id", Controllers.GetSinglePost)
	r.PUT("post/:id", Controllers.UpdatePost)
	r.DELETE("post/:id", Controllers.DeletePost)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error in loading server")
	}

}
