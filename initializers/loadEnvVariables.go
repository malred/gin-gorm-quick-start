package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// 加载.env,里面设置了port,gin会读取并设置端口号
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
