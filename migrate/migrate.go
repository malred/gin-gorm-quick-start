package main

import (
	"quick-start/initializers"
	"quick-start/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	// 创建数据库
	initializers.DB.AutoMigrate(&models.Post{})
}
