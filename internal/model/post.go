package model

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID uint   `gorm:"not null"`
}

func (Post) TableName() string {
	return "posts"
}
