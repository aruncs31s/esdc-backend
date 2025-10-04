package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Project struct {
	ID           int       `gorm:"primaryKey"`
	CreatedBy    int       `gorm:"column:created_by;not null"`
	Name         string    `gorm:"column:name;not null;unique"`
	Image        string    `gorm:"column:image;not null"`
	Description  string    `gorm:"column:description"`
	Link         string    `gorm:"column:link"`
	GithubLink   string    `gorm:"column:github_link"`
	Tags         JSONSlice `gorm:"column:tags;type:text"`
	Contributers JSONInts  `gorm:"column:contributers;type:text"`
}

// JSONSlice is a custom type for []string that can be stored as JSON in the database
type JSONSlice []string

// Value implements the driver.Valuer interface
func (j JSONSlice) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal(j)
	return string(b), err
}

// Scan implements the sql.Scanner interface
func (j *JSONSlice) Scan(value interface{}) error {
	if value == nil {
		*j = JSONSlice{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("cannot scan into JSONSlice")
	}

	return json.Unmarshal(bytes, j)
}

// JSONInts is a custom type for []int that can be stored as JSON in the database
type JSONInts []int

// Value implements the driver.Valuer interface
func (j JSONInts) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "[]", nil
	}
	b, err := json.Marshal(j)
	return string(b), err
}

// Scan implements the sql.Scanner interface
func (j *JSONInts) Scan(value interface{}) error {
	if value == nil {
		*j = JSONInts{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("cannot scan into JSONInts")
	}

	return json.Unmarshal(bytes, j)
}

func (Project) TableName() string {
	return "projects"
}
