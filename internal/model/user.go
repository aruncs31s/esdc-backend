package model

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"column:name;not null" json:"name"`
	Username    string        `gorm:"column:username;unique;not null" json:"username"`
	Email       string        `gorm:"column:email;unique;not null" json:"email"`
	Password    string        `gorm:"column:password;not null" json:"password"`
	Role        string        `gorm:"column:role;not null;default:user" json:"role"`
	Verified    *bool         `gorm:"column:verified;not null;default:false" json:"verified"`
	Status      string        `gorm:"column:status;not null;default:active" json:"status"`
	CreatedAt   int64         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   int64         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Teams       *[]Teams      `gorm:"many2many:team_members;" json:"teams"`
	Github      *Github       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Submissions *[]Submission `gorm:"foreignKey:UserID"`
	Details     *UserDetails  `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`

	// fixed here:
	Projects      *[]Project     `gorm:"many2many:project_contributors;" json:"projects"`
	LikedProjects *[]Project     `gorm:"many2many:project_likes;" json:"liked_projects"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications"`
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
