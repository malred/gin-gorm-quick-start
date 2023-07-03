package middleware

import (
	"fmt"
	"net/http"
	"os"
	"quick-start/initializers"
	"quick-start/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

//解决跨域问题
func RequireAuthHeader(c *gin.Context) {
	// 获取token
	tokenStr := c.Request.Header.Get("Authorization")
	fmt.Println(tokenStr[:7])
	if tokenStr[:6] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// 验证token
	token, err := jwt.Parse(tokenStr[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		fmt.Println(token)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if cliams, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 过期
		if float64(time.Now().Unix()) > cliams["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// 根据jwt携带消息(userId)查询数据库
		var user models.User
		initializers.DB.Table("user").First(&user, cliams["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// c.Set("user", user)
		c.Set("id", user.ID)
		c.Set("username", user.Username)
		c.Set("email", user.Email)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
