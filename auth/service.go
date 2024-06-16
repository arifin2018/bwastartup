package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateKey(UserId int) (string, error)
}

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP")

func (s *JwtService) GenerateKey(UserId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("1")
		log.Println(err.Error())
		return "", err
	}
	return signedToken, err

	// claim := jwt.MapClaims{}
	// claim["user_id"] = UserId

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// signedToken, err := token.SignedString(SECRET_KEY)
	// if err != nil {
	// 	return signedToken, err
	// }

	// return signedToken, nil
}
