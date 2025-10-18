package model

import "time"

type Project struct {
	ID           int             `gorm:"primaryKey"`
	Title        string          `gorm:"column:title;unique"`
	Image        *string         `gorm:"column:image"`
	Description  string          `gorm:"column:description"`
	GithubLink   string          `gorm:"column:github_link"`
	LiveUrl      *string         `gorm:"column:live_url"`
	Likes        int             `gorm:"column:likes;default:0"`
	Category     string          `gorm:"column:category;default:'General'"`
	Cost         int             `gorm:"column:cost;default:0"`
	Status       string          `gorm:"column:status;default:'active'"` // active, inactive, archived
	Visibility   int             `gorm:"column:visibility;default:0"`    // 0: public, 1: private
	CreatedBy    uint            `gorm:"column:created_by;not null"`
	ModifiedBy   *uint           `gorm:"column:modified_by"`
	CreatedAt    *time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	Creator      User            `gorm:"foreignKey:CreatedBy;references:ID" json:"creator"`
	Contributors *[]User         `gorm:"many2many:project_contributors;" json:"contributors"`
	Tags         *[]Tag          `gorm:"many2many:project_tags;" json:"tags"`
	Technologies *[]Technologies `gorm:"many2many:project_technologies;" json:"technologies"`
	LikedBy      []User          `gorm:"many2many:project_likes;" json:"liked_by"`
}

func (Project) TableName() string {
	return "projects"
}
