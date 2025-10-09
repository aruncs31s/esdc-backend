package model

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Username  string `gorm:"unique;not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null;default:user" json:"role"`
	Verified  *bool  `gorm:"not null;default:false" json:"verified"`
	Status    string `gorm:"not null;default:active" json:"status"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`

	Github      *Github       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Submissions *[]Submission `gorm:"foreignKey:UserID"`
	Details     *UserDetails  `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`

	// fixed here:
	Projects []Project `gorm:"many2many:project_contributors;" json:"projects"`
}

type Github struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	UserID   uint   `gorm:"column:user_id"`
	Username string `gorm:"column:username;unique;not null"`
}

func (Github) TableName() string {
	return "github"
}

func (User) TableName() string {
	return "users"
}
