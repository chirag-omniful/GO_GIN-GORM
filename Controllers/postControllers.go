package Controllers

import (
	"Learn/initializers"
	"Learn/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSinglePost(ctx *gin.Context) {

	var user models.Post
	err := ctx.Bind(&user) // should bind json bhi use kr skte hai
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result := initializers.DB.Create(&user)
	//test := result.RowsAffected
	//fmt.Println(test)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"post": user,
	})
}

func CreatePostSelectedFields(ctx *gin.Context) {
	var user models.Post
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result := initializers.DB.Select("Name", "Age").Create(&user)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, user)
}

func CreateInBatches(ctx *gin.Context) {

	var users []*models.Post
	err := ctx.Bind(&users)
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result := initializers.DB.CreateInBatches(users, 2)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, users)
}

func CreateWithMap(ctx *gin.Context) {
	var user models.Post
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result := initializers.DB.Model(models.Post{}).Create(map[string]interface{}{
		"name":        user.Name,
		"age":         user.Age,
		"description": user.Description,
	})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(200, user)
}

func CreateMultiplePost(ctx *gin.Context) {

	var users []*models.Post
	err := ctx.Bind(&users)
	if err != nil {
		fmt.Println("issue in binding the body which comes with the request")
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	result := initializers.DB.Create(users)
	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"post": users,
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
		ctx.JSON(http.StatusBadRequest, err.Error())
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
