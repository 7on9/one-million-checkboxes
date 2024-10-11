package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// r.GET("/hello/:name", getWithParams)
	// r.GET("/enqueue", getEnqueue)
	// r.GET("/dequeue", getDequeue)
	// r.GET("/user/:userCode", getGetUserPosition)
	// r.GET("/queue/options", getQueueOptions)
	// r.POST("/queue/options", postUpdateQueueOptions)
	// services.SetupScanInterval()
	return r
}
