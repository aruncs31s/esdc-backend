package handler

import (
	"esdc-backend/internal/module/common/service"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type FileHandler interface {
	UploadImage(c *gin.Context)
	UploadFile(c *gin.Context)
	UploadMultipleFiles(c *gin.Context)
}

type fileHandler struct {
	responseHelper responsehelper.ResponseHelper
	fileService    service.FileService
}

func NewFileHandler(fileService service.FileService) FileHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &fileHandler{
		responseHelper: responseHelper,
		fileService:    fileService,
	}
}

// sanitizePath removes path traversal attempts and ensures safe path
func sanitizePath(path string) string {
	// Clean the path to remove any .. or . elements
	cleaned := filepath.Clean(path)
	// Remove any leading slashes or dots
	cleaned = strings.TrimPrefix(cleaned, "/")
	cleaned = strings.TrimPrefix(cleaned, "./")
	// Replace any remaining .. with empty string
	cleaned = strings.ReplaceAll(cleaned, "..", "")
	return cleaned
}

// validateUploadDir ensures the upload directory is safe
func validateUploadDir(dir string) string {
	if dir == "" {
		return "files"
	}
	// Sanitize the directory path
	sanitized := sanitizePath(dir)
	// Only allow alphanumeric characters, hyphens, and underscores
	allowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	result := ""
	for _, char := range sanitized {
		if strings.ContainsRune(allowed, char) {
			result += string(char)
		}
	}
	if result == "" {
		return "files"
	}
	return result
}

// UploadImage godoc
// @Summary Upload an image
// @Description Upload a single image file (jpg, jpeg, png, gif, webp) with max size 5MB
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param image formData file true "Image file to upload"
// @Success 200 {object} map[string]interface{} "Image uploaded successfully"
// @Failure 400 {object} map[string]interface{} "Invalid file or no file uploaded"
// @Failure 500 {object} map[string]interface{} "Failed to upload image"
// @Router /upload/image [post]
func (h *fileHandler) UploadImage(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		h.responseHelper.BadRequest(c, "No file uploaded", err.Error())
		return
	}

	// Validate file (only images, max 5MB)
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	maxSize := int64(5 * 1024 * 1024) // 5MB
	if err := h.fileService.ValidateFile(file, allowedExtensions, maxSize); err != nil {
		h.responseHelper.BadRequest(c, "Invalid file", err.Error())
		return
	}

	// Upload file to images directory
	filePath, err := h.fileService.UploadFile(c, file, "images")
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to upload image", err)
		return
	}

	// Sanitize the file path for URL construction
	safeFilePath := sanitizePath(filePath)

	h.responseHelper.Success(c, gin.H{
		"message": "Image uploaded successfully",
		"path":    filePath,
		"url":     fmt.Sprintf("/uploads/%s", safeFilePath),
	})
}

// UploadFile godoc
// @Summary Upload a file
// @Description Upload a single file of any type with max size 10MB
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param file formData file true "File to upload"
// @Param dir query string false "Upload directory" default(files)
// @Success 200 {object} map[string]interface{} "File uploaded successfully"
// @Failure 400 {object} map[string]interface{} "Invalid file or no file uploaded"
// @Failure 500 {object} map[string]interface{} "Failed to upload file"
// @Router /upload/file [post]
func (h *fileHandler) UploadFile(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		h.responseHelper.BadRequest(c, "No file uploaded", err.Error())
		return
	}

	// Get and validate directory parameter
	uploadDir := validateUploadDir(c.Query("dir"))

	// Validate file (max 10MB)
	maxSize := int64(10 * 1024 * 1024) // 10MB
	if err := h.fileService.ValidateFile(file, nil, maxSize); err != nil {
		h.responseHelper.BadRequest(c, "Invalid file", err.Error())
		return
	}

	// Upload file
	filePath, err := h.fileService.UploadFile(c, file, uploadDir)
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to upload file", err)
		return
	}

	// Sanitize the file path for URL construction
	safeFilePath := sanitizePath(filePath)

	h.responseHelper.Success(c, gin.H{
		"message":  "File uploaded successfully",
		"path":     filePath,
		"url":      fmt.Sprintf("/uploads/%s", safeFilePath),
		"filename": file.Filename,
		"size":     service.FormatFileSize(file.Size),
	})
}

// UploadMultipleFiles godoc
// @Summary Upload multiple files
// @Description Upload multiple files of any type with max size 10MB each
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param files formData file true "Files to upload" multiple
// @Param dir query string false "Upload directory" default(files)
// @Success 200 {object} map[string]interface{} "Files upload completed"
// @Failure 400 {object} map[string]interface{} "No files uploaded or failed to parse form"
// @Router /upload/files [post]
func (h *fileHandler) UploadMultipleFiles(c *gin.Context) {
	// Get form
	form, err := c.MultipartForm()
	if err != nil {
		h.responseHelper.BadRequest(c, "Failed to parse form", err.Error())
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		h.responseHelper.BadRequest(c, "No files uploaded", "")
		return
	}

	// Get and validate directory parameter
	uploadDir := validateUploadDir(c.Query("dir"))

	// Upload all files
	var uploadedFiles []gin.H
	var failedFiles []gin.H

	maxSize := int64(10 * 1024 * 1024) // 10MB per file

	for _, file := range files {
		// Validate each file
		if err := h.fileService.ValidateFile(file, nil, maxSize); err != nil {
			failedFiles = append(failedFiles, gin.H{
				"filename": file.Filename,
				"error":    err.Error(),
			})
			continue
		}

		// Upload file
		filePath, err := h.fileService.UploadFile(c, file, uploadDir)
		if err != nil {
			failedFiles = append(failedFiles, gin.H{
				"filename": file.Filename,
				"error":    err.Error(),
			})
			continue
		}

		// Sanitize the file path for URL construction
		safeFilePath := sanitizePath(filePath)

		uploadedFiles = append(uploadedFiles, gin.H{
			"filename": file.Filename,
			"path":     filePath,
			"url":      fmt.Sprintf("/uploads/%s", safeFilePath),
			"size":     service.FormatFileSize(file.Size),
		})
	}

	h.responseHelper.Success(c, gin.H{
		"message":        fmt.Sprintf("Uploaded %d of %d files", len(uploadedFiles), len(files)),
		"uploaded":       uploadedFiles,
		"failed":         failedFiles,
		"uploaded_count": len(uploadedFiles),
		"failed_count":   len(failedFiles),
	})
}
