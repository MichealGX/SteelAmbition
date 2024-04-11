package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ExtractID(c *gin.Context, inf string) (uint, error) {
	ID, err := strconv.Atoi(c.Param(inf))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "Invalid " + inf})
		return 0, err
	}
	uID := uint(ID)
	return uID, nil
}
