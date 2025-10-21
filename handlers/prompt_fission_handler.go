package handlers

import (
	"go-gin-app/db"
	"go-gin-app/models"
	"go-gin-app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PromptFissionRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type PromptFissionResponse struct {
	SubTaskID int64  `json:"sub_task_id"`
	Prompt    string `json:"prompt"`
}

func PromptFission(c *gin.Context) {
	var req PromptFissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Create parent task
	task := models.Task{
		UserID:    int64(userID.(int64)),
		Status:    "pending",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	taskQuery := "INSERT INTO tasks (user_id, status, expires_at) VALUES (?, ?, ?)"
	result, err := db.DB.Exec(taskQuery, task.UserID, task.Status, task.ExpiresAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	taskID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task ID"})
		return
	}

	fissionedPrompts, err := utils.GetOpenAIPromptFission(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []PromptFissionResponse
	subTaskQuery := "INSERT INTO sub_tasks (task_id, prompt, status) VALUES (?, ?, ?)"
	for _, p := range fissionedPrompts {
		subTask := models.SubTask{
			TaskID: taskID,
			Prompt: p,
			Status: "pending",
		}
		subTaskResult, err := db.DB.Exec(subTaskQuery, subTask.TaskID, subTask.Prompt, subTask.Status)
		if err != nil {
			// Log the error, but maybe don't fail the whole request?
			// For now, we'll fail it.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sub_task"})
			return
		}
		subTaskID, err := subTaskResult.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sub_task ID"})
			return
		}
		response = append(response, PromptFissionResponse{SubTaskID: subTaskID, Prompt: p})
	}

	c.JSON(http.StatusOK, response)
}
