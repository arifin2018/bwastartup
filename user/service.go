package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUser) (User, error)
	LoginUser(input LoginUser) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
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

func (s *serviceImpl) LoginUser(input LoginUser) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	if user.Id == 0 {
		return User{}, errors.New("users not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	return user, nil
}

func (s *serviceImpl) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		return true, nil
	}

	return false, nil
}

func (s *serviceImpl) SaveAvatar(Id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(Id)
	if err != nil {
		return user, nil
	}
	user.AvatarFileName = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return user, nil
	}
	return updatedUser, err
}
