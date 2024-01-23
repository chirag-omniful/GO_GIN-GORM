package Controllers

import (
	"Learn/initializers"
	"Learn/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var user models.Post

	err := ctx.Bind(&user) // should bind json bhi use kr skte hai
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
	}

	post := models.Post{Name: user.Name, Age: user.Age, Description: user.Description}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(ctx *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.Post
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("issue in binding body (update post function)")
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Name:        user.Name,
		Age:         user.Age,
		Description: user.Description,
	})

	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	initializers.DB.Delete(&post, id)

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
