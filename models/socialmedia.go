package models

type SocialMedia struct {
	GormModel
	UserID         uint
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~The name of social media is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~The social media URL is required"`
	User           *User
}
