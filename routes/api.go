package routes

import (
	"gin-api/app/controllers/demo"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/demo1", demo.Test1)
	router.POST("/demo2", demo.Test2)
	router.GET("/demo3", demo.TestDb)
	router.GET("/demo4", demo.TestRedis)
}
