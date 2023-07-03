package main

import (
	"github.com/gin-gonic/gin"
	"quick-start/controllers"
	"quick-start/initializers"
	"quick-start/middleware"
)

// 启动时自动调用
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(middleware.Cors())
	{
		r.GET("posts", controllers.PostsIndex)
		r.GET("posts/:id", controllers.PostsShow)
		r.GET("posts/page", controllers.PostsPage)
		r.POST("posts", middleware.RequireAuthHeader, controllers.PostsCreate)
		r.PATCH("posts/:id", middleware.RequireAuthHeader, controllers.PostsUpdate)
		r.DELETE("posts/:id", middleware.RequireAuthHeader, controllers.PostsDelete)
	}
	{
		r.GET("auth/profile", middleware.RequireAuthHeader, controllers.Validate)
		r.POST("users", controllers.Register)
		r.POST("auth/login", controllers.Login)
	}
	r.Run()
}
