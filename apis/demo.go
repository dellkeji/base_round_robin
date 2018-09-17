package apis

import (
	"round_robin_with_weight/utils"

	"github.com/gin-gonic/gin"
)

// GinPing :
func GinPing(c *gin.Context) {
	utils.CommonResponse(c, 200, "Pong", gin.H{})
}
