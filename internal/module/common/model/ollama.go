package model

// This is used to get the already asked questions so that , we can avoid asking the same question again to ollama

type Ollama struct {
	ID        int     `gorm:"primaryKey"`
	ModelName string  `gorm:"column:model_name"`
	Prompt    string  `gorm:"column:prompt"`
	Response  string  `gorm:"column:response"`
	AskedBy   *uint `gorm:"column:asked_by"`
}

func (Ollama) TableName() string {
	return "ollama"
}
