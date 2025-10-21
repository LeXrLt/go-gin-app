package handlers

import (
	"database/sql"
	"go-gin-app/db"
	"go-gin-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DistillationDataRequest struct {
	SubTaskID        int64  `json:"sub_task_id" binding:"required"`
	Prompt           string `json:"prompt" binding:"required"`
	InferenceProcess string `json:"inference_process" binding:"required"`
	ModelOutput      string `json:"model_output" binding:"required"`
}

func UploadDistillationData(c *gin.Context) {
	var req DistillationDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if the sub_task exists and belongs to the user
	var subTask models.SubTask
	var task models.Task
	query := "SELECT st.id, st.task_id, st.status, t.user_id FROM sub_tasks st JOIN tasks t ON st.task_id = t.id WHERE st.id = ? AND t.user_id = ?"
	err := db.DB.QueryRow(query, req.SubTaskID, userID).Scan(&subTask.ID, &subTask.TaskID, &subTask.Status, &task.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "SubTask not found or you don't have permission to access it"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if subTask.Status == "completed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task is already completed"})
		return
	}

	// Update sub_task status to completed
	updateSubTaskQuery := "UPDATE sub_tasks SET status = 'completed' WHERE id = ?"
	_, err = db.DB.Exec(updateSubTaskQuery, req.SubTaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sub_task status"})
		return
	}

	// Check if all sub_tasks for the parent task are completed
	var pendingSubTasks int
	countQuery := "SELECT COUNT(*) FROM sub_tasks WHERE task_id = ? AND status = 'pending'"
	err = db.DB.QueryRow(countQuery, subTask.TaskID).Scan(&pendingSubTasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check sub_tasks status"})
		return
	}

	if pendingSubTasks == 0 {
		// Update parent task status to completed
		updateTaskQuery := "UPDATE tasks SET status = 'completed' WHERE id = ?"
		_, err = db.DB.Exec(updateTaskQuery, subTask.TaskID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task status"})
			return
		}
	}

	// Insert distillation data
	insertQuery := "INSERT INTO distillation_data (user_id, task_id, prompt, inference_process, model_output) VALUES (?, ?, ?, ?, ?)"
	_, err = db.DB.Exec(insertQuery, userID, subTask.TaskID, req.Prompt, req.InferenceProcess, req.ModelOutput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data uploaded and task completed successfully"})
}
