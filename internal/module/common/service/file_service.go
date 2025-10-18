package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	UploadFile(c *gin.Context, file *multipart.FileHeader, uploadDir string) (string, error)
	DeleteFile(filePath string) error
	ValidateFile(file *multipart.FileHeader, allowedExtensions []string, maxSize int64) error
}

type fileService struct {
	baseUploadDir string
}

func NewFileService(baseUploadDir string) FileService {
	return &fileService{
		baseUploadDir: baseUploadDir,
	}
}

// UploadFile handles file upload and returns the file path
func (s *fileService) UploadFile(c *gin.Context, file *multipart.FileHeader, uploadDir string) (string, error) {
	// Create upload directory if it doesn't exist
	fullPath := filepath.Join(s.baseUploadDir, uploadDir)
	if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	randomID := generateRandomID(8)
	filename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), randomID, ext)
	filePath := filepath.Join(fullPath, filename)

	// Save the file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return relative path for storage in database
	relativePath := filepath.Join(uploadDir, filename)
	return relativePath, nil
}

// DeleteFile deletes a file from the filesystem
func (s *fileService) DeleteFile(filePath string) error {
	fullPath := filepath.Join(s.baseUploadDir, filePath)
	if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}

// ValidateFile validates file extension and size
func (s *fileService) ValidateFile(file *multipart.FileHeader, allowedExtensions []string, maxSize int64) error {
	// Check file size
	if file.Size > maxSize {
		return fmt.Errorf("file size exceeds maximum allowed size of %d bytes", maxSize)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if len(allowedExtensions) > 0 {
		allowed := false
		for _, allowedExt := range allowedExtensions {
			if ext == strings.ToLower(allowedExt) {
				allowed = true
				break
			}
		}
		if !allowed {
			return fmt.Errorf("file extension %s is not allowed", ext)
		}
	}

	return nil
}

// generateRandomID generates a random ID string of specified length
func generateRandomID(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		// Fallback to timestamp-based ID
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(bytes)
}

// Helper function to check if file is an image
func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg"}
	for _, imgExt := range imageExtensions {
		if ext == imgExt {
			return true
		}
	}
	return false
}

// Helper function to get file size in human-readable format
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
