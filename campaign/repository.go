package campaign

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
}

type repositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{
		DB: db,
	}
}

func (r *repositoryImpl) FindAll() ([]Campaign, error) {
	var campaigns = new([]Campaign)
	if err := r.DB.Preload("CampaignImage", "is_primary = ?", 1).Find(campaigns).Error; err != nil {
		return *campaigns, nil
	}
	return *campaigns, nil
}

func (r *repositoryImpl) FindByUserId(userId int) ([]Campaign, error) {
	fmt.Println("arifin")
	var campaigns = new([]Campaign)
	if err := r.DB.Where("user_id = ?", userId).Preload("CampaignImage", "is_primary = ?", 1).Find(campaigns).Error; err != nil {
		return *campaigns, nil
	}
	return *campaigns, nil
}
