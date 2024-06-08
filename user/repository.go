package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (repository *repositoryImpl) Save(user User) (User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
