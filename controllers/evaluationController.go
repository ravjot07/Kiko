package controllers

import (
	"context"
	"kiko/models"
	"kiko/services"
	"kiko/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	// Convert ProblemID to ObjectID
	objID, err := primitive.ObjectIDFromHex(submission.ProblemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	// Retrieve the problem from the database
	var problem models.Problem
	err = utils.Client.Database("reactPractice").Collection("problems").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&problem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Problem not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Combine the user code with the test code
	testCode := ""
	for _, test := range problem.Tests {
		testCode += test.TestCode + "\n"
	}

	// Run the code in Docker
	output, err := services.RunCodeInDocker(submission.Code, testCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "output": output})
		return
	}

	// Return the result
	c.JSON(http.StatusOK, gin.H{"output": output})
}
