package utils

import "github.com/gin-gonic/gin"

// CommonResponse : common response for domain api
func CommonResponse(c *gin.Context, code uint, message string, data interface{}) {
	status := 200
	response := gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
	c.JSON(status, response)
}
