package auth

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateKey(UserId int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
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
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
