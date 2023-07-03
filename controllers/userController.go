package controllers

import (
	"net/http"
	"os"
	"quick-start/initializers"
	"quick-start/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func Register(c *gin.Context) {
	var body struct {
		Username        string `json:"username" validate:"min=5"`
		Password        string `json:"password" validate:"min=8"`
		RetypedPassword string `json:"retypedPassword" validate:"min=8,eqfield=Password"`
		Email           string `json:"email" validate:"email"`
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

	var user models.User
	initializers.DB.
		Table("user").
		Where("email = ?", body.Email).
		First(&user)

	// 已存在(其实username也是唯一的,设计在表里了)
	if user != (models.User{}) {
		c.JSON(500, gin.H{
			"error": "the email is already exist!",
		})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user = models.User{
		Username: body.Username,
		Password: string(hash),
		Email:    body.Email,
	}

	result := initializers.DB.
		Table("user").
		Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

// 登录
func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to bind the body!",
		})
		return
	}

	var user models.User
	initializers.DB.
		Table("user").
		Where("username = ?", body.Username).
		Find(&user)

	if user == (models.User{}) {
		c.JSON(500, gin.H{
			"error": "con't find the user!",
		})
		return
	}

	// 校验密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	// 生成jwt令牌(传输协议用https,加密传输,防止jwt泄露)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		// 过期时间
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// 传入密钥,加密
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	// 设置cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenStr, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{
		"token":  tokenStr,
		"userId": user.ID,
	})
}

// 校验,返回用户信息
func Validate(c *gin.Context) {
	id, _ := c.Get("id")
	uname, _ := c.Get("username")
	email, _ := c.Get("email")
	c.JSON(200, gin.H{
		"id":       id,
		"username": uname,
		"email":    email,
	})
}
