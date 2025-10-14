package model

type Teams struct {
	ID        int64     `gorm:"primaryKey;"`
	Name      string    `gorm:"column:name;uniqueIndex"`
	Image     *string   `gorm:"column:image;"`
	GithubUrl *string   `gorm:"column:github_url;"`
	Projects  []Project `gorm:"many2many:project_teams;" json:"projects"`
}

func (Teams) TableName() string {
	return "teams"
}
