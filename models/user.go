package models

import (
	"final-golang-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          uint8         `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required, range(8|200)~The minimum age is 8 and max is 200 -children are not allowed to register :)"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}