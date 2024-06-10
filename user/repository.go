package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(Id int) (User, error)
	Update(user User) (User, error)
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

func (repository *repositoryImpl) FindByEmail(email string) (User, error) {
	user := new(User)
	if err := repository.db.Where("email = ?", email).Find(user).Error; err != nil {
		return *user, nil
	}
	return *user, nil
}

func (repository *repositoryImpl) FindById(Id int) (User, error) {
	user := new(User)
	if err := repository.db.Where("Id = ?", Id).Find(user).Error; err != nil {
		return *user, nil
	}
	return *user, nil
}

func (repository *repositoryImpl) Update(user User) (User, error) {
	if err := repository.db.Save(&user).Error; err != nil {
		return user, nil
	}
	return user, nil
}
