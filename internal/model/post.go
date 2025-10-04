package model

type Post struct {
	ID       int   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID int   `gorm:"not null"`
}

func (Post) TableName() string {
	return "posts"
}
