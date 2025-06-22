package routes

import (
	"dts-developer-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.GET("/:id", controllers.GetTask)
		taskGroup.GET("/", controllers.GetAllTasks)
		taskGroup.PUT("/:id/status", controllers.UpdateTaskStatus)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}
}
