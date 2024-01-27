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
	r.POST("/createInBatches", Controllers.CreateInBatches)
	r.POST("/createWithMap", Controllers.CreateWithMap)
	r.POST("posts", Controllers.CreateMultiplePost)

	r.GET("/posts", Controllers.GetAllPosts)
	r.GET("/posts/:id", Controllers.GetSinglePost)
	r.GET("/lastPost", Controllers.GetLastPost)
	r.GET("/whereCondition", Controllers.GetRowByWhereCondition)
	r.GET("/orCondition", Controllers.GetByOrCondition)

	r.PUT("/post/:id", Controllers.UpdatePost)
	r.PUT("/post", Controllers.UpsertPost)

	r.DELETE("post/:id", Controllers.DeletePost)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error in loading server")
	}

}
