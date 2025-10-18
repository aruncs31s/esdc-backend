package handler

import (
	"esdc-backend/internal/module/common/dto"
	"esdc-backend/internal/module/common/service"
	"strconv"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type NotificationHandler interface {
	SendNotification(c *gin.Context)
	GetSingleNotification(c *gin.Context)
	GetUserNotifications(c *gin.Context)
	MarkAsRead(c *gin.Context)
}

type notificationHandler struct {
	notificationService service.NotificationService
	responseHelper      responsehelper.ResponseHelper
}

func NewNotificationHandler(notificationService service.NotificationService) NotificationHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &notificationHandler{
		notificationService: notificationService,
		responseHelper:      responseHelper,
	}
}

func (h *notificationHandler) SendNotification(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		h.responseHelper.Unauthorized(c, "Unauthorized")
		return
	}
	var data dto.NotificationRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		h.responseHelper.BadRequest(c, err.Error(), "Something Bad")
		return
	}

	err := h.notificationService.SendNotification(data)

	if err != nil {
		h.responseHelper.InternalError(c, "Something wrong", err)
		return
	}
	h.responseHelper.Success(c, "Notification sent successfully")
}
func (h *notificationHandler) MarkAsRead(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		h.responseHelper.Unauthorized(c, "Unauthorized")
		return
	}
	idStr := c.Param("id")
	notificationID, err := strconv.Atoi(idStr)
	if err != nil {
		h.responseHelper.BadRequest(c, "Invalid notification ID", err.Error())
		return
	}

	err = h.notificationService.MarkAsRead(user, uint(notificationID))
	if err != nil {
		h.responseHelper.InternalError(c, "Something wrong", err)
		return
	}
	h.responseHelper.Success(c, "Notification marked as read successfully")
}

// Currently for users only
func (h *notificationHandler) GetUserNotifications(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		h.responseHelper.Unauthorized(c, "Unauthorized")
		return
	}
	notifications, err := h.notificationService.GetUserNotifications(user)
	if err != nil {
		h.responseHelper.InternalError(c, "Something went wrong", err)
		return
	}
	if notifications == nil {
		h.responseHelper.NotFound(c, "No notifications found")
		return
	}
	h.responseHelper.Success(c, notifications)
}
func (h *notificationHandler) GetSingleNotification(c *gin.Context) {
	idStr := c.Param("id")
	notificationID, err := strconv.Atoi(idStr)
	if err != nil {
		h.responseHelper.BadRequest(c, "Invalid notification ID", err.Error())
		return
	}
	notification, err := h.notificationService.GetSingleNotification(uint(notificationID))
	if err != nil {
		h.responseHelper.InternalError(c, "Something went wrong", err)
		return
	}
	if notification == nil {
		h.responseHelper.NotFound(c, "Notification not found")
		return
	}
	h.responseHelper.Success(c, notification)
}
