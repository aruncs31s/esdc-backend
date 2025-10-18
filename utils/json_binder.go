package utils

import (
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

func GetJSONData[T any](c *gin.Context, rh responsehelper.ResponseHelper, badRequestMsg, invalidPayloadMsg string) (T, bool) {
	var data T
	// Get the data
	if err := c.ShouldBindJSON(&data); err != nil {
		rh.BadRequest(c, badRequestMsg, invalidPayloadMsg)
		return data, true
	}
	return data, false
}
