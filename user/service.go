package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUser) (User, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) RegisterUser(input RegisterUser) (User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	user := User{
		Name:       input.Name,
		Email:      input.Email,
		Occupation: input.Occupation,
		Password:   string(password),
		Role:       "user",
	}
	user, err = s.repository.Save(user)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	return user, nil
}
