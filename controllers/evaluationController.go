package controllers

import (
	"kiko/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Submission struct {
	Code      string `json:"code"`
	ProblemID string `json:"problemId"`
}

func EvaluateCode(c *gin.Context) {
	var submission Submission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := services.RunCodeInDocker(submission.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": output})
		return
	}

	c.JSON(http.StatusOK, gin.H{"output": output})
}
