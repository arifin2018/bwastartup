package user

import (
	"time"
)

type User struct {
	Id             uint   `json:"-"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	AvatarFileName string `json:"avatar_file_name"`
	Token          string `json:"token"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RegisterUser struct {
	Id         uint   `json:"-"`
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
