package users

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

func JwtVerification(jwtString string) {
	// 解析 JWT
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil // 指定密钥用于解析 JWT
	})
	if err != nil {
		fmt.Println("Failed to parse token:", err)
		return
	}

	// 验证 JWT 签名
	if !token.Valid {
		fmt.Println("Invalid token")
		return
	}

	// 输出 JWT 中的声明信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Failed to get claims")
		return
	}
	fmt.Println("User ID:", claims["user_id"])
	fmt.Println("Username:", claims["username"])

	// 检查过期时间
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		fmt.Println("Token has expired")
		return
	}

	// 其他自定义校验...
}
