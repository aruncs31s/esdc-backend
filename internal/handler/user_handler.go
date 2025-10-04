package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(c *gin.Context)
	LogOut(c *gin.Context)
	Register(c *gin.Context)
	VerifyEmail(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}
type userHandler struct {
	responseHelper responses.ResponseHelper
	userService    service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	responseHelper := responses.NewResponseHelper()
	return &userHandler{
		responseHelper: responseHelper,
		userService:    userService,
	}
}

func (h *userHandler) Login(c *gin.Context) {
	var loginData dto.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		h.responseHelper.BadRequest(c, "Bad request", "Invalid request payload")
		return
	}

	token, err := h.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		h.responseHelper.InternalError(c, "Could not create token", err)
		return
	}

	h.responseHelper.Success(c, gin.H{"token": token})
}

func (h *userHandler) LogOut(c *gin.Context) {
	// Implementation of logout logic
}

func (h *userHandler) Register(c *gin.Context) {
	var registerData dto.RegisterRequest
	// Get the register Data
	if err := c.ShouldBindJSON(&registerData); err != nil {
		h.responseHelper.BadRequest(c, "Bad request", "Invalid request payload")
		return
	}

	err := h.userService.Register(
		registerData.Username,
		registerData.Email,
		registerData.Password)
	if err != nil {
		h.responseHelper.InternalError(c, "Could not register user", err)
		return
	}
	h.responseHelper.Success(c, nil)
}

func (h *userHandler) VerifyEmail(c *gin.Context) {
	// Implementation of email verification logic
}

func (h *userHandler) ForgotPassword(c *gin.Context) {
	// Implementation of forgot password logic
}

func (h *userHandler) ResetPassword(c *gin.Context) {
	// Implementation of reset password logic
}
