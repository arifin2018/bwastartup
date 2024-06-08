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
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
