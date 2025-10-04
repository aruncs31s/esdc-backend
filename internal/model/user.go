package model

// TODO: Seperate the concerns

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Username    string        `gorm:"unique;not null" json:"username"`
	Email       string        `gorm:"unique;not null" json:"email"`
	Password    string        `gorm:"not null" json:"password"`
	Role        string        `gorm:"not null;default:'user'" json:"role"`
	Verified    *bool         `gorm:"not null;default:false" json:"verified"`
	Status      string        `gorm:"not null;default:'active'" json:"status"`
	CreatedAt   int64         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64         `gorm:"autoUpdateTime" json:"updated_at"`
	Github      *Github       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Submissions *[]Submission `gorm:"foreignKey:UserID;"`
	Details     *UserDetails  `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`
}

type Github struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   uint   `gorm:"not null"`
	Username string `gorm:"not null"`
}

func (Github) TableName() string {
	return "github"
}

func (User) TableName() string {
	return "users"
}
