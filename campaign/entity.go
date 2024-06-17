package campaign

import "time"

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImage    []CampaignImage
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (CampaignImage) TableName() string {
	return "images"
}
