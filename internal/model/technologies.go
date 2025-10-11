package model

type Technologies struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;unique;not null"`
}

func (Technologies) TableName() string {
	return "technologies"
}
