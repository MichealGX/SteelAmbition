package users

import "gorm.io/gorm"

type UserDTO struct {
	gorm.Model
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"psw" binding:"required,min=8,max=16"`
	Email    string `json:"email"`
}
