package handlers

import (
	"go-gin-app/db"
	"go-gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadDistillationData(c *gin.Context) {
	var data models.DistillationData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	data.UserID = int(userID.(int64))

	query := "INSERT INTO distillation_data (user_id, prompt, inference_process, model_output) VALUES (?, ?, ?, ?)"
	_, err := db.DB.Exec(query, data.UserID, data.Prompt, data.InferenceProcess, data.ModelOutput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data uploaded successfully"})
}