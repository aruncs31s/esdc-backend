package model

import "time"

type Project struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"unique"`
	Image       *string
	Description string
	GithubLink  string
	LiveUrl     *string
	CreatedBy   int `gorm:"not null"`
	ModifiedBy  *int
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
	Likes       int        `gorm:"default:0"`
	Category    string     `gorm:"default:'General'"`
	Cost        int        `gorm:"default:0"`

	Contributors []User          `gorm:"many2many:project_contributors;" json:"contributors"`
	Tags         *[]Tag          `gorm:"many2many:project_tags;" json:"tags"`
	Technologies *[]Technologies `gorm:"many2many:project_technologies;" json:"technologies"`
}

func (Project) TableName() string {
	return "projects"
}
