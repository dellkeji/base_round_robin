package apis

import (
	"round_robin_with_weight/utils"

	"github.com/gin-gonic/gin"
)

// LivenessHealth : return container is ok
func LivenessHealth(c *gin.Context) {
	utils.CommonResponse(c, 0, "success", gin.H{})
	return
}
