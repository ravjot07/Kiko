package routes

import (
	"kiko/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProblemRoutes(router *gin.Engine) {
	router.GET("/api/problems", controllers.GetProblems)
}
