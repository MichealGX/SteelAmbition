package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context, Db *sql.DB, userDTO UserDTO) {
	// 检查用户名是否已存在
	var count int
	row := Db.QueryRow("SELECT COUNT(*) FROM User WHERE Username = ?", userDTO.UserName)
	if err := row.Scan(&count); err != nil {
		log.Println("Failed to query database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to query database"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Username already exists"})
		return
	}

	// 插入新用户到数据库
	_, err := Db.Exec("INSERT INTO User (Username, Password, Email) VALUES (?, ?, ?)", userDTO.UserName, userDTO.Password, userDTO.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "Failed to insert user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": 0, "msg": "success"})
}
