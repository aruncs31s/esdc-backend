package model

import "time"

type Submission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"column:user_id"`
	ProblemID   uint      `gorm:"column:problem_id"`
	Attempts    int       `gorm:"column:attempts"`
	SubmittedAt time.Time `gorm:"column:submitted_at"`
	Status      bool      `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Submission) TableName() string {
	return "submissions"
}
