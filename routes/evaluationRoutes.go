package routes

import (
	"kiko/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEvaluationRoutes(router *gin.Engine) {
	router.POST("/api/evaluate", controllers.EvaluateCode)
}
