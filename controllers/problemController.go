package controllers

import (
	"context"
	"kiko/models"
	"kiko/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProblems(c *gin.Context) {
	var problems []models.Problem
	cursor, err := utils.Client.Database("kiko").Collection("problems").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var problem models.Problem
		cursor.Decode(&problem)
		problems = append(problems, problem)
	}
	c.JSON(http.StatusOK, problems)
}
