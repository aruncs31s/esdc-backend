package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AuthRepositoryReader
	AuthRepositoryWriter
}

type AuthRepositoryReader interface {
	FindByEmail(email string) (model.User, error)
}
type AuthRepositoryWriter interface {
	CreateUser(user *model.User) error
}
type authRepository struct {
	reader AuthRepositoryReader
	writer AuthRepositoryWriter
}

type authRepositoryReader struct {
	db *gorm.DB
}
type authRepositoryWriter struct {
	db *gorm.DB
}

func newAuthRepositoryReader(db *gorm.DB) AuthRepositoryReader {
	return &authRepositoryReader{db: db}
}
func newAuthRepositoryWriter(db *gorm.DB) AuthRepositoryWriter {
	return &authRepositoryWriter{db: db}
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	reader := newAuthRepositoryReader(db)
	writer := newAuthRepositoryWriter(db)
	return &authRepository{
		reader: reader,
		writer: writer,
	}
}
func (r *authRepository) FindByEmail(email string) (model.User, error) {
	return r.reader.FindByEmail(email)
}
func (r *authRepository) CreateUser(user *model.User) error {
	return r.writer.CreateUser(user)
}

func (r *authRepositoryReader) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}
func (r *authRepositoryWriter) CreateUser(user *model.User) error {
	result := r.db.Create(&user)
	return result.Error
}
