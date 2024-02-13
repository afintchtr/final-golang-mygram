package models

type Photo struct {
	GormModel
	UserID   uint
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~The title of the photo is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~The photo URL is required"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	User     *User
}
