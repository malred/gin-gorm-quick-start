package controllers

import (
	"quick-start/initializers"
	"quick-start/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// 添加
func PostsCreate(c *gin.Context) {
	// 获取json数据
	var body struct {
		Body  string `json:"body" validate:"min=6,max=255"`
		Title string `json:"title" validate:"min=3,max=20"`
	}
	// 将json数据绑定到结构体
	c.Bind(&body)

	// 校验
	err := validate.Struct(body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	// 创建
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// 查询所有
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// 分页查询
func PostsPage(c *gin.Context) {
	limitStr := c.Query("limit")
	curPageStr := c.Query("currentPage")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.Status(500)
		return
	}
	curPage, err := strconv.Atoi(curPageStr)
	if err != nil {
		c.Status(500)
		return
	}

	var posts []models.Post
	initializers.DB.
		Scopes(models.Paginate(curPage, limit)).
		Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// 根据id查询
func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

// 修改
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string `json:"body" validate:"min=6,max=255"`
		Title string `json:"title" validate:"min=3,max=20"`
	}
	// 将json数据绑定到结构体
	c.Bind(&body)

	
	// 校验
	err := validate.Struct(body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

// 删除
func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
