package models

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~The message in the comment of photo is required"`
	User    *User
	Photo   *Photo
}
