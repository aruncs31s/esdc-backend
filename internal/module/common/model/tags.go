package model

type Tag struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (Tag) TableName() string {
	return "tags"
}
