package handler

import (
	"esdc-backend/internal/module/auth/dto"
	"esdc-backend/internal/module/auth/service"
	"esdc-backend/utils"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}
type authHandler struct {
	authService    service.AuthService
	responseHelper responsehelper.ResponseHelper
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &authHandler{
		authService:    authService,
		responseHelper: responseHelper,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginData body dto.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /login [post]
func (h *authHandler) Login(c *gin.Context) {
	loginData, failed := utils.GetJSONData[dto.LoginRequest](c, h.responseHelper, utils.ErrBadRequest.Error(), utils.ErrDetailBadRequestJSONPayload.Error())
	if failed {
		return
	}
	token, failed := getToken(h, loginData.Email, loginData.Password, c)
	if failed {
		return
	}
	data := map[string]string{"token": token}

	h.responseHelper.Success(c, data)
}

// Register godoc
// @Summary User registration
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param registerData body dto.RegisterRequest true "Registration data"
// @Success 200 {object} map[string]interface{} "Registration successful"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /register [post]
func (h *authHandler) Register(c *gin.Context) {
	registerData, failed := utils.GetJSONData[dto.RegisterRequest](c, h.responseHelper, "Bad request", "Invalid request payload")
	if failed {
		return
	}
	if failed := register(h, registerData, c); failed {
		return
	}
	token, failed := getToken(h, registerData.Email, registerData.Password, c)
	if failed {
		return
	}
	data := map[string]string{"token": token}
	h.responseHelper.Success(c, data)
}

func getToken(h *authHandler, email string, password string, c *gin.Context) (string, bool) {
	token, err := h.authService.Login(email, password)
	if err != nil {
		h.responseHelper.InternalError(c, "Could not login after registration", err)
		return "", true
	}
	return token, false
}

func register(h *authHandler, registerData dto.RegisterRequest, c *gin.Context) bool {
	err := h.authService.Register(registerData)
	if err != nil {
		h.responseHelper.InternalError(c, "Could not register user", err)
		return true
	}
	return false
}
