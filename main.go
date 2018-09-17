package main

import (
	"fmt"
	"os"
	"round_robin_with_weight/storage"

	"round_robin_with_weight/apis"
	"round_robin_with_weight/config"

	"github.com/gin-gonic/gin"
)

// Init :
func init() {
	Execute()
}

func start() error {
	configInfo := config.GlobalConfigurations
	if !configInfo.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// redis client
	redisSession := storage.GetDefaultRedisSession()
	defer redisSession.Close()

	router := gin.Default()
	router.GET("ping", apis.GinPing)
	router.GET("domain", apis.GetAPIDomain)
	router.GET("readiness_health", apis.ReadinessHealth)
	router.GET("liveness_health", apis.LivenessHealth)
	return router.Run(
		fmt.Sprintf("%v:%v", configInfo.Server.Host, configInfo.Server.Port),
	)
}

func main() {
	err := start()
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
}
