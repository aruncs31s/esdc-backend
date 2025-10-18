package model

type UserDetails struct {
	ID             uint    `gorm:"column:id;primaryKey"`
	Phone          *string `gorm:"column:phone"`
	ProfilePicLink string  `gorm:"column:profile_pic_link" json:"profile_pic_link"`
}

func (UserDetails) TableName() string {
	return "user_details"
}
