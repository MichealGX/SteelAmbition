package users

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func SignedToken(user UserDTO) (string, error) {
	// 创建声明（Claim）
	claims := jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 过期时间，这里设置为 24 小时后
	}

	// 创建 JWT 签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取字符串格式的 JWT
	signedToken, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		fmt.Println("Failed to sign token:", err)
		return "", err
	}

	// 输出生成的 JWT
	return signedToken, nil

}

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization请求头
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 校验Authorization格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// 解析JWT字符串
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名密钥
			return []byte("your_secret_key"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Failed to parse JWT token"})
			c.Abort()
			return
		}

		// 校验JWT有效性
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Invalid JWT token"})
			c.Abort()
			return
		}

		// 将解析后的JWT token信息存储到Context中
		c.Set("user", token.Claims)

		// 继续处理请求
		c.Next()
	}
}

// ExtractUser 提取用户信息
func ExtractUser(c *gin.Context) UserDTO {
	user := c.MustGet("user").(jwt.MapClaims)
	userDTO := UserDTO{
		UserID:   int(user["user_id"].(float64)),
		UserName: user["username"].(string),
	}
	return userDTO
}
