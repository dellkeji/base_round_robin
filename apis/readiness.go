package apis

import (
	"round_robin_with_weight/utils"

	"github.com/gin-gonic/gin"
)

// ReadinessHealth : return container is ok
func ReadinessHealth(c *gin.Context) {
	utils.CommonResponse(c, 0, "success", gin.H{})
	return
}
