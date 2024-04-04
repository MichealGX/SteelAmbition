package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(c *gin.Context, Db *sql.DB, userDTO UserDTO) {
	// 查询用户
	var user UserDTO
	var err error
	err = Db.QueryRow("SELECT UserID, Username, Password FROM User WHERE Username = ? AND Password = ?", userDTO.UserName, userDTO.Password).Scan(&user.UserID, &user.UserName, &user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Unauthorized"})
		return
	}

	// 生成 JWT
	var token string
	token, err = SignedToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to sign token"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Login success", "token": token})
}
