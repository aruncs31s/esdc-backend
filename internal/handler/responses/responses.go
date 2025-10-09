package responses

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseHelper interface {
	BadRequest(c *gin.Context, message string, details string)
	NotFound(c *gin.Context, message string)
	Unauthorized(c *gin.Context, message string)
	InternalError(c *gin.Context, message string, err error)
	Success(c *gin.Context, data interface{})
	SuccessWithPagination(c *gin.Context, data interface{}, meta interface{})
	Created(c *gin.Context, data interface{})
	Deleted(c *gin.Context, message string)
	Conflict(c *gin.Context, message string, details string)
}

// Response helper - centralizes response logic
// The context is same in the case of all the responses , but there is no need to , group it in a struct
// only one response per request , so there is no reuse for context.
type responseHelper struct{}

func NewResponseHelper() ResponseHelper {
	return &responseHelper{}
}

func (r *responseHelper) BadRequest(c *gin.Context, message string, details string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"error": gin.H{
			"code":    400,
			"status":  "BAD_REQUEST",
			"message": message,
			"details": details,
		},
	})
}

func (r *responseHelper) NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": false,
		"error": gin.H{
			"code":    404,
			"status":  "NOT_FOUND",
			"message": message,
		},
	})
}

func (r *responseHelper) Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": gin.H{
			"code":    401,
			"status":  "UNAUTHORIZED",
			"message": message,
		},
	})
}

func (r *responseHelper) InternalError(c *gin.Context, message string, err error) {
	log.Printf("Internal error: %v", err)

	c.JSON(http.StatusInternalServerError, gin.H{
		"status": false,
		"error": gin.H{
			"code":    500,
			"status":  "INTERNAL_SERVER_ERROR",
			"message": message,
			"details": err.Error(), // sanitizing this in production
		},
		"data": nil,
	})
}

func (r *responseHelper) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   data,
		"meta":   time.Now().Format(time.RFC3339),
	})
}
func (r *responseHelper) SuccessWithPagination(c *gin.Context, data interface{}, meta interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"data":       data,
		"pagination": meta,
	})
}

func (r *responseHelper) Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   data,
		"meta":   time.Now().Format(time.RFC3339),
	})
}
func (r *responseHelper) Deleted(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, gin.H{
		"status":  true,
		"message": message,
		"meta":    time.Now().Format(time.RFC3339),
	})
}
func (r *responseHelper) Conflict(c *gin.Context, message string, details string) {
	c.JSON(http.StatusConflict, gin.H{
		"status": false,
		"error": gin.H{
			"code":    409,
			"status":  "CONFLICT",
			"message": message,
			"details": details,
		},
	})
}
